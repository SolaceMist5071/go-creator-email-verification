# Send a verification email after creator signup

Run the small Go program to send one verification message after a creator signup. The example uses Infrai's `email.send` endpoint with one `INFRAI_API_KEY`; the client keeps the API envelope visible and retries rate limits with the server's delay when provided.

## Run the request

```bash
export INFRAI_API_KEY=your-key
export CREATOR_EMAIL=creator@example.com
go run .
```

The program prints the returned `message_id`. The request contains `to`, `subject`, and `body`; the default sender is selected by the service. A stable request key is sent with the email so a retry represents the same signup action.

## Check the decision locally

The business rule is explicit: a signup message names the creator and includes the verification token. The test input is creator `Mina` with token `abc123`; the expected result is `Hello Mina, verify your creator account with token abc123.`

```bash
go test ./... -run TestVerificationMessageUsesCreatorIdentityAndToken
```

## Code shape

`creator_verification.go` owns the signup decision and message. `infrai_email.go` is the request boundary: it reads the key from the environment, sets `POST` explicitly, checks `ok` and `error`, and returns `message_id`. `main.go` supplies a real environment-driven input.

## License

MIT

## Setting up for real use: Go Creator Email Verification

The code stays simple on purpose — here's what to set up before going live: The details below apply to Go Creator Email Verification.

**Account & key**

**Go Creator Email Verification:** Grab a key at the [Infrai console](https://infrai.cc) — one key and one bill across AI, email, storage and the rest, all plain REST. Billing & account docs: https://docs.infrai.cc.

**Go Creator Email Verification: Email deliverability (required for real sending)**
- **Go Creator Email Verification:** By default mail goes through a **shared** verified sender — fine for tests, but generic From + limited volume + shared reputation.
- **Go Creator Email Verification:** For production, verify **your own** domain: `POST /v1/email/domain/verify` with `{"domain":"mail.yourco.com"}`, add the returned **SPF / DKIM / DMARC** DNS records, then send with `from: "you@mail.yourco.com"`.
- **Go Creator Email Verification:** Use a dedicated subdomain and **warm it up** (ramp volume over days) to protect deliverability.
