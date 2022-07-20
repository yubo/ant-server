package api

import (
	"fmt"
	"time"

	"github.com/yubo/apiserver/pkg/rest"
	"github.com/yubo/golib/util"
)

type NameParam struct {
	Name string `param:"path" name:"name"`
}

type ListParam struct {
	rest.Pagination
	Query *string `param:"query" name:"query" description:"query user"`
}

type LoginInput struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	AutoLogin bool   `json:"autoLogin"`
	Type      string `json:"type"`
}

type LoginOutput struct {
	Type             string `json:"type"`
	CurrentAuthority string `json:"currentAuthority"`
}

type CaptchaInput struct {
	Phone string `json:"phone"`
}

type User struct {
	Id          int64             `json:"id" sql:",primary_key,auto_increment=1000"`
	Name        string            `json:"name" sql:",where,index,size=32"`
	Avatar      string            `json:"avatar"`
	Email       string            `json:"email"`
	Title       string            `json:"title"`
	Group       string            `json:"group"`
	Tags        map[string]string `json:"tags"`
	NotifyCount int               `json:"notifyCount"`
	UnreadCount int               `json:"unreadCount"`
	Address     string            `json:"address"`
	Phone       string            `json:"phone"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

type CreateUserInput struct {
	Name      *string           `json:"name"`
	Avatar    *string           `json:"avatar"`
	Email     *string           `json:"email"`
	Title     *string           `json:"title"`
	Group     *string           `json:"group"`
	Tags      map[string]string `json:"tags"`
	Address   *string           `json:"address"`
	Phone     *string           `json:"phone"`
	CreatedAt time.Time         `json:"-"`
	UpdatedAt time.Time         `json:"-"`
}

func (p *CreateUserInput) Validate() error {
	if util.StringValue(p.Name) == "" {
		return fmt.Errorf("name not set")
	}

	return nil
}

type ListUserOutput struct {
	Total int64   `json:"total"`
	List  []*User `json:"list"`
}

type UpdateUserInput struct {
	Name      *string           `json:"name" sql:",where"`
	Avatar    *string           `json:"avatar"`
	Email     *string           `json:"email"`
	Title     *string           `json:"title"`
	Group     *string           `json:"group"`
	Tags      map[string]string `json:"tags"`
	Address   *string           `json:"address"`
	Phone     *string           `json:"phone"`
	UpdatedAt time.Time         `json:"-"`
}

func (p *UpdateUserInput) Validate() error {
	if util.StringValue(p.Name) == "" {
		return fmt.Errorf("id & name not set")
	}

	return nil
}
