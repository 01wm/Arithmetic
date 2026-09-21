package likou

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

const blockSize = 4096 // 块大小，可根据实际情况调整

func ComputeBlockHashes(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hasher := md5.New()
	var blockHashes []string

	buffer := make([]byte, blockSize)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		hasher.Reset()
		hasher.Write(buffer[:n])
		blockHash := hex.EncodeToString(hasher.Sum(nil))
		blockHashes = append(blockHashes, blockHash)
	}

	return blockHashes, nil
}
