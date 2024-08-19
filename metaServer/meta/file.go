package meta

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"time"
)

type Chunk struct {
	Id     string
	Len    int
	hash   []byte
	server *DataServer
}

type Status struct {
	CreationTime time.Time
	LastAccess   time.Time
	LastModified time.Time
	Permissions  string
	Owner        User
	Group        []*UserGroup
}

type File struct {
	Name   string
	Size   int
	Hash   []byte
	Chunks []*Chunk
	Status Status
}

// CalculateSHA256 计算文件的SHA-256哈希值
func CalculateSHA256(filePath string) ([]byte, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建一个新的SHA-256哈希对象
	hash := sha256.New()

	// 从文件中读取并更新哈希值
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	// 获取最终的哈希值
	return hash.Sum(nil), nil
}

// VerifyFileIntegrity 比较文件的哈希值以验证其完整性
func VerifyFileIntegrity(filePath string, originalHash []byte) (bool, error) {
	// 计算文件的当前哈希值
	currentHash, err := CalculateSHA256(filePath)
	if err != nil {
		return false, err
	}

	// 比较原始哈希值和当前哈希值
	return fmt.Sprintf("%x", currentHash) == fmt.Sprintf("%x", originalHash), nil
}
