package service

import (
	"time"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 创建token
 * @param           {string} uid
 * @return          {*}
 */
func GenerateToken(uid string) (token string, err error) {
	id, err := utils.NewUlid()
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}
	claim := utils.GenerateClaims(jwt.RegisteredClaims{
		Issuer:    constx.TokenIssuer,
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(constx.TokenExpiresAt * time.Second)},
		Subject:   constx.TokenSub,
		ID:        id,
	}, map[string]interface{}{
		constx.TokenUid: uid,
	})
	token, err = utils.GenerateToken(constx.TokenSecret, claim)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("claim", claim))
		return
	}
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 邮箱注册
 * @param           {model.IUserModel} tx
 * @param           {entity.RegisterEmailReq} rmr
 * @return          {*}
 */
func RegisterByEmail(tx model.IUserModel, rmr entity.RegisterByEmailReq) (uid string, err error) {
	pwd, err := utils.EncryptPwd(rmr.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
		return
	}
	uid, err = utils.NewUlid()
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.UserRegisterErr)
		return
	}
	um := &model.UserModel{
		Email:    rmr.Email,
		NickName: rmr.NickName,
		Pwd:      pwd,
		Uid:      uid,
	}
	return register(tx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-01
 * @Description    : 电话注册
 * @param           {model.IUserModel} tx
 * @param           {entity.RegisterByPhoneReq} rmr
 * @return          {*}
 */
func RegisterByPhone(tx model.IUserModel, rmr entity.RegisterByPhoneReq) (uid string, err error) {
	isPhone := utils.CheckMobile(rmr.Phone)
	if !isPhone {
		err = i18n.NewCodeError(module.UserPhoneErr)
		return
	}
	pwd, err := utils.EncryptPwd(rmr.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
		return
	}
	uid, err = utils.NewUlid()
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.UserRegisterErr)
		return
	}
	um := &model.UserModel{
		Email:    rmr.Phone,
		NickName: rmr.NickName,
		Pwd:      pwd,
		Uid:      uid,
	}
	return register(tx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-01
 * @Description    : 注册通用校验逻辑
 * @param           {model.IUserModel} tx
 * @param           {*model.UserModel} um
 * @return          {*}
 */
func register(tx model.IUserModel, um *model.UserModel) (uid string, err error) {
	_, err = tx.CheckRegisterRepeat(*um)
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", um))
		err = i18n.NewCodeError(module.UserRegisterErr)
		return
	}
	if err != gorm.ErrRecordNotFound {
		err = i18n.NewCodeError(module.UserRegisterRepeatErr)
		return
	}
	_, err = tx.Insert(um)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", um))
		err = i18n.NewCodeError(module.UserRegisterErr)
		return
	}
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 通过邮箱登录
 * @param           {model.IUserModel} tx
 * @param           {entity.LoginByEmailReq} leq
 * @return          {*}
 */
func LoginByEmail(tx model.IUserModel, leq entity.LoginByEmailReq) (uid string, err error) {
	pwd, err := utils.EncryptPwd(leq.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
		return
	}
	um, err := tx.QueryUser(model.UserModel{Email: leq.Email})
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", leq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um.Uid == "" {
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um.Pwd != pwd {
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	uid = um.Uid
	return
}
