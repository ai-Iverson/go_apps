// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cert is the golang structure of table cert for DAO operations like Where/Data.
type Cert struct {
	g.Meta    `orm:"table:cert, do:true"`
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	IsDeleted interface{} //
	Number    interface{} // 证书 ID
	DevId     interface{} // 开发者 ID
	DevOrg    interface{} // 开发者组织
	CertPass  interface{} // 证书密码
	BeginTime interface{} // 颁发时间
	EndTime   interface{} // 过期时间
	Content   interface{} // 证书内容
	Id        interface{} //
}
