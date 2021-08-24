package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_FOLDER = "tests"
var TEST_DATA = "testData"

func writeFile(path string) {
	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func TestFolder(t *testing.T) {

	folderA := filepath.Join(TEST_FOLDER, "A")
	folderB := filepath.Join(TEST_FOLDER, "B")

	os.MkdirAll(folderA, os.ModePerm)
	writeFile(filepath.Join(folderA, "muststay.txt"))

	var foldersToCheck = []string{folderA, folderB}

	initializeFolders(foldersToCheck)

	assert.DirExistsf(t, folderA, "Directory %s isn't conserved", folderA)
	assert.DirExistsf(t, folderB, "Directory %s isn't created", folderB)
	f1, err1 := ioutil.ReadFile(filepath.Join(folderA, "muststay.txt"))
	if err1 != nil {
		log.Fatal(err1)
	}
	f2, err2 := ioutil.ReadFile(filepath.Join(TEST_DATA, "fileContentRef.txt"))
	if err2 != nil {
		log.Fatal(err2)
	}
	assert.Equalf(t, f2, f1, "Folder %s content is changed", folderA)
	os.RemoveAll(TEST_FOLDER)
}
