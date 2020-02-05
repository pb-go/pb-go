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

If `f` is `raw`, you will not get syntax-highlighted output. Other values are invalid.

The code syntax highlighting is done in client side using Google Prettify.js ,

Prettify.js will automatically detect the programming language of the snippet. (Yes, same version as SOF!)

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
DELETE /api/admin?id=<shortid> HTTP/1.1
```

You should set a master key in server config as administrator credential.

The `X-Master-Key` param is the md5-hashed key concatenated with current date and hour in UTC. 
So make sure your time is correct (at least on hour level), and this should be set in `X-Master-Key` of HTTP Header.

For Example:

The password is `123456`, UTC Time represented in RFC822 Format `RFC822 = "02 Jan 06 15:04 MST"` is: 
`02 Jan 06 15:04 UTC`.

So the finally key should be represent as: `{123456}{02 Jan 06 15:04 UTC}`  **(Don't forget `{}`).**

The hashed key is : `3c841b5c136c47dfb764119d41c7c5c1`

The legal HTTP POST request should be:

```http request
DELETE /api/admin?id=<SHORTID> HTTP/1.1
X-Master-Key: 3c841b5c136c47dfb764119d41c7c5c1
```

Response Code:

- HTTP 200, Okay.
- HTTP 403, Authentication Error.
- HTTP 404, Not found.

## Recaptcha

```http request
POST /api/g_verify HTTP/1.1

g-recaptcha-response=<BLAHBLAH>&tempID=<BLAHBLAH>
```

The API path is above. 

Ask user to run reCAPTCHA test. Server-side verification.

If the `recaptcha.enable` inside the server administrators' config is true, then we will try to return 
a URI like this `/showVerify?id=<SNIPPET ID WITH URLSAFE BASE64 ENCODED>` and ask user to continue.

If user failed, this snippet will be automatically expired in 2 mins. Else it will give user the published path.

## Server Status

```http request
GET /status HTTP/1.1
```

This request will return the configuration of a server, return a json.

Including but MAY not limited to:

-  Server is running
-  Server requires recaptcha or not
-  Server force expire time set
-  Server anti-abuse feature enabled or not

The response can be easily interpreted as the word said.

The schema is:

```json
{
  "status": 0,
  "captcha_enabled": true,
  "max_expire": 12,
  "abuse_detection": true,
  "base64_detection": true
}
```

Status code is always `0`, means running, other values should be set according to server config.
