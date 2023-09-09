package cmd

import (
	"bytes"
	"os"
	"testing"
)

// 開始時のos.Argsを保持
var orgArgs = os.Args

func AfterTest() {
	// os.Argsを初期化してテストケースの独立性を保持
	os.Args = orgArgs
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string   // テストケースの名前
		args []string // テストケースで実行するコマンドライン引数
		want string   // 期待されるコマンドの出力
	}{
		{
			name: "TestCatOnce",
			args: []string{"cat"},
			want: "ฅ^•ω•^ฅ にゃー\n",
		},
		{
			name: "TestCatTwice",
			args: []string{"cat", "--count", "2"},
			want: "ฅ^•ω•^ฅ にゃー\nฅ^•ω•^ฅ にゃー\n",
		},
		{
			name: "TestDogOnce",
			args: []string{"dog"},
			want: "꒰ ՞•ﻌ•՞ ꒱ わん！\n",
		},
		{
			name: "TestDogTwice",
			args: []string{"dog", "--count", "2"},
			want: "꒰ ՞•ﻌ•՞ ꒱ わん！\n꒰ ՞•ﻌ•՞ ꒱ わん！\n",
		},
	}

	// テストケースを反復処理
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cmdOutput bytes.Buffer // コマンドの出力をキャプチャするバッファを作成
			rootCmd.SetOut(&cmdOutput) // Cobraのコマンド出力をバッファに設定

			// テストケースの引数をos.Argsに追加
			for _, arg := range tt.args {
				os.Args = append(os.Args, arg)
			}

			// Execute() 関数を呼び出し、コマンドを実行
			Execute()

			// コマンドの出力と期待される出力を比較し、テスト結果を検証
			if cmdOutput.String() != tt.want {
				t.Errorf("Execute() = \n%v\n, want\n%v\n", cmdOutput.String(), tt.want)
			}

			// テストケースの後、os.Argsを元に戻す
			AfterTest()
		})
	}
}
