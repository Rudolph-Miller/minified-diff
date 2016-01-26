package minified_diff

import (
	"github.com/ditashi/jsbeautifier-go/jsbeautifier"
	"github.com/sergi/go-diff/diffmatchpatch"
	"strings"
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

func prefix(c diffmatchpatch.Operation) string {
	switch c {
	case diffmatchpatch.DiffEqual:
		return " "
	case diffmatchpatch.DiffInsert:
		return "+"
	case diffmatchpatch.DiffDelete:
		return "-"
	}
	return " "
}

func formatDiff(lineDiff []diffmatchpatch.Diff) *string {
	result := ""
	for _, diff := range lineDiff {
		texts := strings.Split(diff.Text, "\n")
		for _, text := range texts {
			if len(text) > 0 {
				result += prefix(diff.Type)
				result += text
				result += "\n"
			}
		}
	}
	return &result
}

func MinifiedDiff(file1 string, file2 string) *string {
	src1 := beautify(file1)
	src2 := beautify(file2)
	return formatDiff(lineDiff(*src1, *src2))
}
