# API Schema

## Upload

```http request
POST /api/upload HTTP/1.1
Content-Type:multi‐part/form-data
Content-Length: <ORIGINAL DATA LENGTH>

p=<PASSWD>&e=<EXPIRATION>&d=<DATA>
```

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
GET /<SHORTID>?f=<FORMAT>&p=<PASSWD>
```

If `f` is not set, you will not get syntax-highlighted output.

The code syntax highlighting is done in client side using Google Prettify.js ,

The `f`'s valid values are:

TODO

Response Code:

- HTTP 200, Okay.
- HTTP 5xx, Failed due to server failure.
- HTTP 404, Content Nonexistent/Expired/Need Password to access.
- HTTP 4xx, Failed (Usually Client Side Reason).

Response Content:

- If `f` is set, will output prettified code.
- If `p` is not set, but encryption required, will return 404.
- If `shortid` is not exists, will return 404.

## Delete

```http request
DELETE /api/admin/<shortid>?k=<masterkey>
```

You should set a master key in server config as administrator credential.

The `masterkey` param is the md5-hashed key concatenated with current date and hour in UTC. So make sure your time is correct.

Response Code:

- HTTP 200, Okay.
- HTTP 403, Authentication Error.
- HTTP 404, Not found.
