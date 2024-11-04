package v1

type Document struct {
	ObjectMeta `json:"metadata,omitempty"`
	DocumentId string  `json:"document_id,omitempty" gorm:"unique;column:documentId;type:varchar(20);not null"`
	Title      string  `json:"title,omitempty" gorm:"column:title;type:varchar(30);not null"`
	Authors    Authors `json:"author,omitempty" gorm:"column:author;type:varchar(30);not null"`
	Tags       Tags    `json:"tags,omitempty" gorm:"column:author;type:varchar(30);not null"`

	UpdatedUserId string `json:"user,omitempty" gorm:"colum:updatedUserId;type:varchar(20);not null"`
}

type DocumentList struct {
	ListMeta `json:",inline"`
	Items    []*Document `json:"items"`
}
