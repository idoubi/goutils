package crypt

import (
	"crypto/aes"
	"crypto/cipher"
)

// AesGcmEncrypt AEAD_AES_GCM 算法加密
func AesGcmEncrypt(rawData, associatedData, nonce, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	encryptedData := aesgcm.Seal(nil, nonce, rawData, associatedData)

	return encryptedData, nil
}

// AesGcmDecrypt AEAD_AES_GCM 算法解密
func AesGcmDecrypt(encryptedData, associatedData, nonce, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	decryptedData, err := aesgcm.Open(nil, nonce, encryptedData, associatedData)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}
