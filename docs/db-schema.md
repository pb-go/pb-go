## Database Schema

Start a new database in mongodb.

Create a new collection called: "userdata" with default TTL 24h.

```mongodb
db.createCollection("userdata",{
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["waitVerify", "shortId", "userIP", "expireAt", "data", "pwdIsSet", "passwd"],
            properties: {
                waitVerify: {
                    bsonType: "bool",
                    description: "If captcha enabled, this will mark if publish or not."
                },
                shortId: {
                    bsonType: "string",
                    description: "shorter than 5 bytes, nanoid"
                },
                userIP: {
                    bsonType: "decimal",
                    description: "save user IP, including IPv6 support"
                },
                expireAt: {
                    bsonType: "date",
                    description: "expire time, max 24h"
                },
                data: {
                    bsonType: "binData",
                    description: "utf8 only, stored after chacha20 encrypted. userdata."
                },
                pwdIsSet: {
                    bsonType: "bool",
                    description: "check if password is set to use encryption."
                },
                passwd: {
                    bsonType: "string",
                    description: "blake2b hashed password"
                }
            }
        }
    },
    validationAction: "error"
    })
```

Set TTL to get document expired.

```mongodb
db.userdata.createIndex(
    {"expireAt": 1},
    {expireAfterSeconds: 0}
    );
db.userdata.createIndex(
    {"shortId": 1},
    {unique: true}
    );
```

