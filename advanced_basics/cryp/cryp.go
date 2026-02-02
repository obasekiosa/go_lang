package cryp

import (
	"crypto/sha256"
	"fmt"
)

func littleCrypt() {
	buf := []byte{}
	fmt.Println(buf)
	hash := sha256.Sum256([]uint8("x"))
	fmt.Printf("%X\n", hash)
	fmt.Println(hash)
	fmt.Println(buf)
}
