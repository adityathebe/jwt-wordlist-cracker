package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
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

func parseJWT(jwt string) (JWT, error) {
	x := strings.Split(jwt, ".")
	if len(x) != 3 {
		return JWT{}, errors.New("invalid JWT. Not enough or too many segments")
	}
	return JWT{
		header:    x[0],
		payload:   x[1],
		signature: x[2],
	}, nil
}

func hs256(data, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func main() {
	token := flag.String("t", "", "JWT string")
	wordlistPath := flag.String("w", "", "Path to wordlist")
	flag.Parse()
	if *wordlistPath == "" || *token == "" {
		flag.Usage()
		return
	}

	jwt, err := parseJWT(*token)
	if err != nil {
		log.Fatalln(err)
	}

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
