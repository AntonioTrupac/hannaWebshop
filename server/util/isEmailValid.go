package util

import (
	"fmt"
	"net"
	"net/mail"
	"strings"
)

// IsEmailValid checks if the email provided passes the required structure
// and length test. It also checks the domain has a valid MX record.
func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		fmt.Println("Email or domain is not valid!")
		return false
	}

	parts := strings.Split(email, "@")
	mx, error := net.LookupMX(parts[1])

	if error != nil || len(mx) == 0 {
		fmt.Println("Domain is not valid!")
		return false
	}

	return true
}
