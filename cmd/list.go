package cmd

import (
	"github.com/snyt45/che-go/internal/helper"
	"github.com/snyt45/che-go/internal/yaml"
	"github.com/urfave/cli/v2"
)

// ListCommandは、リスト形式でコマンド一覧を表示します。
// デフォルトではSearchモードで起動します。インタラクティブにコマンドを絞り込めます。
// また、モードは / で切り替えることができます。
func ListCommand(c *cli.Context) error {
	path, err := yaml.YamlPath()
	if err != nil {
		return err
	}
	cheatYaml, err := yaml.LoadYaml(path)
	if err != nil {
		return err
	}
	// コマンドを選択するプロンプト
	commands := cheatYaml.Commands
	label := "Select command"
	_, err = helper.SelectCommand(label, commands)
	if err != nil {
		return err
	}
	return nil
}
