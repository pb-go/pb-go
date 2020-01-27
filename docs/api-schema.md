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

Response Content:

```http request
Content-Type: text/plain
Content-Length: <ORIGINAL DATA LENGTH>

https://<DOMAIN>:<PORT>/<SHORTID> 
```

## Show

```http request
GET /<SHORTID>?f=<FORMAT>&p=<PASSWD>
```

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

