package extract

import (
	"log"
	"os"
	"regexp"
	"strings"
)

// Contains one code block
type SourceBlock struct {
	T, Content string
}

// Extract a source block from a certain block in st
func ParseBlock(b string) (block SourceBlock) {
	// Trim the backquotes and newlines
	b = strings.Trim(b, "` ")

	var t, content string
	hasType := true

	// If the string starts with a new line, then is has no type
	if b[0] == '\n' {
		t = ""
		hasType = false
	}

	// Strip the outer new line
	b = strings.Trim(b, "\n")
	
	lines := strings.Split(b, "\n")
	if hasType {
		t = lines[0] // The first line is the type of block
		lines = lines[1:]
	}
	content = strings.Join(lines, "\n")
	block = SourceBlock{t, content}
	return

}

// Given a file, exctract all source blocks
func ExtractBlocks(f []byte) (blocks []SourceBlock, err error) {
	re := regexp.MustCompile("```[\\w\\W]*?```")
	var s SourceBlock
	for _, m := range re.FindAll(f, -1) {
		cleanF := strings.ReplaceAll(string(m), "\r", "")
		s = ParseBlock(cleanF)
		blocks = append(blocks, s)
	}
	return
}

// Given a list of blocks, concatenate the blocks of the same type
func ConcatenateBlocks(blocks []SourceBlock) (concatenatedBlocks []SourceBlock) {
	return blocks
}

// Extracts all code blocks from a given file
func ExtractBlocksFromFile(filePath string, joinBlocks bool) (blocks []SourceBlock, err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Could not read file %s", filePath)
	}
	blocks, err = ExtractBlocks(data)
	if err != nil {
		log.Printf("Failed to extract blocks on file %s", filePath)
	}

	if joinBlocks{
		blocks = ConcatenateBlocks(blocks)
	}

	return
}
