# jwt-parser

##### Simple cli jwt parser

[![GitHub](https://img.shields.io/github/license/laststem/jwt-parser)](https://github.com/laststem/jwt-parser/blob/master/LICENSE)

### Install
```bash
go get -u github.com/laststem/jwt-parser
```

### Usage
- Only show parsed token: jwt-parser {token}
- show parsed token with specific key: jwt-parser {token} {key} {key} {key} ...

## Example
```bash
$ jwt-parser eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
{
    "iat": 1516239022,
    "name": "John Doe",
    "sub": "1234567890"
}

$ jwt-parser eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c iat
1516239022

# if you have nested fields and want to see a specific field, just put listing keys on arguments.
# if claim looks like this
{
    "key1": {
        "key2": {
            "key3": "hello"      
        }
    },
    "other": ["jwt", "parser", "cli"]
}
$ jwt-parser <token> key1 key2
{
    "key3": "hello"
}

$ jwt-parser <token> key1 key2 key3
"hello"

$ jwt-parser <token> other
[
    "jwt",
    "parser",
    "cli"
]
``` 
