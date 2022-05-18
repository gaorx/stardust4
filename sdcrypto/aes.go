package sdcrypto

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/gaorx/stardust4/sderr"
)

var (
	AES Encrypter = &EncrypterFunc{
		Encrypter: AESEncrypt,
		Decrypter: AESDecrypt,
	}
	AESCRC32 Encrypter = &CRC32Encrypter{AES}
)

func AESEncrypt(key, data []byte) ([]byte, error) {
	return AESEncryptPadding(key, data, Pkcs5)
}

func AESDecrypt(key, encrypted []byte) ([]byte, error) {
	return AESDecryptPadding(key, encrypted, UnPkcs5)
}

func AESEncryptPadding(key, data []byte, p Padding) ([]byte, error) {
	if p == nil {
		return nil, sderr.New("sdcrypto aes nil padding")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcrypto aes create cipher error")
	}
	data, err = p(data, block.BlockSize())
	if err != nil {
		return nil, sderr.Wrap(err, "sdcrypto aes padding error")
	}
	encrypter := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	encrypted := make([]byte, len(data))
	encrypter.CryptBlocks(encrypted, data)
	return encrypted, nil
}

func AESDecryptPadding(key, encrypted []byte, p Unpadding) ([]byte, error) {
	if p == nil {
		return nil, sderr.New("sdcrypto unaes nil unpadding")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcrypto deaes create cipher error")
	}
	decrypter := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	data := make([]byte, len(encrypted))
	decrypter.CryptBlocks(data, encrypted)
	r, err := p(data, block.BlockSize())
	if err != nil {
		return nil, sderr.Wrap(err, "sdcrypto deaes unpadding error")
	}
	return r, nil
}
