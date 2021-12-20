package util

import (
	"os/exec"
	"strconv"
)

// テストファイルを生成する処理（ファイルの存在チェックはしないので、すでに存在する場合は強制上書きする）
func GenerateTestingFile(filePath string, size uint32) error {
	sizeStr := strconv.FormatUint(uint64(size), 10)

	// テストファイル生成
	cmd := exec.Command("dd", "if=/dev/zero", "of="+filePath, "bs="+sizeStr, "count=1")

	// 処理を実行
	if err := cmd.Start(); err != nil {
		return err
	}

	// 処理が終了するまで待つ
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
