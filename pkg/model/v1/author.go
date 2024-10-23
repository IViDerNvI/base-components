package v1

import (
	"database/sql/driver"

	"github.com/lib/pq"
)

type AuthorsUnion pq.StringArray

func (au *AuthorsUnion) Scan(src interface{}) error {
	return (*pq.StringArray)(au).Scan(src)
}

func (au *AuthorsUnion) Value() (driver.Value, error) {
	return (*pq.StringArray)(au).Value()
}

func (au *AuthorsUnion) AddAuthor(author string) *AuthorsUnion {
	*au = append((*au), author)
	return au
}

func (au *AuthorsUnion) GetCount() int {
	return len((*au))
}
