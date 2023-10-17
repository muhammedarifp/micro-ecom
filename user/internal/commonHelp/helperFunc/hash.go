package helperfunc

import "golang.org/x/crypto/bcrypt"

func PasswordHashing(pass string) string {
	passByte, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(passByte)
}

func ComparePassword(hashPass, orgPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(orgPass))

}
