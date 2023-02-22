package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/snyt45/che-go/internal/helper"
	"github.com/snyt45/che-go/internal/yaml"
	"github.com/urfave/cli/v2"
)

// EditCommandは、対話形式で編集するコマンドを選択して編集します。
// 入力した内容は~/cheat.ymlに保存されます。
// 受け付けるコマンド情報は次の通りです。
//   Command:     文字列
//   Tags:        カンマ区切りの文字列
//   Description: 文字列
//   Recommend:   1 ～ 5の数字
//   Favorite:    bool
//   Details:     複数行の文字列
func EditCommand(c *cli.Context) error {
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
	label := "Edit command"
	selectedIndex, err := helper.SelectCommand(label, commands)
	if err != nil {
		return err
	}
	selectedCommand := commands[selectedIndex]

	// Command(Prompt)
	validate := func(input string) error {
		// コマンドは必須です。
		if input == "" {
			return fmt.Errorf("command is required")
		}
		// 既に存在するコマンドは入力不可です。
		exists := checkCommandExistsExcludeSelf(selectedCommand.Command, cheatYaml, strings.TrimSpace(input))
		if exists {
			return fmt.Errorf("command '%s' already exists in cheat sheet", input)
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Command",
		Validate: validate,
		Default:  selectedCommand.Command,
	}
	commandStr, err := prompt.Run()
	if err != nil {
		return err
	}
	command := strings.TrimSpace(commandStr)

	// Tags(Prompt)
	prompt = promptui.Prompt{
		Label:   "Tags(Comma Separated Input)",
		Default: strings.Join(selectedCommand.Tags, ","),
	}
	tagsStr, err := prompt.Run()
	if err != nil {
		return err
	}
	tags := strings.Split(strings.ReplaceAll(tagsStr, " ", ""), ",")

	// Description(Prompt)
	prompt = promptui.Prompt{
		Label:   "Description",
		Default: selectedCommand.Description,
	}
	descriptionStr, err := prompt.Run()
	if err != nil {
		return err
	}
	description := strings.TrimSpace(descriptionStr)

	// Recommend(Select)
	recommends := []string{"1", "2", "3", "4", "5"}
	promptRecommend := promptui.Select{
		Label:     "Recommend",
		Items:     recommends,
		CursorPos: selectedCommand.Recommend - 1,
	}
	_, recommendStr, err := promptRecommend.Run()
	if err != nil {
		return err
	}
	recommend, _ := strconv.Atoi(recommendStr)

	// Favorite(Select)
	var selecttedFavorite int
	if selectedCommand.Favorite == true {
		selecttedFavorite = 0
	} else {
		selecttedFavorite = 1
	}
	favorites := []string{"true", "false"}
	promptFavorite := promptui.Select{
		Label:     "Favorite",
		Items:     favorites,
		CursorPos: selecttedFavorite,
	}
	_, favoriteStr, err := promptFavorite.Run()
	if err != nil {
		return err
	}
	favorite, _ := strconv.ParseBool(favoriteStr)

	// Details(Prompt)
	prompt = promptui.Prompt{
		Label: "Details(multiline(end with 'done'))",
	}

	// Detailsは1行ずつ入力を受け付けます。
	// 'done'と入力することで入力を完了します。
	var details []string
	for {
		result, err := prompt.Run()
		if err != nil {
			return err
		}
		if strings.TrimSpace(result) == "done" {
			break
		}
		details = append(details, result)
	}

	cmd := yaml.Command{
		ID:          selectedCommand.ID,
		Command:     command,
		Tags:        tags,
		Description: description,
		Recommend:   recommend,
		Favorite:    favorite,
		Details:     strings.Join(details, "\n"),
	}

	// コマンド更新
	cheatYaml.Commands[selectedIndex] = cmd
	err = yaml.WriteYaml(path, cheatYaml)
	if err != nil {
		return err
	}

	// 処理が成功した旨と入力内容をコンソールに出力します。
	fmt.Printf(
		"Command edited successfully!\n  ID: %d\n  Command: %s\n  Tags: %s\n  Description: %s\n  Recommend: %d\n  Favorite: %t\n  Details: \n    %s\n",
		selectedCommand.ID, command, tags, description, recommend, favorite, strings.Join(details, "\n    "),
	)

	return nil
}

// checkCommandExistsExcludeSelfは、既にコマンドが登録されていないかを確認します。
func checkCommandExistsExcludeSelf(selfCommand string, cheatYaml *yaml.CheatYaml, command string) bool {
	for _, c := range cheatYaml.Commands {
		// 自身のコマンドの場合はskip
		if c.Command == selfCommand {
			continue
		}
		if c.Command == command {
			return true
		}
	}
	return false
}
