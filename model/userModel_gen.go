// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheXinclockinUserIdPrefix             = "cache:xinclockin:user:id:"
	cacheXinclockinUserIdentityNumberPrefix = "cache:xinclockin:user:identityNumber:"
	cacheXinclockinUserUidPrefix            = "cache:xinclockin:user:uid:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByIdentityNumber(ctx context.Context, identityNumber string) (*User, error)
		FindOneByUid(ctx context.Context, uid string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id             int64          `db:"id"`
		IdentityNumber string         `db:"identity_number"`
		Uid            string         `db:"uid"`
		FullName       sql.NullString `db:"full_name"`
		AvatarUrl      sql.NullString `db:"avatar_url"`
		SessionId      sql.NullString `db:"session_id"`
		Biography      sql.NullString `db:"biography"`
		PrivateKey     sql.NullString `db:"private_key"`
		ClientId       sql.NullString `db:"client_id"`
		Contract       sql.NullString `db:"contract"`
		IsMvmUser      int64          `db:"is_mvm_user"`
		CreatedAt      time.Time      `db:"created_at"`
		UpdatedAt      time.Time      `db:"updated_at"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	xinclockinUserIdKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdPrefix, id)
	xinclockinUserIdentityNumberKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdentityNumberPrefix, data.IdentityNumber)
	xinclockinUserUidKey := fmt.Sprintf("%s%v", cacheXinclockinUserUidPrefix, data.Uid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, xinclockinUserIdKey, xinclockinUserIdentityNumberKey, xinclockinUserUidKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	xinclockinUserIdKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, xinclockinUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByIdentityNumber(ctx context.Context, identityNumber string) (*User, error) {
	xinclockinUserIdentityNumberKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdentityNumberPrefix, identityNumber)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, xinclockinUserIdentityNumberKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `identity_number` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, identityNumber); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUid(ctx context.Context, uid string) (*User, error) {
	xinclockinUserUidKey := fmt.Sprintf("%s%v", cacheXinclockinUserUidPrefix, uid)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, xinclockinUserUidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	xinclockinUserIdKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdPrefix, data.Id)
	xinclockinUserIdentityNumberKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdentityNumberPrefix, data.IdentityNumber)
	xinclockinUserUidKey := fmt.Sprintf("%s%v", cacheXinclockinUserUidPrefix, data.Uid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.IdentityNumber, data.Uid, data.FullName, data.AvatarUrl, data.SessionId, data.Biography, data.PrivateKey, data.ClientId, data.Contract, data.IsMvmUser)
	}, xinclockinUserIdKey, xinclockinUserIdentityNumberKey, xinclockinUserUidKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	xinclockinUserIdKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdPrefix, data.Id)
	xinclockinUserIdentityNumberKey := fmt.Sprintf("%s%v", cacheXinclockinUserIdentityNumberPrefix, data.IdentityNumber)
	xinclockinUserUidKey := fmt.Sprintf("%s%v", cacheXinclockinUserUidPrefix, data.Uid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.IdentityNumber, newData.Uid, newData.FullName, newData.AvatarUrl, newData.SessionId, newData.Biography, newData.PrivateKey, newData.ClientId, newData.Contract, newData.IsMvmUser, newData.Id)
	}, xinclockinUserIdKey, xinclockinUserIdentityNumberKey, xinclockinUserUidKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheXinclockinUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
