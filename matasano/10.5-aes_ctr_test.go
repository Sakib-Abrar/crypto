package matasano

import (
	"crypto/rand"
	"testing"

	"github.com/nindalf/crypto/aes"
)

func TestCTREncryptDecrypt(t *testing.T) {
	input := "ATTACK AT DAWN!!ATTACK AT DAWN!!"
	testCTR(t, input)
	testCTR(t, input[0:1])
	testCTR(t, input[0:15])
	testCTR(t, input[0:16])
	testCTR(t, input[0:17])
}

func testCTR(t *testing.T, input string) {
	b := []byte(input)
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)

	EncryptAESCTR(b, key, iv)
	if string(b) == input {
		t.Fatalf("Failed to encrypt - %s", input)
	}

	DecryptAESCTR(b, key, iv)
	if string(b) != input {
		t.Fatalf("Failed to decrypt - %s, %s", input, string(b))
	}
}

func TestCTRDecrypt(t *testing.T) {
	c := []byte{0xe4, 0x62, 0x18, 0xc0, 0xa5, 0x3c, 0xbe, 0xca, 0x69, 0x5a, 0xe4, 0x5f, 0xaa, 0x89, 0x52, 0xaa, 0xe, 0x37, 0xd, 0xdf, 0xd9, 0x55, 0x5, 0x3f, 0x6f, 0x78, 0x96, 0xc0, 0x45, 0x2, 0xfc, 0xc7, 0x1b, 0x23, 0xd6, 0x35, 0x7d, 0xbe, 0x62, 0xe, 0x52, 0x8d, 0x70, 0xf7, 0xb0, 0xd1, 0x85, 0x32, 0xf5, 0x11}
	expected := "Always avoid the random gibberish while decrypting"

	key := []byte{0x36, 0xf1, 0x83, 0x57, 0xbe, 0x4d, 0xbd, 0x77, 0xf0, 0x50, 0x51, 0x5c, 0x73, 0xfc, 0xf9, 0xf2}
	iv := []byte{0x77, 0xb, 0x80, 0x25, 0x9e, 0xc3, 0x3b, 0xeb, 0x25, 0x61, 0x35, 0x8a, 0x9f, 0x2d, 0xc6, 0x17}

	DecryptAESCTR(c, key, iv)

	if string(c) != expected {
		t.Fatalf("Failed to decrypt in CTR mode. Expected - %s\tFound - %s", expected, string(c))
	}
}