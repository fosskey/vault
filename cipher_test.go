package vault

import (
	"bytes"
	"testing"
)

func TestCipher(t *testing.T) {
	var password = []byte("MyP@ssw0rd")
	var message = []byte("Hello World!")

	ciphertext, err := encrypt(password, message)
	if err != nil {
		t.Fatal(err)
	}

	if len(ciphertext) <= len(message) {
		t.Fatal("Encrypted result must be longer than the original message")
	}

	plaintext, err := decrypt(password, ciphertext)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(plaintext, message) {
		t.Fatal("Decrypted result did not match with original message")
	}
}
