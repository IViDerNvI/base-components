package v1

import (
	"database/sql/driver"

	"github.com/lib/pq"
)

type TagsUnion pq.StringArray

func (tu *TagsUnion) Scan(src interface{}) error {
	return (*pq.StringArray)(tu).Scan(src)
}

func (tu *TagsUnion) Value() (driver.Value, error) {
	return (*pq.StringArray)(tu).Value()
}

func (tu *TagsUnion) AddTag(tag string) *TagsUnion {
	*tu = append((*tu), tag)
	return tu
}

func (tu *TagsUnion) GetCount() int {
	return len((*tu))
}
