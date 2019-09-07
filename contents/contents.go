package contents

import (
	"io"

	"github.com/gernest/front"
)

var contents = []map[string]interface{}{}

// Parse parses contents from r.
func Parse(r io.Reader) error {
	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	content, body, err := m.Parse(r)
	if err != nil {
		return err
	}

	content["body"] = body
	contents = append(contents, content)

	return nil
}

// Index returns contents.
func Index() []map[string]interface{} {
	return contents
}
