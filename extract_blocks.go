package main

import (
	"os"
)

// Contains one code block
type SourceBlock struct {
	t, content string
}

// Extracts all code blocks from a given file
func ExtractBlocks(filePath string) (blocks []SourceBlock, err error) {
	data, err := os.ReadFile(filePath)
	blocks = append(blocks, SourceBlock{t: "asda", content: string(data)})
	return
}
