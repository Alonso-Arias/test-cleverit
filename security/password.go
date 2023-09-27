package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"regexp"

	"github.com/trustelem/zxcvbn"

	"strconv"

	errs "github.com/Alonso-Arias/test-cleverit/errors"
	"github.com/Alonso-Arias/test-cleverit/log"
	"github.com/raja/argon2pw"
)

var loggerf = log.LoggerJSON().WithField("package", "security")

var SecretKey = "******"

type PasswordHash interface {
	Hash(p string) (string, error)
	Compare(p string, hash string) (bool, error)
}

type PasswordHashImpl struct {
}

func (ph PasswordHashImpl) Hash(p string) (string, error) {
	log := loggerf.WithField("func", "Hash")
	hashedPassword, err := argon2pw.GenerateSaltedHash(p)
	if err != nil {
		log.Error("hashedPassword fail")
		return "", err
	}

	return hashedPassword, nil
}

func (ph PasswordHashImpl) Compare(p string, hash string) (bool, error) {
	log := loggerf.WithField("func", "Compare")
	valid, err := argon2pw.CompareHashWithPassword(hash, p)
	if err != nil {
		log.Error("CompareHashWithPassword fail")
		return false, err
	}
	//log.Info("CompareHashWithPassword valid : " + valid)
	return valid, nil
}

type StrengthPassword string

const (
	Weak   StrengthPassword = "WEAK"
	Medium StrengthPassword = "MEDIUM"
	Strong StrengthPassword = "STRONG"
)

var containUpper = regexp.MustCompile(`[A-Z]`).MatchString
var containLower = regexp.MustCompile(`[a-z]`).MatchString
var containDigit = regexp.MustCompile(`[0-9]`).MatchString

type PasswordPolicy interface {
	Validate(password string) (StrengthPassword, error)
}

type PasswordPolicyImpl struct {
}

func (pp PasswordPolicyImpl) Validate(p string) (StrengthPassword, error) {

	log := loggerf.WithField("func", "Validate")

	prefix := "[" + p + "]"

	result := Weak
	evaluation := zxcvbn.PasswordStrength(p, nil)

	log.Info(prefix + " Password points :" + strconv.Itoa(evaluation.Score))

	if evaluation.Score >= 3 && evaluation.Score <= 4 {
		result = Medium
	}

	if evaluation.Score >= 5 {
		result = Strong
	}

	if len(p) < 8 {
		return result, errs.LenPassPolicy
	}

	log.Info(prefix + " Policy containUpper")

	if !containUpper(p) {
		return result, errs.UpperPassPolicy
	}

	log.Info(prefix + " Policy containLower")

	if !containLower(p) {
		return result, errs.LowerPassPolicy
	}

	log.Info(prefix + " Policy containDigit")

	if !containDigit(p) {
		return result, errs.DigitPassPolicy
	}

	return result, nil
}

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

type DataEncrypt interface {
	Encrypt(key, text string) (string, error)
	Decrypt(key, text string) (string, error)
}

type DataEncryptImpl struct {
}

// encrypt encrypts plain string with a secret key and returns encrypt string.
func (pp DataEncryptImpl) Encrypt(secret, plainData string) (string, error) {
	cipherBlock, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(plainData), nil)), nil
}

// decrypt decrypts encrypt string with a secret key and returns plain string.
func (pp DataEncryptImpl) Decrypt(secret, encodedData string) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainData), nil
}
