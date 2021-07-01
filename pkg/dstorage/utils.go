package dstorage

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
