package pkg

import "encoding/base64"

func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Base64Decode(text string) string {
	data, _ :=  base64.StdEncoding.DecodeString(text)
	return string(data)
}
