package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AesCbcEncrypt AES_CBC 算法加密
func AesCbcEncrypt(rawData, key, iv []byte) ([]byte, error) {
	keyLen := len(key)
	if keyLen < 16 || keyLen%16 != 0 { // 16位：aes-128；32-位：aes-256
		return nil, errors.New("invalid key length")
	}
	// 设置默认的初始化向量
	if len(iv) == 0 {
		iv = key[:aes.BlockSize]
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充
	rawData = Pkcs7Padding(rawData, keyLen)

	mode := cipher.NewCBCEncrypter(block, iv)
	encryptedData := make([]byte, len(rawData))

	mode.CryptBlocks(encryptedData, rawData)

	return encryptedData, nil
}

// AesCbcDecrypt AES_CBC 算法解密
func AesCbcDecrypt(encryptedData, key, iv []byte) ([]byte, error) {
	keyLen := len(key)
	if keyLen < 16 || keyLen%16 != 0 { // 16位：aes-128；32-位：aes-256
		return nil, errors.New("invalid key length")
	}
	// 设置默认的初始化向量
	if len(iv) == 0 {
		iv = key[:aes.BlockSize]
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decryptedData := make([]byte, len(encryptedData))

	mode.CryptBlocks(decryptedData, encryptedData)

	// 去除填充
	decryptedData = Pkcs7UnPadding(decryptedData)

	return decryptedData, nil
}

// Pkcs7Padding PKCS#7 填充
func Pkcs7Padding(text []byte, blockSize int) []byte {
	// 计算待填充的长度
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}

	return append(text, paddingText...)
}

// Pkcs7UnPadding PKCS#7 去除填充
func Pkcs7UnPadding(text []byte) []byte {
	// 取出填充的数据 以此来获得填充数据长度
	unPadding := int(text[len(text)-1])

	return text[:(len(text) - unPadding)]
}
