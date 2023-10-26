package service

import (
	"Uploader/internal/errorext"
	"bytes"
	"github.com/minio/minio-go/v7"
	"io"
	"mime/multipart"
	"strings"

	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"unicode"
)

func isEmailValid(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regex, email)

	if err != nil {
		return false
	}

	return match
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func comparePasswords(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func checkLettersInPassword(password string) bool {
	hasUppercase := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true

			break
		}
	}

	hasLowercase := false

	for _, char := range password {
		if unicode.IsLower(char) {
			hasLowercase = true

			break
		}
	}

	hasDigit := false

	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true

			break
		}
	}

	hasSpecialChar := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}

func isPasswordValid(password string) bool {
	if len(password) < 8 || len(password) > 50 {
		return false
	}

	return checkLettersInPassword(password)
}

func checkEmailAndPassword(email, password string) error {
	if ok := isPasswordValid(password); !ok {
		return errorext.NewValidationError(fmt.Sprintf("password is invalid"))
	}

	if ok := isEmailValid(email); !ok {
		return errorext.NewValidationError(fmt.Sprintf("email is invalid"))
	}

	return nil
}

func getFileExtension(text string) string {
	splitByDot := strings.Split(text, ".")
	return splitByDot[len(splitByDot)-1]
}

func getBufferFromMultipartFile(file multipart.File) (*bytes.Buffer, error) {
	reader, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(reader)
	return buff, nil
}

func getFileFromMinioObject(object *minio.Object) (*bytes.Buffer, error) {
	reader, err := io.ReadAll(object)

	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(reader)

	return buff, nil

}

func generateBucketName(userId uint) string {
	return fmt.Sprintf("user-%d", userId)
}
