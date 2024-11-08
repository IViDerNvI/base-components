package v1

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ObjectMeta `json:"metadata,omitempty"`
	UserId     string    `json:"user_id,omitempty" gorm:"unique;column:userId;type:varchar(20);not null"`
	Account    string    `json:"account,omitempty" gorm:"column:account;type:varchar(20);not null"`
	Password   string    `json:"password,omitempty" gorm:"column:password;type:varchar(20);not null"`
	Phone      string    `json:"phone,omitempty" gorm:"column:phone;type:varchar(15)"`
	Email      string    `json:"email,omitempty" gorm:"column:email;type:varchar(30)"`
	LoginAt    time.Time `json:"loginat,omitempty" gorm:"column:loginAt"`
	IsAdmin    bool      `json:"is_admin,omitempty" gorm:"column:isAdmin"`
	IsDropped  bool      `json:"is_dropped,omitempty" gorm:"column:isDropped"`
	IsBanned   bool      `json:"is_banned,omitempty" gorm:"column:isBanned"`
}

type UserList struct {
	ListMeta `json:",inline"`
	Items    []*User `json:"items"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) IsPasswordCorrect(pwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd)); err != nil {
		return fmt.Errorf("failed to compile password: %w", pwd)
	}
	return nil
}
