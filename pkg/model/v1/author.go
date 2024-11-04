package v1

import (
	stringutil "github.com/ividernvi/base-components/pkg/util/stringutil"
)

type Authors stringutil.StringSet

func CreateAuthors(authors ...string) Authors {
	return Authors(stringutil.CreateStringSet(authors...))
}

func (authors1 Authors) Equals(authors2 Authors) bool {
	return stringutil.StringSet(authors1).Equals(stringutil.StringSet(authors2))
}

func (authors Authors) AddAuthor(author ...string) {
	for _, au := range author {
		stringutil.StringSet(authors).Add(au)
	}
}
