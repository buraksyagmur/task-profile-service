// task-profile-service/utils/verification_code.go

package utils

import (
	"crypto/rand"
	"math/big"
	"time"
)
func GenerateVerificationCode() (string, error) {
	code, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}
	code.Add(code, big.NewInt(100000))
	return code.String(), nil
}
func GetExpirationTime() time.Time {
	return time.Now().Add(10 * time.Minute)
}
