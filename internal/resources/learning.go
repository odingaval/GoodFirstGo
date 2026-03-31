package resources

// GetLearningResources returns curated learning URLs for a language
func GetLearningResources(lang string) []string {
	resources := map[string][]string{
		"go": {
			"https://go.dev/tour/welcome/1",
			"https://go.dev/doc/effective_go",
			"https://go.dev/doc/",
		},
		"rust": {
			"https://www.rust-lang.org/learn",
			"https://doc.rust-lang.org/book/",
		},
		"python": {
			"https://docs.python.org/3/tutorial/",
			"https://realpython.com/",
		},
		"javascript": {
			"https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide",
			"https://javascript.info/",
		},
		"java": {
			"https://docs.oracle.com/javase/tutorial/",
		},
		// Add more languages as needed
	}

	if res, ok := resources[lang]; ok {
		return res
	}
	return []string{"https://github.com/goodfirstissue"} // Fallback
}
