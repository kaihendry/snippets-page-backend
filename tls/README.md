**How to run an HTTP/2 server? Generate a self-signed X.509 TLS certificate**\
Run the following command to generate cert.pem and key.pem files:\
`go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost`\
For prod you should obtain a certificate from CA.