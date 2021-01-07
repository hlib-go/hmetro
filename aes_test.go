package hmetro

import (
	"encoding/base64"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestAes(t *testing.T) {
	origData := []byte("Hello World") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	log.Println("原文：", string(origData))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	log.Println("------------------ ECB模式 --------------------")
	encrypted = AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	log.Println("------------------ CFB模式 --------------------")
	encrypted = AesEncryptCFB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptCFB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}

func TestAesDecryptCBC(t *testing.T) {
	type args struct {
		encrypted []byte
		key       []byte
	}
	tests := []struct {
		name          string
		args          args
		wantDecrypted []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDecrypted := AesDecryptCBC(tt.args.encrypted, tt.args.key); !reflect.DeepEqual(gotDecrypted, tt.wantDecrypted) {
				t.Errorf("AesDecryptCBC() = %v, want %v", gotDecrypted, tt.wantDecrypted)
			}
		})
	}
}

func TestAesDecryptCFB(t *testing.T) {
	type args struct {
		encrypted []byte
		key       []byte
	}
	tests := []struct {
		name          string
		args          args
		wantDecrypted []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDecrypted := AesDecryptCFB(tt.args.encrypted, tt.args.key); !reflect.DeepEqual(gotDecrypted, tt.wantDecrypted) {
				t.Errorf("AesDecryptCFB() = %v, want %v", gotDecrypted, tt.wantDecrypted)
			}
		})
	}
}

func TestAesDecryptECB(t *testing.T) {
	type args struct {
		encrypted []byte
		key       []byte
	}
	tests := []struct {
		name          string
		args          args
		wantDecrypted []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDecrypted := AesDecryptECB(tt.args.encrypted, tt.args.key); !reflect.DeepEqual(gotDecrypted, tt.wantDecrypted) {
				t.Errorf("AesDecryptECB() = %v, want %v", gotDecrypted, tt.wantDecrypted)
			}
		})
	}
}

func TestAesEncryptCBC(t *testing.T) {
	type args struct {
		origData []byte
		key      []byte
	}
	tests := []struct {
		name          string
		args          args
		wantEncrypted []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEncrypted := AesEncryptCBC(tt.args.origData, tt.args.key); !reflect.DeepEqual(gotEncrypted, tt.wantEncrypted) {
				t.Errorf("AesEncryptCBC() = %v, want %v", gotEncrypted, tt.wantEncrypted)
			}
		})
	}
}

func TestAesEncryptCFB(t *testing.T) {
	type args struct {
		origData []byte
		key      []byte
	}
	tests := []struct {
		name          string
		args          args
		wantEncrypted []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEncrypted := AesEncryptCFB(tt.args.origData, tt.args.key); !reflect.DeepEqual(gotEncrypted, tt.wantEncrypted) {
				t.Errorf("AesEncryptCFB() = %v, want %v", gotEncrypted, tt.wantEncrypted)
			}
		})
	}
}

func TestAesEncryptECB(t *testing.T) {
	type args struct {
		origData []byte
		key      []byte
	}
	tests := []struct {
		name          string
		args          args
		wantEncrypted []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEncrypted := AesEncryptECB(tt.args.origData, tt.args.key); !reflect.DeepEqual(gotEncrypted, tt.wantEncrypted) {
				t.Errorf("AesEncryptECB() = %v, want %v", gotEncrypted, tt.wantEncrypted)
			}
		})
	}
}

func Test_pkcs5Padding(t *testing.T) {
	type args struct {
		ciphertext []byte
		blockSize  int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkcs5Padding(tt.args.ciphertext, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkcs5Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pkcs5UnPadding(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkcs5UnPadding(tt.args.origData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkcs5UnPadding() = %v, want %v", got, tt.want)
			}
		})
	}
}
