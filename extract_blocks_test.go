package main

import (
	"testing"
)

const TestString1 = "Text Here \n```yml\nSomeText\nMoreText\n```\nMore Text here"
const TestString2 = "Text Here \n```\nSomeText\nMoreText\n```\nMore Text here"
const TestBlock = "```yml\nwtv\n```"

func TestParseBlock(t *testing.T) {
	block := ParseBlock(TestBlock)
	result := SourceBlock{t: "yml", content: "wtv"}
	if block != result {
		t.Fatalf("Blocks do not match. Expected %v but got %+v", result, block)
	}
		
}

// func TestExtractBlocksFromStr(t *testing.T) {
// }

// func TestExtractBlocksFromFile(t *testing.T) {
// 	var result [2]SourceBlock
// 	result[0] = SourceBlock{t: "teser", content: "asd"}
// 	result[1] = SourceBlock{t: "teser", content: "asd"}

// 	path := filepath.Join(".", "test", "fixtures", "test.md")
// 	blocks, err := ExtractBlocksFromFile(path)
// 	if err != nil {
// 		t.Error("Found error")
// 	}

// 	for i := range result {
// 		if result[i] != blocks[i] {
// 			t.Error("Results do not match")			
// 		}
// 	}
//}
