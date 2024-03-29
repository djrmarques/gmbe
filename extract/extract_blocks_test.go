package extract

import (
	"fmt"
	"path/filepath"
	"testing"
)

const TestString1 = "Text Here\n```yml\nSomeText\nMoreText\n```\nMore Text here"
const TestString2 = "Text Here\n```\nSomeText\nMoreText\n```\nMore Text here"
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
		tErrM = fmt.Sprintf("Block Type not Equal:\n%s\n!=\n%s", b1.T, b2.T)
	}

	if !cEqual {
		cErrM = fmt.Sprintf("Block Content not equal:\n%s\n!=\n%s", b1.Content, b2.Content)
	}

	err = fmt.Errorf(tErrM + cErrM)
	return
}


func TestExtractBlocksFromStr1(t *testing.T) {
	result := SourceBlock{T: "yml", Content: "SomeText\nMoreText"}
	blocks, _ := ExtractBlocks([]byte(TestString1))
	if len(blocks) != 1 {
		t.Errorf("Expected 1 blocks, but found %d", len(blocks))
	}
	be, err := blockIsEqual(result, blocks[0])
	if !be {
		t.Error(err)
	}

}

func TestExtractBlocksFromStr2(t *testing.T) {
	result := SourceBlock{T: "", Content: "SomeText\nMoreText"}
	blocks, _ := ExtractBlocks([]byte(TestString2))
	if len(blocks) != 1 {
		t.Errorf("Expected 1 blocks, but found %d", len(blocks))
	}
	be, err := blockIsEqual(result, blocks[0])
	if !be {
		t.Error(err)
	}

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
	result[0] = SourceBlock{T: "python", Content: "Python code line 1\nPython code line 2"}
	result[1] = SourceBlock{T: "yaml", Content: "something:\n    - here\n    - here"}
	result[2] = SourceBlock{T: "", Content: "Unknown Format"}
	result[3] = SourceBlock{T: "python", Content: "Python code line 3\nPython code line 4"}
	
	path := filepath.Join("..", "test", "fixtures", "test.md")
	blocks, err := ExtractBlocksFromFile(path, false)
	if err != nil {
		t.Error("Found error")
	}

	if n_blocks := len(blocks); n_blocks != 4 {
		t.Errorf("Expected 4 blocks, but got %d", n_blocks)
	}

	var b SourceBlock
	for i, r := range result {
		b = blocks[i]
		be, err := blockIsEqual(r, b)
		if !be {
			t.Error(err)
			fmt.Printf("\n%q\n%q\n", r.Content, b.Content)
		}
	}
}

func TestConcatenateBlocks(t *testing.T) {
	var result [3]SourceBlock
	result[0] = SourceBlock{T: "python", Content: "pythoncode1"}
	result[1] = SourceBlock{T: "yaml", Content: "something:\n - here\n - here"}
	result[2] = SourceBlock{T: "python", Content: " pythoncode2"}

	concatenatedResults := ConcatenateBlocks(result[:])
	if nRes := len(concatenatedResults); nRes != 2 {
		t.Errorf("Length must be 2 but got %d", nRes)
	}
	

}

func TestExtractBlocksFromFileJoined(t *testing.T) {
	var result [3]SourceBlock
	result[0] = SourceBlock{T: "python", Content: "Python code line 1\nPython code line 2\nPython code line 3\nPython code line 4"}
	result[1] = SourceBlock{T: "yaml", Content: "something:\n    - here\n    - here"}
	result[2] = SourceBlock{T: "", Content: "Unknown Format"}
	
	path := filepath.Join("..", "test", "fixtures", "test.md")
	blocks, err := ExtractBlocksFromFile(path, true)
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
