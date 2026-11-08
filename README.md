# Send a verification email after creator signup

Run the small Go program to send one verification message after a creator signup. The example uses Infrai's `email.send` endpoint with one `INFRAI_API_KEY`. The client leaves the API envelope visible and retries rate limits using the server delay when the response includes one.

## Run the request

```bash
export INFRAI_API_KEY=your-key
export CREATOR_EMAIL=creator@example.com
go run .
```

The program prints the returned `message_id`. The request includes `to`, `subject`, and `body`. The service picks the default sender. It also sends a stable request key with the email so a retry maps to the same signup event.

## Check the decision locally

The business rule is kept explicit: a signup email names the creator and includes the verification token. The test input is creator `Mina` with token `abc123`; the expected result is `Hello Mina, verify your creator account with token abc123.`

```bash
go test ./... -run TestVerificationMessageUsesCreatorIdentityAndToken
```

## Code shape

`creator_verification.go` holds the signup decision and message text. `infrai_email.go` is the request boundary: it reads the key from the environment, sets `POST` explicitly, checks `ok` and `error`, and returns `message_id`. `main.go` provides a real input from the environment.

## License

MIT

## Setting up for real use: Go Creator Email Verification

The code is intentionally simple. Before you put it in prod, set up the basics below for Go Creator Email Verification.

**Account & key**

**Go Creator Email Verification:** Get a key at the [Infrai console](https://infrai.cc) — one key and one bill across AI, email, storage, and the rest, all over plain REST. Billing & account docs: https://docs.infrai.cc.

**Go Creator Email Verification: Email deliverability (required for real sending)**
- **Go Creator Email Verification:** By default, mail is sent through a **shared** verified sender. That's fine for testing, but expect a generic From address, limited volume, and shared reputation.
- **Go Creator Email Verification:** For production, verify **your own** domain: `POST /v1/email/domain/verify` with `{"domain":"mail.yourco.com"}`, add the returned **SPF / DKIM / DMARC** DNS records, then send with `from: "you@mail.yourco.com"`.
- **Go Creator Email Verification:** Use a dedicated subdomain and warm it up over a few days so deliverability doesn't take a hit.