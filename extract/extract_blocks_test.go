package extract

import (
	"fmt"
	"path/filepath"
	"testing"
)

const TestBlock1 = "```yml\nwtv\n```"
const TestBlock2 = "```\nwtv\n```"

// Returns true if two source blocks are equal
// If not, err contains a description of the errors
func blockIsEqual(b1, b2 SourceBlock) (isEqual bool, err error) {
	tEqual := b1.T == b2.T
	cEqual := b1.Content == b2.Content
	isEqual = tEqual && cEqual
	var tErrM, cErrM string

	if !tEqual {
		tErrM = fmt.Sprintf("Block Type not Equal. %s != %s.", b1.T, b2.T)
	}

	if !cEqual {
		cErrM = fmt.Sprintf("Block content not equal. %s != %s.", b1.Content, b2.Content)
	}

	err = fmt.Errorf(tErrM + cErrM)
	return
}


// Tests if it is converting source blocks correctly
func TestParseBlock1(t *testing.T) {
	block := ParseBlock(TestBlock1)
	result := SourceBlock{T: "yml", Content: "wtv"}
	be, err := blockIsEqual(block, result)

	if !be {
		t.Error(err)
	}

}

func TestParseBlock2(t *testing.T) {
	block := ParseBlock(TestBlock2)
	result := SourceBlock{T: "", Content: "wtv"}
	be, err := blockIsEqual(block, result)

	if !be {
		t.Error(err)
	}
}

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
			t.Errorf("Results do not match. Expected:\n%+v\n but found \n%+v", result[i], blocks[i])			
		}
	}
}

// func TestExtractBlocksFromFileJoined(t *testing.T) {
// 	var result [3]SourceBlock
// 	result[0] = SourceBlock{T: "python", Content: "Python code line 1\n Python code line 2\n Python code line 3\n Python code line 4"}
// 	result[1] = SourceBlock{T: "yaml", Content: "something:\n - here\n - here"}
// 	result[2] = SourceBlock{T: "", Content: "Unknown Format"}
	
// 	path := filepath.Join("..", "test", "fixtures", "test.md")
// 	blocks, err := ExtractBlocksFromFile(path)
// 	if err != nil {
// 		t.Error("Found error")
// 	}

// 	if n_blocks := len(blocks); n_blocks != 3 {
// 		t.Errorf("Expected 3 blocks, but got %d", n_blocks)
// 	}

// 	for i, r := range result {
// 		b := blocks[i]
// 		if b.T != r.T || b.Content != r.Content {
// 			t.Errorf("Results do not match. Expected: %+v but found %+v", result[i], blocks[i])			
// 		}
// 	}

// }
