package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// JWT is JSON Web Token
type JWT struct {
	header    string
	payload   string
	signature string
}

func parseJWT(jwt string) JWT {
	x := strings.Split(jwt, ".")
	return JWT{
		header:    x[0],
		payload:   x[1],
		signature: x[2],
	}
}

func hs256(data, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return strings.ReplaceAll(base64.URLEncoding.EncodeToString(mac.Sum(nil)), "=", "")
}

func main() {
	token := flag.String("t", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjpudWxsfQ.Tr0VvdP6rVBGBGuI_luxGCOaz6BbhC6IxRTlKOW8UjM", "JWT string")
	wordlistPath := flag.String("w", "", "Path to wordlist")
	flag.Parse()
	if *wordlistPath == "" {
		log.Fatalf("Please provide a wordlist")
	}

	jwt := parseJWT(*token)

	file, err := os.Open(*wordlistPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		secret := scanner.Text()
		sign := hs256(jwt.header+"."+jwt.payload, secret)
		if sign == jwt.signature {
			fmt.Println(secret)
			return
		}
	}
}
