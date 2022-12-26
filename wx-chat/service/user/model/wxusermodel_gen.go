// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	wxUserFieldNames          = builder.RawFieldNames(&WxUser{})
	wxUserRows                = strings.Join(wxUserFieldNames, ",")
	wxUserRowsExpectAutoSet   = strings.Join(stringx.Remove(wxUserFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")
	wxUserRowsWithPlaceHolder = strings.Join(stringx.Remove(wxUserFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"
)

type (
	wxUserModel interface {
		Insert(ctx context.Context, data *WxUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*WxUser, error)
		FindOneByUsername(ctx context.Context, username string) (*WxUser, error)
		Update(ctx context.Context, data *WxUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWxUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	WxUser struct {
		Id         int64     `db:"id"`
		Username   string    `db:"username"` // username
		Password   string    `db:"password"` // password
		Gender     int64     `db:"gender"`   // 0-un status｜1-male｜2-female
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newWxUserModel(conn sqlx.SqlConn) *defaultWxUserModel {
	return &defaultWxUserModel{
		conn:  conn,
		table: "`wx_user`",
	}
}

func (m *defaultWxUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultWxUserModel) FindOne(ctx context.Context, id int64) (*WxUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wxUserRows, m.table)
	var resp WxUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWxUserModel) FindOneByUsername(ctx context.Context, username string) (*WxUser, error) {
	var resp WxUser
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", wxUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWxUserModel) Insert(ctx context.Context, data *WxUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, wxUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Gender)
	return ret, err
}

func (m *defaultWxUserModel) Update(ctx context.Context, newData *WxUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wxUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Username, newData.Password, newData.Gender, newData.Id)
	return err
}

func (m *defaultWxUserModel) tableName() string {
	return m.table
}
