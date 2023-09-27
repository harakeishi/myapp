package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	tests := []struct {
		name    string // テストケースの名前
		command string // テストケースで実行するコマンド
		count   int    // テストケースでcountフラグに設定する値
		want    string // 期待されるコマンドの出力
	}{
		{
			name:    "TestCatOnce",
			command: "cat",
			count:   1,
			want:    "ฅ^•ω•^ฅ にゃー\n",
		},
		{
			name:    "TestCatTwice",
			command: "cat",
			count:   2,
			want:    "ฅ^•ω•^ฅ にゃー\nฅ^•ω•^ฅ にゃー\n",
		},
		{
			name:    "TestDogOnce",
			command: "dog",
			count:   1,
			want:    "꒰ ՞•ﻌ•՞ ꒱ わん！\n",
		},
		{
			name:    "TestDogTwice",
			command: "dog",
			count:   2,
			want:    "꒰ ՞•ﻌ•՞ ꒱ わん！\n꒰ ՞•ﻌ•՞ ꒱ わん！\n",
		},
	}

	// テストケースを反復処理
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cmdOutput bytes.Buffer // コマンドの出力をキャプチャするバッファを作成
			rootCmd.SetOut(&cmdOutput) // Cobraのコマンド出力をバッファに設定

			// Commands() 関数を呼び出し、サブコマンドを取得
			commands := rootCmd.Commands()
			for _, cmd := range commands {
				// テストケースのコマンドと一致するサブコマンドを実行
				if cmd.Name() == tt.command {
					// テストケースのコマンドライン引数を設定し、コマンドを実行
					cmd.Flag("count").Value.Set(fmt.Sprint(tt.count))
					cmd.RunE(cmd, []string{})
				}
			}

			// コマンドの出力と期待される出力を比較し、テスト結果を検証
			if cmdOutput.String() != tt.want {
				t.Errorf("Execute() = \n%v\n, want\n%v\n", cmdOutput.String(), tt.want)
			}
		})
	}
}
