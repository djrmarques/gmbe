package extract

import (
	"testing"
	"path/filepath"
	"reflect"
)

const TestString1 = "Text Here \n```yml\nSomeText\nMoreText\n```\nMore Text here"
const TestString2 = "Text Here \n```\nSomeText\nMoreText\n```\nMore Text here"


func TestExtractBlocksFromStr(t *testing.T) {
	blocks, _ := ExtractBlocks([]byte(TestString1))
	if len(blocks) != 1 {
		t.Errorf("Expected 1 blocks, but found %d", len(blocks))
	}
	
}

func TestExtractBlocksFromFile(t *testing.T) {
	var result [2]SourceBlock
	result[0] = SourceBlock{T: "python", Content: "Python code line 1\n Python code line 2"}
	result[1] = SourceBlock{T: "yaml", Content: "something:\n - here\n - here"}

	path := filepath.Join("..", "test", "fixtures", "test.md")
	blocks, err := ExtractBlocksFromFile(path)
	if err != nil {
		t.Error("Found error")
	}

	for i := range result {
		if reflect.DeepEqual(result[i],  blocks[i]) {
			t.Errorf("Results do not match. Expected: %+v but found %+v", result[i], blocks[i])			
		}
	}
}
