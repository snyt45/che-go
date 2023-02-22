<div align="center">
  <h1>che-go</h1>
  <h4>コマンドラインチートシート管理ツール</h4>
  <p><b>実験的なバージョンのため仕様が大きく変わる可能性があります。</b></p>
  <a href="https://asciinema.org/a/R1WiEWGkUDq7IT0Wj7zn31mio" target="_blank"><img src="https://asciinema.org/a/R1WiEWGkUDq7IT0Wj7zn31mio.svg" width="640" height="480" /></a>
</div>

che-goは、コマンドライン上で簡単にチートシートを管理できるツールです。 YAML形式のファイルに記述されたコマンドやコマンドのオプション、説明などを管理できます。

## インストール方法

che-goコマンドを以下の方法でインストールしてください。

```
$ go install github.com/snyt45/che-go@latest
```

che-goコマンドがインストールされていることを確認してください。

```
$ che-go --help
```

## 使い方

che-goコマンド実行時に`~/cheat.yml`が存在しない場合、`~/cheat.yml`を作成します。  
che-goコマンドでは`cheat.yml`を便利に管理するコマンドラインインターフェイスを提供します。  

### コマンドを追加する

対話形式でコマンドを追加できます。

```
$ che-go add
```

### コマンドを編集する

対話形式でコマンドを追加できます。

```
$ che-go edit
```

### コマンド一覧を表示する

現在登録されているコマンド一覧を表示します。

```
$ che-go list
```

### コマンドを削除する

コマンドを削除できます。

```
$ che-go remove
```

## 注意
promptuiの既出のバグで入力が端末の長さを超えて入力を行うと再レンダリングが発生します。  
`che-go add`、`che-go edit`で発生する可能性があります。  
その場合は直接`cheat.yml`を修正してください。  

https://github.com/manifoldco/promptui/issues/207  
https://github.com/manifoldco/promptui/issues/92

## ライセンス
[MIT](https://github.com/snyt45/che-go/blob/main/LICENSE)

## 関連プロジェクト
[che](https://github.com/snyt45/che) – The predecessor of che-go.
