package main

import (
	"fmt"
	"golab/encryption/sign"
)

func main() {
	src := "hello world"
	key := "123456"

	fmt.Println(sign.Md5Str(src))
	fmt.Println(sign.HmacSha1(src, key))
	fmt.Println(sign.HmacSha256(src, key))
	fmt.Println(sign.HmacSha512(src, key))
	fmt.Println(sign.Sha1str(src))
	fmt.Println(sign.Sha256str(src))
	fmt.Println(sign.Sha512str(src))
	fmt.Println(sign.Base64EncodeStr(src))
	fmt.Println(sign.Base64DecodeStr(sign.Base64EncodeStr(src)))

}
