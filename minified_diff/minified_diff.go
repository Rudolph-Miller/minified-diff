package minified_diff

import (
	"github.com/ditashi/jsbeautifier-go/jsbeautifier"
)

func MinifiedDiff(file1 string, file2 string) *string {
	options := jsbeautifier.DefaultOptions()
	result1 := jsbeautifier.BeautifyFile(file1, options)
	result2 := jsbeautifier.BeautifyFile(file2, options)
	result := *result1 + " " + *result2
	return &result
}
