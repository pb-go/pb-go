# Debug

Since our program request you to use TLS and requests must have an `X-Real-IP` exists, you need to reverse proxy the program instance.

If you don't do that, your requests will always be 502-ed.

We use CaddyServer v1, the prebuilt binary with full-plugin support we needed can be get from [here][1], `Caddyfile` is the config file it used.

[1]: https://filebin.kmahyyg.xyz/caddy_v1.tar.zst

## TLS Cert

Go back to search `mkcert` on GitHub, that will help you.
