// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	Id       interface{} // 用户唯一标识
	Username interface{} // 用户姓名
	Mobile   interface{} // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	Picture  interface{} // 用户头像 url
	Email    interface{} // 用户邮箱
}
