package main


import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"bytes"
)

const (
	key = "909dfef0d4242e7ed94abcc2be9cc747"
	iv  = "37cf1d51a2192ad5"
)

func main() {
	str := "我勒个去"
	es, _ := AesEncrypt(str, []byte(key))
	fmt.Println(es)

	dfd:="8b78161dd4ae844629deb86143449339fb8e001b93b3b4b2861bf8a3837bf97059a7a9e49739cce002fef5e64467a77eee8a21714aa7aed77c03984353429322e71b9721ee66849b3a029224ee4d5688908022f7a776179074dcbe25ec6fa3b33b9ed9495a0000ea4586fd543f081ca3bb0f7c7da5e65f3f97806c6229658adcb3be1d4d61fcf774fbbd443cec2a30167c2b8addafae1a8cbfda7231845d14cf2d55537612b00d2ba3b7e18c034c46f38d574a8a1b4bdd5a1b962786d0b20ff7b6205ed476f30cd91aee13fc35bb221804d09c05182038536d844e03533861fcaa3f257ebd4ba8b8ac4b12897524aece"
	ds, _ := AesDecrypt(dfd, []byte(key))
	fmt.Println(string(ds))

}

func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	fmt.Println(blockSize,"sdfsadfsdf")

	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

