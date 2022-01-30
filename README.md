# Redirect2Me HTTPS Proxy  [<img alt="r2proxy logo" src="assets/favicon.svg" height="90" align="right" />](https://redirect2.me/)

This is a server that proxies HTTPS requests to a separate (non-HTTPS) server, automatically handling certificate provisioning and renewals.

## How it works

* The server listens on port 443.
* If necessary, a certificate is automatically provisioned.
* All requests are proxied to the specified target `host:port`.

## Configuration

### Allowed hostnames

In order to prevent abuse, you may need to limit the hostnames that are allowed.  The `allowed` parameter:
* `*` - all hostnames (default)
* `api:url` - call an external API (Coming soon)
* `list:host1,host2,...` - list of allowed hostnames (Coming soon)
* `etld1` - only hostnames a single level under a public suffix (or `www` + single level) (Coming soon)

### DNS Check

Certificate provisioning will only work if the DNS is configured correctly, so this is checked before provisioning starts.  You can disable this by setting `--dnscheck=false`.

### Certificate storage

Coming soon

## Contributions

Contributions are welcome!

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Pico CSS](https://www.vectorlogo.zone/logos/picocss/picocss-ar21.svg)](https://picocss.com/ "CSS")
[![Python](https://www.vectorlogo.zone/logos/python/python-ar21.svg)](https://www.python.org/ "Test origin webserver")
[![svgrepo](https://www.vectorlogo.zone/logos/svgrepo/svgrepo-ar21.svg)](https://www.svgrepo.com/svg/31307/server "Icon")

