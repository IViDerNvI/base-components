package v1

import (
	stringutil "github.com/ividernvi/base-components/pkg/util/stringutil"
)

type Tags stringutil.StringSet

func CreateTags(tags ...string) Tags {
	return Tags(stringutil.CreateStringSet(tags...))
}

func (tags1 Tags) Equals(tags2 Tags) bool {
	return stringutil.StringSet(tags1).Equals(stringutil.StringSet(tags2))
}

func (tags Tags) AddTag(tag ...string) {
	for _, tg := range tag {
		stringutil.StringSet(tags).Add(tg)
	}
}
