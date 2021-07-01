package dstorage

import (
	"crypto/sha1"
	"encoding/base64"
)

func splitFile(file []byte, noParts int) [][]byte {
	chunks := make([][]byte, 0)
	chunkSize := (len(file) + noParts - 1) / noParts
	for i := 0; i < len(file); i += chunkSize {
		end := i + chunkSize

		if end > len(file) {
			end = len(file)
		}

		chunks = append(chunks, file[i:end])
	}
	return chunks
}

func mergeChunks(chunks *map[int][]byte) []byte {
	file := make([]byte, 0)
	// карты отсортированы поэтому просто цикл
	for _, v := range *chunks {
		file = append(file, v...)
	}

	return file
}

func calculateHash(hh []byte) string {
	hasher := sha1.New()
	hasher.Write(hh)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func checkChunks(chunks map[int][]byte, hashes []string) bool {
	for i, v := range chunks {
		if hashes[i] != calculateHash(v) {
			return false
		}
	}
	return true
}
