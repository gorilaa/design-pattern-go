package main

import (
	"fmt"

	"design-pattern/pattern/creational/prototype/node"
)

func main() {
	fileOne := &node.File{Name: "File One"}
	fileTwo := &node.File{Name: "File Two"}
	fileThree := &node.File{Name: "File Three"}

	folderOne := &node.Folder{
		Name:     "Folder One",
		Children: []node.INode{fileOne},
	}

	folderTwo := &node.Folder{
		Name:     "Folder Two",
		Children: []node.INode{folderOne, fileTwo, fileThree},
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folderTwo.Print("  ")

	cloneFolder := folderTwo.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
