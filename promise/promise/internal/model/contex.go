package model

import (
	"promise/utility/token"

	"github.com/gogf/gf/v2/frame/g"
)

// Context 请求上下文结构
type Context struct {
	Token *token.MyCacheToken // token信息，包含上下文用户信息
	Data  g.Map               // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id       uint   // 用户ID
	Passport string // 用户账号
	Name     string // 用户名称
	Avatar   string // 用户头像
	IsAdmin  bool   // 是否是管理员
}
