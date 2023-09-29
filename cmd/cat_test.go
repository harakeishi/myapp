package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCatCommand(t *testing.T) {
	tests := []struct {
		name    string // テストケースの名前
		count   int    // テストケースでcountフラグに設定する値
		want    string // 期待されるコマンドの出力
	}{
		{
			name:    "TestCatOnce",
			count:   1,
			want:    "ฅ^•ω•^ฅ にゃー\n",
		},
		{
			name:    "TestCatTwice",
			count:   2,
			want:    "ฅ^•ω•^ฅ にゃー\nฅ^•ω•^ฅ にゃー\n",
		},
	}

	// テストケースを反復処理
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var cmdOutput bytes.Buffer // コマンドの出力をキャプチャするバッファを作成
			catCmd.SetOut(&cmdOutput) // Cobraのコマンド出力をバッファに設定

			// countフラグにテストケースで設定した値を設定
			catCmd.Flag("count").Value.Set(fmt.Sprint(tt.count))
			// コマンド実行
			catCmd.RunE(catCmd, []string{})

			// コマンドの出力と期待される出力を比較し、テスト結果を検証
			if cmdOutput.String() != tt.want {
				t.Errorf("Execute() = \n%v\n, want\n%v\n", cmdOutput.String(), tt.want)
			}
		})
	}
}
