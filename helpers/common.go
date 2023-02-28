package helpers

import (
	"fmt"
	"time"

	filehelpers "github.com/vnhacker1337/goutils/helpers"

	funk "github.com/thoas/go-funk"
)

func Diff2Files(f1 string, f2 string) ([]string, []string) {

	// file1_slice := fileutil.LoadFile(f1)
	// file2_slice := fileutil.LoadFile(f2)

	file1_slice := filehelpers.LoadFile(f1)
	file2_slice := filehelpers.LoadFile(f2)

	only_file1, only_file2 := funk.DifferenceString(file1_slice, file2_slice)

	return only_file1, only_file2
}

func GetCurrentTime() string {
	currentTime := time.Now()
	stime := fmt.Sprintf(currentTime.Format("01-02-2006"))
	return stime
}
