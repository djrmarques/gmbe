package extract

import (
	"testing"
	"path/filepath"
)


func TestExtractBlocksFromFileSingle(t *testing.T) {
	var result [4]SourceBlock
	result[0] = SourceBlock{T: "python", Content: "Python code line 1\n Python code line 2"}
	result[1] = SourceBlock{T: "yaml", Content: "something:\n - here\n - here"}
	result[2] = SourceBlock{T: "", Content: "Unknown Format"}
	result[3] = SourceBlock{T: "python", Content: "Python code line 3\n Python code line 4"}
	
	path := filepath.Join("..", "test", "fixtures", "test.md")
	blocks, err := ExtractBlocksFromFile(path)
	if err != nil {
		t.Error("Found error")
	}

	if n_blocks := len(blocks); n_blocks != 4 {
		t.Errorf("Expected 4 blocks, but got %d", n_blocks)
	}

	for i, r := range result {
		b := blocks[i]
		if b.T != r.T || b.Content != r.Content {
			t.Errorf("Results do not match. Expected: %+v but found %+v", result[i], blocks[i])			
		}
	}
}

func TestExtractBlocksFromFileJoined(t *testing.T) {
	var result [3]SourceBlock
	result[0] = SourceBlock{T: "python", Content: "Python code line 1\n Python code line 2\n Python code line 3\n Python code line 4"}
	result[1] = SourceBlock{T: "yaml", Content: "something:\n - here\n - here"}
	result[2] = SourceBlock{T: "", Content: "Unknown Format"}
	
	path := filepath.Join("..", "test", "fixtures", "test.md")
	blocks, err := ExtractBlocksFromFile(path)
	if err != nil {
		t.Error("Found error")
	}

	if n_blocks := len(blocks); n_blocks != 3 {
		t.Errorf("Expected 3 blocks, but got %d", n_blocks)
	}

	for i, r := range result {
		b := blocks[i]
		if b.T != r.T || b.Content != r.Content {
			t.Errorf("Results do not match. Expected: %+v but found %+v", result[i], blocks[i])			
		}
	}

}
