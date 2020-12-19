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
go run main.go \ 
-w wordlist.txt \ 
-t eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjpudWxsfQ.Tr0VvdP6rVBGBGuI_luxGCOaz6BbhC6IxRTlKOW8UjM
```