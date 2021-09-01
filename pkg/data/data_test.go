package data

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_FOLDER = "tests"
var TEST_DATA = "testData"

func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func TestPathExists(t *testing.T) {

	type testCase struct {
		I string
		R bool
		E bool
	}

	var tests = []testCase{
		{TEST_DATA, true, false},
		{filepath.Join(TEST_DATA, "fileContentRef.txt"), true, false},
		{TEST_FOLDER, false, false},
		{"\000x", false, true}}

	for _, v := range tests {
		R, E := PathExists(v.I)
		assert.Equalf(t, v.R, R, "Results does not match on value %s", v.I)
		assert.Equalf(t, E != nil, v.E, "Errors does not match on value %s", v.I)
	}

}

func TestFolder(t *testing.T) {

	folderA := filepath.Join(TEST_FOLDER, "A")
	folderB := filepath.Join(TEST_FOLDER, "B")
	refFile := filepath.Join(TEST_DATA, "fileContentRef.txt")
	testFile := filepath.Join(folderA, "muststay.txt")

	os.MkdirAll(folderA, os.ModePerm)
	defer os.RemoveAll(TEST_FOLDER)
	err := copy(refFile, testFile)
	assert.NoErrorf(t, err, "Copying ref file to folder %s triggers an error", folderA)

	var foldersToCheck = []string{folderA, folderB}

	InitializeFolders(foldersToCheck)

	assert.DirExistsf(t, folderA, "Directory %s isn't conserved", folderA)
	assert.DirExistsf(t, folderB, "Directory %s isn't created", folderB)
	f1, err1 := ioutil.ReadFile(testFile)
	if err1 != nil {
		log.Fatal(err1)
	}
	f2, err2 := ioutil.ReadFile(refFile)
	if err2 != nil {
		log.Fatal(err2)
	}
	assert.Equalf(t, f2, f1, "Folder %s content is changed", folderA)
}
