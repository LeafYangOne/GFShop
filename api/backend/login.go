package backend

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"GFShop/internal/model/entity"
)

type LoginDoReq struct {
	g.Meta   `path:"/backend/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	// Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}

type LoginDoRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// for jwt
// type RefreshTokenReq struct {
// 	g.Meta `path:"/backend/refresh_token" method:"post"`
// }

// LoginRes for token
type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`    //是否超管
	RoleIds     string                  `json:"role_ids"`    //角色
	Permissions []entity.PermissionInfo `json:"permissions"` //权限列表
}
type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
type LogoutReq struct {
	g.Meta `path:"/backend/logout" method:"post"`
}

type LogoutRes struct {
}
