package main

import (
	"os"
	"log"
	"regexp"
	"strings"
)

const BlockRegex = "```.*```"

// Contains one code block
type SourceBlock struct {
	t, content string
}

// Extract a source block from a certain block in st
func ParseBlock(b string) (block SourceBlock) {
	// Trim the backquotes
	b = strings.Trim(b, "`\n")

	// Get the block type from the first line
	lines := strings.Split(b, "\n")
	t := lines[0]
	content := strings.Join(lines[1:], "\n")
	block = SourceBlock{t, content}
	return
	
}

// Given a file, extract all source blocks
func ExtractBlocks(f []byte) (blocks []SourceBlock, err error) {
	re := regexp.MustCompile(BlockRegex)
	var s SourceBlock
	for _, m := range re.FindAll(f, -1) {
		log.Println(m)
		s = SourceBlock{t: "asda", content: "sdad"}
		blocks = append(blocks, s)
	}
	return
}

// Extracts all code blocks from a given file
func ExtractBlocksFromFile(filePath string) (blocks []SourceBlock, err error) {
	data, err := os.ReadFile(filePath)
	blocks = append(blocks, SourceBlock{t: "asda", content: string(data)})
	return
}
