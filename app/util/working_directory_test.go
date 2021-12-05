package util_test

import (
	"os"
	"strings"
	"testing"

	util "github.com/hazuki3417/testing-file-generator/util"
)

func Test_Util_WorkingDirectory(t *testing.T) {

	workDir := util.WorkingDirectory{
		BaseDir: "/tmp",
	}

	t.Run("Create", func(t *testing.T) {
		str, err := workDir.Create()
		if err != nil {
			t.Errorf(err.Error())
		}

		if !strings.Contains(str, "/tmp/") {
			t.Errorf("Path() = %v expected %v", str, "/tmp/{timestamp}")
		}

		if _, err := os.Stat(str); err != nil {
			t.Errorf("directory generation failed")
		}
	})

	t.Run("Path", func(t *testing.T) {
		if str := workDir.Path(); !strings.Contains(str, "/tmp/") {
			t.Errorf("Path() = %v expected %v", str, "/tmp/{timestamp}")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		str := workDir.Path()
		if err := workDir.Delete(); err != nil {
			t.Errorf(err.Error())
		}
		if _, err := os.Stat(str); err == nil {
			t.Errorf("directory deletion failed")
		}
	})
}
