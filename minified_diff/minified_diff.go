package minified_diff

import (
	"fmt"
	"github.com/ditashi/jsbeautifier-go/jsbeautifier"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func beautify(src string) *string {
	options := jsbeautifier.DefaultOptions()
	return jsbeautifier.BeautifyFile(src, options)
}

func lineDiff(src1, src2 string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(src1, src2)
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	return result
}

func formatDiff(diff []diffmatchpatch.Diff) {
	fmt.Println(diff)
}

func MinifiedDiff(file1 string, file2 string) {
	src1 := beautify(file1)
	src2 := beautify(file2)
	formatDiff(lineDiff(*src1, *src2))
}
