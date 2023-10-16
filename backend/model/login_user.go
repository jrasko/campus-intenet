package model

import (
	"errors"
	"fmt"
	"slices"

	"github.com/alexedwards/argon2id"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleFinance Role = "financer"
	RoleViewer  Role = "viewer"
)

const (
	FieldUsername = "username"
	FieldRole     = "role"
)

type LoginUser struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Role         Role   `json:"role"`
}

func (u LoginUser) Validate() error {
	allowedRoles := []Role{RoleAdmin, RoleFinance, RoleViewer}
	if !slices.Contains(allowedRoles, u.Role) {
		return fmt.Errorf("invalid role %s", u.Role)
	}
	if u.Username == "" {
		return errors.New("username must not be empty")
	}
	if _, _, _, err := argon2id.DecodeHash(u.PasswordHash); err != nil {
		return err
	}
	return nil
}
