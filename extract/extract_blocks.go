package extract

import (
	"fmt"
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

// Given a file, extract all source blocks
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

// Given a list of source blocks, concatenate all blocks of the same type
func ConcatenateBlocks(blocks []SourceBlock) (concatenatedBlocks []SourceBlock) {
	blockTypeMap := make(map[string]SourceBlock)

	for _, b := range blocks {
		if v, ok := blockTypeMap[b.T]; ok {
			v.Content += "\n" + b.Content
			blockTypeMap[b.T] = v // Assign the updated value back to the map
		} else {
			blockTypeMap[b.T] = b
		}
	}

	// Create an array of the map values
	concatenatedBlocks = make([]SourceBlock, 0, len(blockTypeMap))
	for _, v := range blockTypeMap {
		concatenatedBlocks = append(concatenatedBlocks, v)
	}

	return
}


// Extracts all code blocks from a given file
func ExtractBlocksFromFile(filePath string, joinBlocks bool) (blocks []SourceBlock, err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Could not read file %s", filePath)
	}
	blocks, err = ExtractBlocks(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to extract blocks on file %s", filePath)
	}

	//TODO: This parse if very badly implemented
	if joinBlocks{
		log.Print("Concatenating blocks.")
		blocks = ConcatenateBlocks(blocks)
	}

	return blocks, nil
}
