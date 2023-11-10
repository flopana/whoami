package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

// GenUUID generates a random uuid
func GenUUID() string {
	var bytes = getRandBytes(16)
	var uuid = make([]byte, 36)
	encodeHex(&uuid, &bytes)

	return string(uuid)
}

func getRandBytes(n int) []byte {
	var bytes = make([]byte, n)
	_, err := rand.Read(bytes)

	if err != nil {
		panic(err)
	}
	return bytes
}

// encodeHex encodes a byte slice into a hex string in the format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func encodeHex(dst *[]byte, src *[]byte) {
	hex.Encode(*dst, (*src)[:4])
	(*dst)[8] = '-'
	hex.Encode((*dst)[9:13], (*src)[4:6])
	(*dst)[13] = '-'
	hex.Encode((*dst)[14:18], (*src)[6:8])
	(*dst)[18] = '-'
	hex.Encode((*dst)[19:23], (*src)[8:10])
	(*dst)[23] = '-'
	hex.Encode((*dst)[24:], (*src)[10:])
}
