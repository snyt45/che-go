package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/snyt45/che-go/internal/yaml"
	"github.com/urfave/cli/v2"
)

// AddCommandは、対話形式で追加するコマンド情報を受け付ます。
// 入力した内容は~/cheat.ymlに保存されます。
// 受け付けるコマンド情報は次の通りです。
//   Command:     文字列
//   Tags:        カンマ区切りの文字列
//   Description: 文字列
//   Recommend:   1 ～ 5の数字
//   Favorite:    bool
//   Details:     複数行の文字列
func AddCommand(c *cli.Context) error {
	path, err := yaml.YamlPath()
	if err != nil {
		return err
	}
	cheatYaml, err := yaml.LoadYaml(path)
	if err != nil {
		return err
	}

	// Command(Prompt)
	validate := func(input string) error {
		// コマンドは必須です。
		if input == "" {
			return fmt.Errorf("command is required")
		}
		// 既に存在するコマンドは入力不可です。
		exists := checkCommandExists(cheatYaml, strings.TrimSpace(input))
		if exists {
			return fmt.Errorf("command '%s' already exists in cheat sheet", input)
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Command",
		Validate: validate,
	}
	commandStr, err := prompt.Run()
	if err != nil {
		return err
	}
	command := strings.TrimSpace(commandStr)

	// Tags(Prompt)
	prompt = promptui.Prompt{
		Label: "Tags(Comma Separated Input)",
	}
	tagsStr, err := prompt.Run()
	if err != nil {
		return err
	}
	tags := strings.Split(strings.ReplaceAll(tagsStr, " ", ""), ",")

	// Description(Prompt)
	prompt = promptui.Prompt{
		Label: "Description",
	}
	descriptionStr, err := prompt.Run()
	if err != nil {
		return err
	}
	description := strings.TrimSpace(descriptionStr)

	// Recommend(Select)
	recommends := []string{"1", "2", "3", "4", "5"}
	promptRecommend := promptui.Select{
		Label: "Recommend",
		Items: recommends,
	}
	_, recommendStr, err := promptRecommend.Run()
	if err != nil {
		return err
	}
	recommend, _ := strconv.Atoi(recommendStr)

	// Favorite(Select)
	favorites := []string{"true", "false"}
	promptFavorite := promptui.Select{
		Label: "Favorite",
		Items: favorites,
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
		Command:     command,
		Tags:        tags,
		Description: description,
		Recommend:   recommend,
		Favorite:    favorite,
		Details:     strings.Join(details, "\n"),
	}

	// コマンド追加
	cheatYaml = addCommandToCheatYaml(cheatYaml, cmd)
	err = yaml.WriteYaml(path, cheatYaml)
	if err != nil {
		return err
	}

	// 処理が成功した旨と入力内容をコンソールに出力します。
	fmt.Printf(
		"Command added successfully!\n  Command: %s\n  Tags: %s\n  Description: %s\n  Recommend: %d\n  Favorite: %t\n  Details: \n    %s\n",
		command, tags, description, recommend, favorite, strings.Join(details, "\n    "),
	)

	return nil
}

// addCommandToCheatYamlは、CheatYaml構造体に入力したCommandを追加します。
// 追加する際にCommand.IDは最大IDに+1した値で追加されます。
func addCommandToCheatYaml(cheatYaml *yaml.CheatYaml, command yaml.Command) *yaml.CheatYaml {
	// Generate a unique ID for the command
	var id int
	fmt.Println(len(cheatYaml.Commands))
	if len(cheatYaml.Commands) > 0 {
		id = maxID(cheatYaml.Commands) + 1
	} else {
		id = 1
	}
	command.ID = id

	cheatYaml.Commands = append(cheatYaml.Commands, command)

	return cheatYaml
}

// checkCommandExistsは、既にコマンドが登録されていないかを確認します。
func checkCommandExists(cheatYaml *yaml.CheatYaml, command string) bool {
	for _, c := range cheatYaml.Commands {
		if c.Command == command {
			return true
		}
	}
	return false
}

// maxIDは、CheatYaml構造体の中で最大IDを求めます。
func maxID(commnads []yaml.Command) int {
	ids := []int{}
	for _, c := range commnads {
		ids = append(ids, c.ID)
	}
	sort.Sort(sort.IntSlice(ids))
	return ids[len(ids)-1]
}
