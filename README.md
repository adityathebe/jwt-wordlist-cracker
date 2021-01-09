# JWT CRACKER

Use wordlist to crack JWT HS256

## Build

```bash
go build -o jwtcracker main.go
```

## Example

```bash
# wordlist.txt

hacker
jwt
insecurity
pentesterlab
hacking
```

```bash
./jwtcracker \
-w wordlist.txt \
-t eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjpudWxsfQ.Tr0VvdP6rVBGBGuI_luxGCOaz6BbhC6IxRTlKOW8UjM
```

## John the Ripper (Alternative)

```bash
# File: jwt.txt
# eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjpudWxsfQ.Tr0VvdP6rVBGBGuI_luxGCOaz6BbhC6IxRTlKOW8UjM

john jwt.txt --wordlist=wordlist.txt --format=HMAC-SHA256
```

## Resources

- [Wordlist - public JWT secrets found with Google dorking and Google BigQuery](https://github.com/wallarm/jwt-secrets)
