// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package v1

import (
	"time"

	"gorm.io/gorm"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type ObjectMeta struct {
	ID         uint64         `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	InstanceID string         `json:"instanceID,omitempty" gorm:"unique;column:instanceID;type:varchar(32);not null"`
	Name       string         `json:"name,omitempty" gorm:"column:name;type:varchar(64);not null" validate:"name"`
	CreatedAt  time.Time      `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	Deleted    gorm.DeletedAt `json:"deltedAt,omitempty" `
}

type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

type ListOptions struct {
	TypeMeta       `json:",inline"`
	LabelSelector  string `json:"labelSelector,omitempty" form:"labelSelector"`
	FieldSelector  string `json:"fieldSelector,omitempty" form:"fieldSelector"`
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
	Offset         *int64 `json:"offset,omitempty" form:"offset"`
	Limit          *int64 `json:"limit,omitempty" form:"limit"`
}

type GetOptions struct {
	TypeMeta `json:",inline"`
}

type DeleteOptions struct {
	TypeMeta `json:",inline"`
	Unscoped bool `json:"unscoped"`
}

type CreateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}

type UpdateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}
