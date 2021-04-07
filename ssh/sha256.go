package ssh

import (
	cryptoSha256 "crypto/sha256"
	"os"

	minioSha256 "github.com/minio/sha256-simd"
)

var (
	sha256New    = cryptoSha256.New
	sha256Sum256 = cryptoSha256.Sum256
)

func init() {
	if os.Getenv("SFTPGO_MINIO_SHA256_SIMD") == "1" {
		sha256New = minioSha256.New
		sha256Sum256 = minioSha256.Sum256
	}
}
