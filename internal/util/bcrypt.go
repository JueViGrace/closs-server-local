package util

import (
	"crypto/md5"
	"encoding/hex"
)

func ValidatePassword(reqPass, encPass string) bool {
	hash := md5.Sum([]byte(reqPass))
	hashedPass := hex.EncodeToString(hash[:])
	return hashedPass == encPass
}
