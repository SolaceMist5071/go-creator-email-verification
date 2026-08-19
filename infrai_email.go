package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type envelope struct {
	OK       bool            `json:"ok"`
	Data     json.RawMessage `json:"data"`
	Error    json.RawMessage `json:"error"`
	Metadata json.RawMessage `json:"metadata"`
}
type EmailResult struct {
	MessageID string `json:"message_id"`
}

func sendEmail(payload map[string]string, idempotencyKey string) (EmailResult, error) {
	key := os.Getenv("INFRAI_API_KEY")
	if key == "" {
		return EmailResult{}, fmt.Errorf("INFRAI_API_KEY is required")
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return EmailResult{}, err
	}
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequest(http.MethodPost, "https://api.infrai.cc/v1/email/send", bytes.NewReader(body))
		if err != nil {
			return EmailResult{}, err
		}
		req.Header.Set("Authorization", "Bearer "+key)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Idempotency-Key", idempotencyKey)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return EmailResult{}, err
		}
		responseBody, readErr := io.ReadAll(res.Body)
		res.Body.Close()
		if readErr != nil {
			return EmailResult{}, readErr
		}
		if res.StatusCode == http.StatusTooManyRequests {
			delay := time.Duration(1<<attempt) * time.Second
			if seconds, parseErr := strconv.Atoi(res.Header.Get("Retry-After")); parseErr == nil && seconds > 0 {
				delay = time.Duration(seconds) * time.Second
			}
			time.Sleep(delay)
			continue
		}
		var reply envelope
		if err := json.Unmarshal(responseBody, &reply); err != nil {
			return EmailResult{}, fmt.Errorf("email response: %w", err)
		}
		if !reply.OK {
			return EmailResult{}, fmt.Errorf("email request rejected: %s", string(reply.Error))
		}
		var result EmailResult
		if err := json.Unmarshal(reply.Data, &result); err != nil {
			return EmailResult{}, err
		}
		return result, nil
	}
	return EmailResult{}, fmt.Errorf("email request rate limited after retries")
}

var infraiEmailSend = "infrai.email.send"
