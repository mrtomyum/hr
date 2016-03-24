package models

import (
	"fmt"
	"github.com/elithrar/simple-scrypt"
	"github.com/grafana/grafana/pkg/cmd/grafana-cli/log"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password []byte
	PersonID uint
}

func (u *User) ChangePass(p string) error {
	hash, err := scrypt.GenerateFromPassword([]byte(p), scrypt.DefaultParams)
	if err != nil {
		log.Error(err)
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) VerifyPass(p string) error {

	err := scrypt.CompareHashAndPassword(u.Password, []byte(p))
	if err != nil {
		return err
	}
	fmt.Println("User []byte Password Hash:", u.Password)
	fmt.Println("User []byte Password:", []byte(p))
	return nil
}
