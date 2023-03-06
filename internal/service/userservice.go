package service

import (
	"errors"
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

func ParseToken(token string) (jtoken *jwt.Token, err error) {
	return utils.ParseToken(token, constx.TokenSecret)
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
		Email:        rmr.Email,
		NickName:     rmr.NickName,
		Pwd:          pwd,
		CreatedBy:    uid,
		RegisterType: constx.EmailRT,
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
		Phone:        rmr.Phone,
		NickName:     rmr.NickName,
		Pwd:          pwd,
		CreatedBy:    uid,
		RegisterType: constx.PhoneRT,
	}
	return register(tx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-04
 * @Description    : 通过uid注册
 * @param           {model.IUserModel} tx
 * @return          {*}
 */
func RegisterByUid(tx model.IUserModel) (uid string, err error) {
	uid, err = utils.NewUlid()
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.UserRegisterErr)
		return
	}
	um := &model.UserModel{
		CreatedBy: uid,
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
	uid = um.CreatedBy
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
	if um.CreatedBy == "" {
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um.Pwd != pwd {
		err = i18n.NewCodeError(module.UserPwdErr)
		return
	}
	uid = um.CreatedBy
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-06
 * @Description    : 通过手机号登录
 * @param           {model.IUserModel} tx
 * @param           {entity.LoginByPhoneReq} lpq
 * @return          {*}
 */
func LoginByPhone(tx model.IUserModel, lpq entity.LoginByPhoneReq) (uid string, err error) {
	pwd, err := utils.EncryptPwd(lpq.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
		return
	}
	um, err := tx.QueryUser(model.UserModel{Phone: lpq.Phone})
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", lpq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Info(err.Error(), zap.Any("model", lpq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um.Pwd != pwd {
		loggerx.ZapLog.Info("pwd validate failure", zap.Any("model", lpq))
		err = i18n.NewCodeError(module.UserPwdErr)
		return
	}
	uid = um.CreatedBy
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-04
 * @Description    : 通过uid登录
 * @param           {model.IUserModel} tx
 * @param           {entity.LoginByUidReq} luq
 * @return          {*}
 */
func LoginByUid(tx model.IUserModel, luq entity.LoginByUidReq) (uid string, err error) {
	um, err := tx.QueryUser(model.UserModel{CreatedBy: luq.Uid})
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", luq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", luq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um.Email != "" ||
		um.Phone != "" {
		err = i18n.NewCodeError(module.UserBindNoUidLoginErr)
		return
	}
	uid = luq.Uid
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-06
 * @Description    : 给用户绑定邮箱
 * @param           {model.IUserModel} tx
 * @param           {string} uid
 * @param           {entity.BindEmailReq} req
 * @return          {*}
 */
func BindEmailByUid(tx model.IUserModel, uid string, req entity.BindEmailReq) (err error) {
	um1, err := tx.QueryUser(model.UserModel{CreatedBy: uid})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", uid), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um1.Email != "" {
		err = i18n.NewCodeError(module.UserBindedEmailErr)
		return
	}
	_, err = tx.QueryUser(model.UserModel{Email: req.Email})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.UserEmailBindByOtherErr)
		return
	}
	err = tx.Update(model.UserModel{CreatedBy: uid, Email: req.Email})
	if err != nil {
		err = i18n.NewCodeError(module.UserUpdateErr)
		return
	}
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-06
 * @Description    : 给用户绑定电话
 * @param           {model.IUserModel} tx
 * @param           {string} uid
 * @param           {entity.BindPhoneReq} req
 * @return          {*}
 */
func BindPhoneByUid(tx model.IUserModel, uid string, req entity.BindPhoneReq) (err error) {
	um1, err := tx.QueryUser(model.UserModel{CreatedBy: uid})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", uid), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um1.Email != "" {
		err = i18n.NewCodeError(module.UserBindedPhoneErr)
		return
	}
	_, err = tx.QueryUser(model.UserModel{Phone: req.Phone})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.UserEmailBindByOtherErr)
		return
	}
	err = tx.Update(model.UserModel{CreatedBy: uid, Phone: req.Phone})
	if err != nil {
		err = i18n.NewCodeError(module.UserUpdateErr)
		return
	}
	return
}
