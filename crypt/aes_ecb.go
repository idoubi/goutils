package crypt

import "crypto/aes"

// AesEcbEncrypt AES_ECB 算法加密
func AesEcbEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	// 填充
	rawData = Pkcs7Padding(rawData, blockSize)

	dataLen := len(rawData)
	encryptedData := make([]byte, dataLen)

	blockData := make([]byte, blockSize)

	for i := 0; i < dataLen; i += blockSize {
		block.Encrypt(blockData, rawData[i:i+blockSize])
		copy(encryptedData[i:i+blockSize], blockData)
	}

	return encryptedData, nil
}

// AesEcbDecrypt AES_ECB 算法解密
func AesEcbDecrypt(encryptedData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dataLen := len(encryptedData)
	decryptedData := make([]byte, dataLen)

	blockSize := block.BlockSize()
	blockData := make([]byte, blockSize)

	for i := 0; i < dataLen; i += blockSize {
		block.Decrypt(blockData, encryptedData[i:i+blockSize])
		copy(decryptedData[i:i+blockSize], blockData)
	}

	// 去除填充
	decryptedData = Pkcs7UnPadding(decryptedData)

	return decryptedData, nil
}
