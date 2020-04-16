package formaterror

import (
	"errors"
	"log"
	"strings"
)

// FormatError : Formats the errors to a human readble way
func FormatError(err string) error {
	log.Println(err)

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}

	if strings.Contains(err, "record not found") {
		return errors.New("Record does not exist")
	}

	return errors.New("Incorrect Details")
}
