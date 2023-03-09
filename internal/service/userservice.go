package service

import (
	"errors"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 邮箱注册
 * @param           {model.IUserModel} tx
 * @param           {entity.RegisterEmailReq} rmr
 * @return          {*}
 */
func RegisterByEmail(ctx *contextx.Contextx, tx model.IUserModel, ttx model.ITeamModel,
	rmr entity.RegisterByEmailReq) (resp model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(rmr.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(ctx.Language, module.EncryptPwdErr)
		return
	}
	um := &model.UserModel{
		Email:        rmr.Email,
		NickName:     rmr.NickName,
		Pwd:          pwd,
		CreatedBy:    utils.NewUlid(),
		TeamIdPort:   utils.NewUlid(),
		RegisterType: constx.EmailRT,
	}
	return register(ctx, tx, ttx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-01
 * @Description    : 电话注册
 * @param           {model.IUserModel} tx
 * @param           {entity.RegisterByPhoneReq} rmr
 * @return          {*}
 */
func RegisterByPhone(ctx *contextx.Contextx, tx model.IUserModel, ttx model.ITeamModel,
	rmr entity.RegisterByPhoneReq) (resp model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(rmr.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(ctx.Language, module.EncryptPwdErr)
		return
	}
	um := &model.UserModel{
		Phone:        rmr.Phone,
		NickName:     rmr.NickName,
		Pwd:          pwd,
		CreatedBy:    utils.NewUlid(),
		TeamIdPort:   utils.NewUlid(),
		RegisterType: constx.PhoneRT,
	}
	return register(ctx, tx, ttx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-04
 * @Description    : 通过uid注册
 * @param           {model.IUserModel} tx
 * @return          {*}
 */
func RegisterByUid(ctx *contextx.Contextx, utx model.IUserModel, ttx model.ITeamModel) (umresp model.UserModel, err error) {
	um := &model.UserModel{
		CreatedBy:  utils.NewUlid(),
		TeamIdPort: utils.NewUlid(),
	}
	return register(ctx, utx, ttx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-01
 * @Description    : 注册通用校验逻辑
 * @param           {model.IUserModel} tx
 * @param           {*model.UserModel} um
 * @return          {*}
 */
func register(ctx *contextx.Contextx, utx model.IUserModel, ttx model.ITeamModel, um *model.UserModel) (umresp model.UserModel, err error) {
	_, err = utx.CheckRegisterRepeat(*um)
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", um))
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterRepeatErr)
		return
	}

	_, err = ttx.Query(model.TeamModel{TeamId: um.TeamIdPort})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(ctx.Language, module.TeamRepeatErr)
		return
	}

	_, err = utx.Insert(um)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", um))
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		return
	}

	tx2 := utx.GetTx()
	err = container.InitData(tx2, ctx.Language, um.CreatedBy, um.TeamIdPort)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(ctx.Language, module.UserDataInit)
		return
	}
	umresp = *um
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
func LoginByEmail(ctx *contextx.Contextx, tx model.IUserModel, leq entity.LoginByEmailReq) (um model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(leq.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(ctx.Language, module.EncryptPwdErr)
		return
	}
	um, err = tx.QueryUser(model.UserModel{Email: leq.Email})
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", leq))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if um.Pwd != pwd {
		err = i18n.NewCodeError(ctx.Language, module.UserPwdErr)
		return
	}
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
func LoginByPhone(ctx *contextx.Contextx, tx model.IUserModel, lpq entity.LoginByPhoneReq) (um model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(lpq.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(ctx.Language, module.EncryptPwdErr)
		return
	}
	um, err = tx.QueryUser(model.UserModel{Phone: lpq.Phone})
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", lpq))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Info(err.Error(), zap.Any("model", lpq))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if um.Pwd != pwd {
		loggerx.ZapLog.Info("pwd validate failure", zap.Any("model", lpq))
		err = i18n.NewCodeError(ctx.Language, module.UserPwdErr)
		return
	}
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
func LoginByUid(ctx *contextx.Contextx, tx model.IUserModel, luq entity.LoginByUidReq) (um model.UserModel, err error) {
	um, err = tx.QueryUser(model.UserModel{CreatedBy: luq.Uid})
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", luq))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", luq))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if um.Email != "" ||
		um.Phone != "" {
		err = i18n.NewCodeError(ctx.Language, module.UserBindNoUidLoginErr)
		return
	}
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
func BindEmailByUid(ctx *contextx.Contextx, tx model.IUserModel, uid string, req entity.BindEmailReq) (err error) {
	um1, err := tx.QueryUser(model.UserModel{CreatedBy: uid})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", uid), zap.Any("model", req))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if um1.Email != "" {
		err = i18n.NewCodeError(ctx.Language, module.UserBindedEmailErr)
		return
	}
	_, err = tx.QueryUser(model.UserModel{Email: req.Email})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", req))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(ctx.Language, module.UserEmailBindByOtherErr)
		return
	}
	err = tx.Update(model.UserModel{CreatedBy: uid, Email: req.Email})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserUpdateErr)
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
func BindPhoneByUid(ctx *contextx.Contextx, tx model.IUserModel, uid string, req entity.BindPhoneReq) (err error) {
	um1, err := tx.QueryUser(model.UserModel{CreatedBy: uid})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", uid), zap.Any("model", req))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if um1.Email != "" {
		err = i18n.NewCodeError(ctx.Language, module.UserBindedPhoneErr)
		return
	}
	_, err = tx.QueryUser(model.UserModel{Phone: req.Phone})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", req))
		err = i18n.NewCodeError(ctx.Language, module.UserNotFoundErr)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(ctx.Language, module.UserEmailBindByOtherErr)
		return
	}
	err = tx.Update(model.UserModel{CreatedBy: uid, Phone: req.Phone})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserUpdateErr)
		return
	}
	return
}
