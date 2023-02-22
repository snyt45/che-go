package yaml

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteLoadYaml(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "test_cheat*.yml")
	if err != nil {
		t.Fatal(err)
	}
	testDataPath := tempFile.Name()
	defer os.Remove(testDataPath)

	cheatYaml := &CheatYaml{
		Commands: []Command{
			{
				ID:          1,
				Command:     "ls",
				Tags:        []string{"utility", "filesystem"},
				Description: "List directory contents",
				Recommend:   5,
				Favorite:    true,
				Details:     "ls is a command to list computer files in Unix and Unix-like operating systems.\nls is specified by POSIX and the Single UNIX Specification.",
			},
		},
	}
	err = WriteYaml(testDataPath, cheatYaml)
	if err != nil {
		t.Errorf("WriteYaml returned an unexpected error: %v", err)
	}

	loadedYaml, err := LoadYaml(testDataPath)
	if err != nil {
		t.Errorf("LoadYaml returned an unexpected error: %v", err)
	}

	if len(loadedYaml.Commands) != 1 {
		t.Errorf("Unexpected number of commands in CheatYaml: %v", len(loadedYaml.Commands))
	}

	if loadedYaml.Commands[0].Command != "ls" {
		t.Errorf("Unexpected command in CheatYaml: %v", loadedYaml.Commands[0])
	}
}

func TestCreateYaml(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir) // テスト終了時に一時ディレクトリを削除する

	filePath := filepath.Join(dir, "cheat.yml")
	err = CreateYaml(filePath)
	if err != nil {
		t.Errorf("CreateYaml returned an unexpected error: %v", err)
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	content := []byte("commands:\n")
	if string(b) != string(content) {
		t.Fatalf("content mismatch: expected=%#v, actual=%#v", string(content), string(b))
	}
}
