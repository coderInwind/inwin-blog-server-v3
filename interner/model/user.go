package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BasicModel
	Username string
	Password string
	RoleId   int
	Role     Role `gorm:"ForeignKey:RoleId"`
}

//加密密码
func (u *User) EncryptPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) //加密的难度
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	fmt.Println(err)
	return err == nil
}
