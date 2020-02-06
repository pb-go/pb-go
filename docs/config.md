# Server configuration 

## Configuration Example

```yaml
network:
  listen: 127.0.0.1:8181
  host: localhost
  mongodb_url: mongodb+srv://
recaptcha:
  enable: false
  secret_key:
  site_key:
security:
  master_key:
  encryption_key:
  encryption_nonce:
content:
  detect_abuse: true
  expire_hrs: 24
```

Default listens at `127.0.0.1:8181`, Configure reCAPTCHA v2 related key and enable in `recaptcha` part.

The host will be used to generate response to user, you should input your domain here, the finally output 
will be like this: `https://<network.host>/<otherdata>`, so make sure you've covered domain and port correctly.

The `masterkey` must be longer than 12 bytes, `encryption_key` will be used for storage encryption, must be equals to 32 bytes.

`encryption_nonce` will be used cooperate with `encryption_key`, must be equals to 12 bytes.

If `content.detect_abuse` is enabled, the system will only allow to upload pure text.

`content.expire_hrs` defined the maximum TTL by default, CANNOT and SHOULD NOT BE LONGER than 24, the unit is hour.

> Note: `expire_hrs` is the maximum TTL allowed, user may override their own snippet TTL by define themselves, but
> it cannot be longer than 24h, cuz the database will automatically expire after 24h in consideration of storage and 
> anti-abuse. Set to `0` means burn after read, the valid values must be 0~24h.

## Database

MongoDB should be configured as described in DB schema file to make sure the data TTL is correctly set. You do need to add SCRAM authentication to your DB and make sure your DB listened just on localhost.

Please DO NOT APPEND database name in URI.

### FAQ about DB

1. I always encounter MongoDB Connection Error especially when trying to establish first connection.

If you use `mongodb+srv://` link, please check [here](https://godoc.org/go.mongodb.org/mongo-driver/mongo#hdr-Potential_DNS_Issues) 
and try switch to another DNS. Otherwise, check your internet connection is stable and connected or not.

Official Explanation:

> Building with Go 1.11+ and using connection strings with the "mongodb+srv" scheme 
> is incompatible with some DNS servers in the wild due to the change introduced 
> in https://github.com/golang/go/issues/10622. If you receive an error with the message 
> "cannot unmarshal DNS message" while running an operation, we suggest you use a different DNS server.

### Garbage Collection (DB Storage Reuse)

Run following command: That will block DB operation! Have a backup first:

```js
db = db.getSiblingDB('pbgo');
db.runCommand( { compact : 'userdata' } );
```

You could try to write a cron job to do this at the no-request time.

Save the above request to a file, name it as `compact_mdb.js`

Then Create a Cron Job like this:

```crontab
15 2 */2 * * /usr/bin/mongo -u <USERNAME> -p <PASSWORD> --host <HOST> --port <PORT> --eval /usr/local/compact_mdb.js
```


## Anti-Abuse

Our application will just do the content check by examined the uploaded data, we will only allow the pure text. Any binary file or unknown file will be rejected immediately.

About the rate limitation, since our application doesn't implement it and is a public stateless service. The only thing we can do is recording your IP and do rate limitation according to that.

Normally, it should be set to about 10 reqs/min/IP.

Our goal is to offer convenient service to developers, not for spammer.

If your public instance is experiencing abusing, please do enable recaptcha and set API key. If enabled, All requests will need to pass recaptcha test in 5 mins before finally published. 

## Security

We use the password you offered or the default encryption password set in server configuration using CHACHA20 algorithm to make sure your data safety.

BUT, DO REMEMBER, we don't guarantee about unintended data loss like hardware failure and maintenance.

Do Remember: `masterkey` must be longer than 12 bytes. Encryption key should be 32 bytes and corresponding nonce should be 12 bytes.

One thing you must listen: correctly set the reverse proxy and offer a `X-Real-IP` header in http request, if this is not offered, we'll decline the request immediately.

## Network

Set the listen port and domain and IP in server configuration, in YAML format.
