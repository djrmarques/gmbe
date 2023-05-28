package main

import (
	"path/filepath"
	"testing"
)

func TestExtractBlocks(t *testing.T) {
	var result [2]SourceBlock
	result[0] = SourceBlock{t: "teser", content: "asd"}
	result[1] = SourceBlock{t: "teser", content: "asd"}

	path := filepath.Join(".", "test", "fixtures", "test.md")
	blocks, err := ExtractBlocks(path)
	if err != nil {
		t.Error("Found error")
	}

	for i := range result {
		if result[i] != blocks[i] {
			t.Error("Results do not match")			
		}
	}
}

q
