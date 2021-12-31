package util

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// 作業ディレクトリを制御する関数たち（開始）
////////////////////////////////////////////////////////////////////////////////
type WorkingDirectory struct {
	BaseDir string
	Prefix  string
	workDir string
}

func (w *WorkingDirectory) clear() {
	w.workDir = ""
}

func (w *WorkingDirectory) baseDirExists() bool {
	if f, err := os.Stat(w.BaseDir); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

func (w *WorkingDirectory) Path() string {
	return w.workDir
}

func (w *WorkingDirectory) Create() (string, error) {
	if !w.baseDirExists() {
		return "", errors.New("the specified directory does not exist")
	}

	workdirName := strconv.FormatInt(time.Now().Unix(), 10)
	if w.Prefix != "" {
		workdirName = w.Prefix + "-" + workdirName
	}
	tmp := filepath.Join(w.BaseDir, workdirName)

	if err := os.Mkdir(tmp, 0777); err != nil {
		return "", err
	}

	w.workDir = tmp
	return w.workDir, nil
}

func (w *WorkingDirectory) Delete() error {
	defer w.clear()

	if w.workDir == "" {
		return errors.New("working directory does not exist")
	}

	return os.RemoveAll(w.workDir)
}
