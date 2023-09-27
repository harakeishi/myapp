package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDogCommand(t *testing.T) {
	tests := []struct {
		name    string // テストケースの名前
		command string // テストケースで実行するコマンド
		count   int    // テストケースでcountフラグに設定する値
		want    string // 期待されるコマンドの出力
	}{
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

			// countフラグにテストケースで設定した値を設定
			dogCmd.Flag("count").Value.Set(fmt.Sprint(tt.count))
			// コマンド実行
			dogCmd.RunE(dogCmd, []string{})

			// コマンドの出力と期待される出力を比較し、テスト結果を検証
			if cmdOutput.String() != tt.want {
				t.Errorf("Execute() = \n%v\n, want\n%v\n", cmdOutput.String(), tt.want)
			}
		})
	}
}
