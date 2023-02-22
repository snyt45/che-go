package cmd

import (
	"fmt"

	"github.com/snyt45/che-go/internal/helper"
	"github.com/snyt45/che-go/internal/yaml"
	"github.com/urfave/cli/v2"
)

// RemoveCommandは、対話形式で削除するコマンドを選択して削除します。
// 削除した内容は~/cheat.ymlから削除されます。
func RemoveCommand(c *cli.Context) error {
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
	label := "Delete command"
	selectedIndex, err := helper.SelectCommand(label, commands)
	if err != nil {
		return err
	}

	// Remove selected command
	selectedID := commands[selectedIndex].ID
	selectedCommand := commands[selectedIndex].Command
	for i, c := range commands {
		if selectedID == c.ID {
			// 削除する要素の1個前までのスライスを作成、削除する要素の1個先から最後の要素までのスライスを作成し追加する
			// ex) [cmd1, cmd2, cmd3, cmd4, cmd5] の cmd3を削除する場合、
			//     [cmd1, cmd2]のスライスに、[cmd4, cmd5]のスライスをappend
			cheatYaml.Commands = append(cheatYaml.Commands[:i], cheatYaml.Commands[i+1:]...)
			err = yaml.WriteYaml(path, cheatYaml)
			if err != nil {
				return err
			}
			break
		}
	}
	fmt.Printf("Command removed successfully. '%s'", selectedCommand)
	return nil
}
