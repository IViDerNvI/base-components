package v1

import "testing"

func TestAddAuthor(t *testing.T) {
	cases := []struct {
		Name     string
		Id       int
		Authors  Authors
		AuthorA  []string
		Expected Authors
	}{
		{"BasicTest", 1, CreateAuthors("Adam"), []string{"Eve"}, CreateAuthors("Adam", "Eve")},
		{"BasicTest", 2, CreateAuthors("Rose", "Jim"), []string{"John", "Joney"}, CreateAuthors("Rose", "Jim", "John", "Joney")},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			c.Authors.AddAuthor(c.AuthorA...)
			if !c.Authors.Equals(c.Expected) {
				t.Fatalf("%s#%2d: expected %v, but %v got",
					c.Name, c.Id, c.Authors, c.Expected)
			}
		})
	}
}
