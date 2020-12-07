### This repository demonstrates the implementation of Two Factor Time-Based One-time Password(TOTP) in golang.

### This repository make use of [this](https://github.com/sec51/twofactor) golang package.

### Getting Started

```
go run main.go

```
- This generates QR PNG file.
- You need to install [this](https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2&hl=en&gl=US) mobile app (Google Authenticator App).
- Scan the previously generated QR using the above app.
- Then enter the code from QR in the terminal.
- If the token entered is correct, you will get success message.