package security

import "golang.org/x/crypto/bcrypt"

//Hash receive string and apply hash
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

//VerifyPassword compare one pass and one hash and return if is equal
func VerifyPassword(passHash, passString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passString))
}
