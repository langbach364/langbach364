package main

import (
	"time"
	"math/rand"
	"github.com/go-mail/mail"
)

var codeStore = make(map[string]string)
var codeExpiry = make(map[string]time.Time)

const charset = "0123456789"
const tokenLength = 6

var seededRandPkg *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRandPkg.Intn(len(charset))]
	}
	return string(b)
}

func randomToken(length int) string {
	return StringWithCharset(length, charset)
}

func Send_Email(sender string, password string, recevier string) {
	m := mail.NewMessage()
	token := randomToken(tokenLength)

	m.SetHeader("From", sender)
	m.SetHeader("To", recevier)
	m.SetBody("text/plain", token)
	m.SetHeader("Subject", "Cảm ơn bạn đã xem")

	d := mail.NewDialer("smtp.gmail.com", 587, sender, password)
	err := d.DialAndSend(m)
	check_err(err)

	codeStore[recevier] = token
	codeExpiry[recevier] = time.Now().Add(5 * time.Minute)
}


func verify_code(email string, code string) bool {
	storedCode, ok := codeStore[email]
	if !ok {
		return false
	}

	if time.Now().After(codeExpiry[email]) {
		delete(codeStore, email)
		delete(codeExpiry, email)
		return false
	}

	return storedCode == code
}
