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
	AES_CRC32 Encrypter = &CRC32Encrypter{AES}
)

func AESEncrypt(key, data []byte) ([]byte, error) {
	return AESEncryptPadding(key, data, Pkcs5)
}

func AESDecrypt(key, crypted []byte) ([]byte, error) {
	return AESDecryptPadding(key, crypted, UnPkcs5)
}

func AESEncryptPadding(key, data []byte, p Padding) ([]byte, error) {
	if p == nil {
		return nil, sderr.New("nil padding")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcrypto aes: encrypt error")
	}
	data, err = p(data, block.BlockSize())
	if err != nil {
		return nil, sderr.Wrap(err, "sdcrypto aes: padding error")
	}
	encrypter := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	crypted := make([]byte, len(data))
	encrypter.CryptBlocks(crypted, data)
	return crypted, nil
}

func AESDecryptPadding(key, crypted []byte, p Unpadding) ([]byte, error) {
	if p == nil {
		return nil, sderr.New("nil unpadding")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, sderr.WithStack(err)
	}
	decrypter := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	data := make([]byte, len(crypted))
	decrypter.CryptBlocks(data, crypted)
	r, err := p(data, block.BlockSize())
	if err != nil {
		return nil, sderr.Wrap(err, "sdcrypto aes: unpadding error")
	}
	return r, nil
}
