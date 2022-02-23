package tools

func FileValidator(extension string) bool {
	if extension != ".json" && extension != "cvl" {
		return false
	}
	return true
}
