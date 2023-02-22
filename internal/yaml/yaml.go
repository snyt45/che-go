package yaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Command struct {
	ID          int      `yaml:"id"`
	Command     string   `yaml:"command"`
	Tags        []string `yaml:"tags"`
	Description string   `yaml:"description"`
	Recommend   int      `yaml:"recommend"`
	Favorite    bool     `yaml:"favorite"`
	Details     string   `yaml:"details"`
}

type CheatYaml struct {
	Commands []Command `yaml:"commands"`
}

// YAMLファイルを読み込み、YAMLファイルをCheatYaml構造体に変換
// YAMLファイルの情報を保持したCheatYaml構造体のポインタを返す
func LoadYaml(path string) (*CheatYaml, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("LoadYaml ReadFile error: %w", err)
	}

	var cheatYaml CheatYaml
	err = yaml.Unmarshal(content, &cheatYaml)
	if err != nil {
		return nil, fmt.Errorf("LoadYaml Unmarshal error: %w", err)
	}

	return &cheatYaml, nil
}

// CheatYaml構造体をYAMLフォーマットに変換し、YAMLファイルに書き込み
func WriteYaml(path string, cheatYaml *CheatYaml) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("WriteYaml ReadFile error: %w", err)
	}
	content, err = yaml.Marshal(cheatYaml)
	if err != nil {
		return fmt.Errorf("WriteYaml Marshal error: %w", err)
	}

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return fmt.Errorf("WriteYaml WriteFile error: %w", err)
	}

	return nil
}

func CreateYaml(path string) error {
	_, err := os.Stat(path)
	// CheatYamlが存在しない場合、作成する
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("CreateYaml error: %w", err)
		}
		defer f.Close()

		_, err = f.WriteString("commands:\n")
		if err != nil {
			return fmt.Errorf("CreateYaml error: %w", err)
		}
		if err != nil {
			return fmt.Errorf("CreateYaml error: %w", err)
		}
		// 想定しないエラーの場合
	} else if err != nil {
		return fmt.Errorf("CreateYaml error: %w", err)
	}

	// 既にCheatYamlが存在する場合、作成しない
	return nil
}

func YamlPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("YamlPath error: %w", err)
	}
	path := filepath.Join(home, "cheat.yml")
	return path, nil
}
