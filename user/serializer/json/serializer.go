package json

import (
	"encoding/json"

	"github.com/danielhoward314/tenantwin/user/svc"
)

type User struct{}

func (u *User) Decode(input []byte) (*svc.User, error) {
	user := &svc.User{}
	if err := json.Unmarshal(input, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Encode(user *svc.User) ([]byte, error) {
	rawMsg, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return rawMsg, nil
}
