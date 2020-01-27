# API Schema

## Upload

```http request
POST /api/upload HTTP/1.1
Content-Type:multi‐part/form-data

p=<PASSWD>&e=<EXPIRATION>&d=<DATA>
```

`EXPIRATION` should be integer from 1 to 24, 
set to 0 will result in burning after read,
if `EXPIRATION` is not set, default is 24, which means expired after 24h.

`PASSWD` should be the password you wanna use,
its length must be larger than or equal to 6.
if `PASSWD` is not set, no password will be required.
But your data will still be encrypted stored on our server.

Response Code:

- HTTP 200, Okay.
- HTTP 4xx, Failed (Usually Client Side Reason).
- HTTP 5xx, Failed due to server failure.
- HTTP 403, Content failed to pass checking.

Response Content:

```http request
Content-Type: text/plain
Content-Length: <ORIGINAL DATA LENGTH>

https://<DOMAIN>:<PORT>/<SHORTID> 
```

If you getting a notice about CAPTCHA test, it will show you another URL instead of the URL contains shortID. If you failed to pass, the content you uploaded will be deleted after 5 mins and will NOT be published.

## Show

```http request
GET /<SHORTID>?f=<FORMAT>&p=<PASSWD> HTTP/1.1
```

If `f` is raw, you will not get syntax-highlighted output.

The code syntax highlighting is done in client side using Google Prettify.js ,

Response Code:

- HTTP 200, Okay.
- HTTP 5xx, Failed due to server failure.
- HTTP 404, Content Nonexistent/Expired/Need Password to access.
- HTTP 4xx, Failed (Usually Client Side Reason).

Response Content:

- If `f` is not set, will output prettified code.
- If `p` is not set, but encryption required, will return 404.
- If `shortid` is not exists, will return 404.

## Delete

```http request
DELETE /api/admin/<shortid>?k=<masterkey> HTTP/1.1
```

You should set a master key in server config as administrator credential.

The `masterkey` param is the md5-hashed key concatenated with current date and hour in UTC. So make sure your time is correct.

Response Code:

- HTTP 200, Okay.
- HTTP 403, Authentication Error.
- HTTP 404, Not found.

## Recaptcha

