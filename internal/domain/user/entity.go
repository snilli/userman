package user

import (
	"slices"
	"time"
	"userman/internal/domain/common"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	*common.BaseEntity
	Name     string
	Email    string
	Password string
	Role     string
}

type UpdateUser struct {
	Name  string
	Email string
}

func NewUser(name string, email string, password string) *User {
	now := time.Now()
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     "user",
		BaseEntity: &common.BaseEntity{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	user.setPassword(password)
	return user
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) SetRole(role string) {
	if contain := slices.Contains([]string{"admin", "user"}, role); contain {
		u.Role = role
		u.BaseEntity.SetChange("role", role)

		u.BaseEntity.UpdatedAt = time.Now()
		u.BaseEntity.SetChange("updated_at", u.BaseEntity.UpdatedAt)
	}
}

func (u *User) Update(input *UpdateUser) {
	updated := false
	if input.Name != "" && input.Name != u.Name {
		u.Name = input.Name
		u.BaseEntity.SetChange("name", input.Name)
		updated = true
	}

	if input.Email != "" && input.Email != u.Email {
		u.Email = input.Email
		u.BaseEntity.SetChange("email", input.Email)
		updated = true
	}

	if updated {
		u.BaseEntity.UpdatedAt = time.Now()
		u.BaseEntity.SetChange("updated_at", u.BaseEntity.UpdatedAt)
	}
}

func (u *User) setPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.BaseEntity.SetChange("password", u.Password)

	return nil
}

func (u *User) CheckPassword(plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err
}

var _ common.Entity = User{}
