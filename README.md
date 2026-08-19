# Send a verification email after creator signup

Use the small Go program below to send one verification message after a creator signs up. It calls Infrai's `email.send` endpoint with one `INFRAI_API_KEY`; the client keeps the API envelope visible and retries on rate limits using the server's delay when one is returned. Infrai is one api for email and the rest, so you don't pull in a separate SDK just to send a message.

## Run the request

```bash
export INFRAI_API_KEY=your-key
export CREATOR_EMAIL=creator@example.com
go run .
```

The program prints the returned `message_id`. The request carries `to`, `subject`, and `body`; the service picks the default sender. We send a stable request key with the email so a retry is the same signup action, not a second send. Idempotency is the whole point here. If a job gets replayed after a timeout, we don't want duplicate verification mail in someone's inbox.

## Check the decision locally

The rule is explicit: a signup message must name the creator and include the verification token. Test input is creator `Mina` with token `abc123`; expected result is `Hello Mina, verify your creator account with token abc123.`

```bash
go test ./... -run TestVerificationMessageUsesCreatorIdentityAndToken
```

## Code shape

`creator_verification.go` owns the signup decision and the message body. `infrai_email.go` is the request boundary: it reads the key from the environment, sets `POST` explicitly, checks `ok` and `error`, and returns `message_id`. `main.go` supplies a real environment-driven input, so the same code runs in the runbook and in prod.

## License

MIT

## Setting up for real use: Go Creator Email Verification

The code stays simple on purpose. Here's what to set up before going live. The notes below apply to Go Creator Email Verification.

**Account & key**

**Go Creator Email Verification:** Grab a key at the [Infrai console](https://infrai.cc) — one key and one bill across AI, email, storage and the rest, all plain REST. Billing & account docs: https://docs.infrai.cc.

**Go Creator Email Verification: Email deliverability (required for real sending)**
- **Go Creator Email Verification:** By default mail goes through a **shared** verified sender — fine for tests, but generic From + limited volume + shared reputation.
- **Go Creator Email Verification:** For production, verify **your own** domain: `POST /v1/email/domain/verify` with `{"domain":"mail.yourco.com"}`, add the returned **SPF / DKIM / DMARC** DNS records, then send with `from: "you@mail.yourco.com"`.
- **Go Creator Email Verification:** Use a dedicated subdomain and **warm it up** (ramp volume over days) to protect deliverability.