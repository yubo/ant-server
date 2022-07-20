package models

import (
	"context"

	"github.com/yubo/ant-server/pkg/api"
	"github.com/yubo/apiserver/pkg/models"
	"github.com/yubo/apiserver/pkg/storage"
	"github.com/yubo/golib/orm"
)

func NewUser() *User {
	o := &User{DB: models.DB()}
	return o
}

// user implements the user interface.
type User struct {
	orm.DB
}

func (p *User) Name() string {
	return "user"
}

func (p *User) NewObj() interface{} {
	return &api.User{}
}

func (p *User) Create(ctx context.Context, obj *api.CreateUserInput) error {
	// use snamekcassedName(objs.Name()) as table name
	return p.Insert(obj, orm.WithTable(p.Name()))
}

// Get retrieves the User from the db for a given name.
func (p *User) Get(ctx context.Context, name string) (ret *api.User, err error) {
	err = p.Query("select * from user where name=?", name).Row(&ret)
	return
}

// List lists all Secrets in the indexer.
func (p *User) List(ctx context.Context, opts storage.ListOptions) (list []*api.User, err error) {
	err = p.DB.List(&list,
		orm.WithTable(p.Name()),
		orm.WithTotal(opts.Total),
		orm.WithSelector(opts.Query),
		orm.WithOrderby(opts.Orderby...),
		orm.WithLimit(opts.Offset, opts.Limit),
	)
	return
}

func (p *User) Update(ctx context.Context, obj *api.UpdateUserInput) error {
	// use snamekcassedName(objs.Name()) as table name
	return p.DB.Update(obj, orm.WithTable(p.Name()))
}

func (p *User) Delete(ctx context.Context, name string) error {
	_, err := p.DB.Exec("delete demo where name=?", name)
	return err
}

func init() {
	models.Register(&User{})
}
