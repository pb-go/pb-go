## Database Schema

Start a new database in mongodb.

Create a new collection called: "userdata" with default TTL 24h.

```mongodb
db.createCollection("userdata",{
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["shortId", "userIP", "expireAt", "data", "pwdIsSet", "passwd"],
            properties: {
                shortId: {
                    bsonType: "string",
                    description: "shorter than 5 bytes, nanoid"         
                },
                userIP: {
                    bsonType: "Decimal128",
                    description: "save user IP, including IPv6 support"
                },
                expireAt: {}           
            }
        }   
    },
    validationAction: "error"
)
```

Set TTL to get document expired.

```mongodb
db.userdata.createIndex(
    {"expireAt": 1},
    {expireAfterSeconds: 0}
    );
```