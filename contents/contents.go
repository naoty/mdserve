package contents

import (
	"os"
	"path/filepath"

	"github.com/gernest/front"
	"github.com/russross/blackfriday"
)

var contents = map[string]map[string]interface{}{}

// ParseAll parses contents in passed dir.
func ParseAll(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		return Parse(path)
	})
}

// Parse parses contents from file at passed path.
func Parse(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	frontmatter, body, err := m.Parse(file)
	if err != nil {
		return err
	}

	html := blackfriday.Run([]byte(body))

	content := map[string]interface{}{}
	content["path"] = path
	content["html"] = string(html)
	content["frontmatter"] = frontmatter
	contents[path] = content

	return nil
}

// Index returns contents.
func Index() []map[string]interface{} {
	list := []map[string]interface{}{}

	for _, content := range contents {
		list = append(list, content)
	}

	return list
}

// Get returns a content matched with passed path.
func Get(path string) (map[string]interface{}, bool) {
	content, ok := contents[path]
	return content, ok
}
