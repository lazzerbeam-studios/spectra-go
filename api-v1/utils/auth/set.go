package auth

var SecretKey []byte

func SetSecret(secret string) {
	SecretKey = []byte(secret)
}
