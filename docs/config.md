# Server configuration 

## Configuration Example



## Database

MongoDB should be configured as described in DB schema file to make sure the data TTL is correctly set. You do need to add SCRAM authentication to your DB and make sure your DB listened just on localhost.

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

## Network

Set the listen port and domain and IP in server configuration, in YAML format.
