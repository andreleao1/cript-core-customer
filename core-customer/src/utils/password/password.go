package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log/slog"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)

	return bytes, err
}

func HashPassword(password string) (string, error) {
	salt, err := GenerateRandomBytes(16)

	if err != nil {
		slog.Error("Error generating random bytes.")
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hashedPassword := fmt.Sprintf("%s.%s", base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(hash))

	return hashedPassword, nil
}

func CheckPasswordHash(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, ".")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	compareHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if subtle.ConstantTimeCompare(hash, compareHash) == 1 {
		return true, nil
	}

	return false, nil
}
