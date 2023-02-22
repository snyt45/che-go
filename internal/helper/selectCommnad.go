package helper

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/snyt45/che-go/internal/yaml"
)

func SelectCommand(label string, commands []yaml.Command) (int, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   `> {{ if eq .Favorite true }}{{ "★" | yellow }}{{ else }}{{ " " }}{{ end }} {{ .Command | cyan }}`,
		Inactive: `  {{ if eq .Favorite true }}{{ "★" | yellow }}{{ else }}{{ " " }}{{ end }} {{ .Command | cyan }}`,
		Selected: "\u2713 {{ .Command | red | cyan }}",
		Details: `
--------- Command ----------
{{ "Command:" | faint }}      {{ .Command }}
{{ "Tags:" | faint }}         {{ .Tags }}
{{ "Description:" | faint }}  {{ .Description }}
{{ "Recommend:" | faint }}    {{ if eq .Recommend 1 }} {{ "✰✰✰✰★" | blue }} {{ else if eq .Recommend 2 }} {{ "✰✰✰★★" | blue }} {{ else if eq .Recommend 3 }} {{ "✰✰★★★" | blue }} {{ else if eq .Recommend 4 }} {{ "✰★★★★" | blue }}	{{ else if eq .Recommend 5 }} {{ "★★★★★" | blue }} {{ end }}
{{ "Favorite:" | faint }}     {{ if eq .Favorite true }} {{ "★" | yellow }} {{ else }} {{ "" }} {{ end }}
{{ "Details:" | faint }}      {{ .Details }} `,
	}

	// 各要素ごとにinputと一致するかの処理が走る
	// 一致した要素を検索結果として返す
	searcher := func(input string, index int) bool {
		command := commands[index]
		name := strings.Replace(strings.ToLower(command.Command), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:             label,
		Items:             commands,
		Templates:         templates,
		Size:              20, // 最大表示件数
		Searcher:          searcher,
		StartInSearchMode: true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return 0, fmt.Errorf("Prompt failed %v\n", err)
	}

	return i, nil
}
