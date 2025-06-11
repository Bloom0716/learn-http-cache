package database

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/Bloom0716/learn-http-cache/internal/model"
)

func (r *Repository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *model.User) (*model.User, error) {
	etagSource := fmt.Sprintf("%d:%s", user.ID, user.Name)
	hasher := sha1.New()
	hasher.Write([]byte(etagSource))
	user.Etag = hex.EncodeToString(hasher.Sum(nil))

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
