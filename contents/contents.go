package contents

// Index returns contents.
func Index() []map[string]string {
	contents := []map[string]string{}

	contents = append(contents, map[string]string{
		"title": "Test1",
		"body":  "this is test content",
	})
	contents = append(contents, map[string]string{
		"title": "Test2",
		"body":  "this is test content",
	})

	return contents
}
