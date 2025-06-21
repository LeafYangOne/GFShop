package login

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"GFShop/internal/dao"
	"GFShop/internal/model"
	"GFShop/internal/model/entity"
	"GFShop/internal/service"
	"GFShop/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}
func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New(`账号或密码错误`)
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	return nil
}
