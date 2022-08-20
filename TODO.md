# To Do

- [x] get proxy working for http
- [x] status page
- [x] favicon/robots
- [x] option to also proxy HTTP on port 80
- [x] License.txt
- [x] ability to switch to production LetsEncryptCA
- [x] text on index.html

- [x] air for local dev
- [ ] certmagic tls-alpn challenge
- [ ] ondemand decision function: always
- [ ] ondemand decision function: dns check

- [ ] build with GHA
- [ ] deploy with GHA

- [ ] ondemand decision function: only eTLD+1
- [ ] ondemand decision function: allowlist
- [ ] ondemand decision function: http api call
- [ ] certmagic custom storage
- [ ] certmagic renewals

long term:
- local 404 handler
- structured logging (zap?)
- metrics
- logging
- graceful shutdown
- allowlist from db
- use pure go DNS resolver
- serving http: [ none | redirect | proxy ]
- control over "Host" header sent to target
- HSTS
- timeouts
- Dockerfile
- [ZeroSSL](https://zerossl.com/documentation/acme/) option
- [BuyPass](https://www.buypass.com/products/tls-ssl-certificates/go-ssl) option
- [SSL.com](https://www.ssl.com/guide/ssl-tls-certificate-issuance-and-revocation-with-acme/) option

https://github.com/pilu/fresh
https://blog.charmes.net/post/reverse-proxy-go/
https://www.integralist.co.uk/posts/golang-reverse-proxy/
https://github.com/kasvith/simplelb/blob/master/main.go)
https://kasvith.github.io/posts/lets-create-a-simple-lb-go/

PR:
- https://letsencrypt.org/docs/client-options/
- https://github.com/dariubs/awesome-proxy#reverse-proxy

Package:
- https://go-team.pages.debian.net/packaging.html
- https://github.com/Debian/dh-make-golang
