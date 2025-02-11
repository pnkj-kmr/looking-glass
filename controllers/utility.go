package controllers

import (
	"strings"
)

// keyCipher
// var keyCipher []byte = []byte("thatislongpassphrase32bitJHreqwy")

func maskPasswd(plaintext string) string {
	// return strings.Repeat("*", len(plaintext))
	return strings.Repeat("*", 10)
}

func encryptAES(plaintext string) (string, error) {
	return plaintext, nil
	// TODO
	// c, err := aes.NewCipher(keyCipher)
	// if err != nil {
	// 	logger.Error("Encrypt Error", err)
	// 	return "", err
	// }
	// out := make([]byte, len(plaintext))
	// c.Encrypt(out, []byte(plaintext))
	// return hex.EncodeToString(out), nil
}

func decryptAES(ct string) (string, error) {
	return ct, nil
	// TODO
	// ciphertext, err := hex.DecodeString(ct)
	// if err != nil {
	// 	logger.Error("Decrypt Error", err)
	// 	return "", err
	// }
	// c, err := aes.NewCipher(keyCipher)
	// if err != nil {
	// 	logger.Error("Decrypt Cipher Error", err)
	// 	return "", err
	// }
	// pt := make([]byte, len(ciphertext))
	// c.Decrypt(pt, ciphertext)
	// s := string(pt[:])
	// return s, nil
}
