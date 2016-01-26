package minified_diff

func MinifiedDiff(file1 string, file2 string) *string {
	result := file1 + " " + file2
	return &result
}
