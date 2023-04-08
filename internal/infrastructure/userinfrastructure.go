package infrastructure

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
func RegisterByEmail(ctx *contextx.Contextx,
	rmr entity.RegisterByEmailReq) (resp model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(rmr.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
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
	return register(ctx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-01
 * @Description    : 电话注册
 * @param           {model.IUserModel} tx
 * @param           {entity.RegisterByPhoneReq} rmr
 * @return          {*}
 */
func RegisterByPhone(ctx *contextx.Contextx,
	rmr entity.RegisterByPhoneReq) (resp model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(rmr.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
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
	return register(ctx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-04
 * @Description    : 通过uid注册
 * @param           {model.IUserModel} tx
 * @return          {*}
 */
func RegisterByUid(ctx *contextx.Contextx) (umresp model.UserModel, err error) {
	um := &model.UserModel{
		CreatedBy:  utils.NewUlid(),
		TeamIdPort: utils.NewUlid(),
	}
	return register(ctx, um)
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-01
 * @Description    : 注册通用校验逻辑
 * @param           {model.IUserModel} tx
 * @param           {*model.UserModel} um
 * @return          {*}
 */
func register(ctx *contextx.Contextx, um *model.UserModel) (umresp model.UserModel, err error) {
	_, err = ctx.BaseTx.UserModel.CheckRegisterRepeat(*um)
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", um))
		err = i18n.NewCodeError(module.UserRegisterErr)
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.UserRegisterRepeatErr)
		return
	}

	_, err = ctx.BaseTx.TeamModel.First(model.TeamModel{TeamId: um.TeamIdPort})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.TeamRepeatErr)
		return
	}

	// _, err = ctx.BaseTx.UserModel.Insert(um)
	// if err != nil {
	// 	loggerx.ZapLog.Error(err.Error(), zap.Any("model", um))
	// 	err = i18n.NewCodeError(module.UserRegisterErr)
	// 	return
	// }

	tx2 := ctx.BaseTx.Db
	err = container.InitData(tx2, ctx.Language, um.CreatedBy, um.TeamIdPort)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.UserDataInit)
		return
	}
	umresp = *um
	ctx.UID = um.CreatedBy
	ctx.TID = um.TeamIdPort
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
func LoginByEmail(ctx *contextx.Contextx, leq entity.LoginByEmailReq) (um model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(leq.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
		return
	}
	um, err = ctx.BaseTx.UserModel.QueryUser(model.UserModel{Email: leq.Email})
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", leq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um.Pwd != pwd {
		err = i18n.NewCodeError(module.UserPwdErr)
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
func LoginByPhone(ctx *contextx.Contextx, lpq entity.LoginByPhoneReq) (um model.UserModel, err error) {
	pwd, err := utils.EncryptPwd(lpq.Pwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("pwd", pwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
		return
	}
	um, err = ctx.BaseTx.UserModel.QueryUser(model.UserModel{Phone: lpq.Phone})
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
func LoginByUid(ctx *contextx.Contextx, luq entity.LoginByUidReq) (um model.UserModel, err error) {
	um, err = ctx.BaseTx.UserModel.QueryUser(model.UserModel{CreatedBy: luq.Uid})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", luq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", luq))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	return
}

func GetUserInfo(ctx *contextx.Contextx, uid string) (u model.UserModel, err error) {
	u, err = ctx.BaseTx.UserModel.QueryUser(model.UserModel{CreatedBy: uid})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", uid))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", uid))
		err = i18n.NewCodeError(module.UserNotFoundErr)
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
func BindEmailByUid(ctx *contextx.Contextx, req entity.BindEmailReq) (err error) {
	_, err = ctx.BaseTx.UserModel.QueryUser(model.UserModel{CreatedBy: ctx.UID})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", ctx.UID), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	u2, err := ctx.BaseTx.UserModel.QueryUser(model.UserModel{Email: req.Email})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) && u2.CreatedBy != ctx.UID {
		err = i18n.NewCodeError(module.UserEmailBindByOtherErr)
		return
	}
	err = ctx.BaseTx.UserModel.Update(model.UserModel{CreatedBy: ctx.UID, Email: req.Email})
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
func BindPhoneByUid(ctx *contextx.Contextx, req entity.BindPhoneReq) (err error) {
	_, err = ctx.BaseTx.UserModel.QueryUser(model.UserModel{CreatedBy: ctx.UID})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", ctx.UID), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	um2, err := ctx.BaseTx.UserModel.QueryUser(model.UserModel{Phone: req.Phone})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", req))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) && um2.CreatedBy != ctx.UID {
		err = i18n.NewCodeError(module.UserPhoneBindByOtherErr)
		return
	}
	err = ctx.BaseTx.UserModel.Update(model.UserModel{CreatedBy: ctx.UID, Phone: req.Phone})
	if err != nil {
		err = i18n.NewCodeError(module.UserUpdateErr)
		return
	}
	return
}

func UserResetPwd(ctx *contextx.Contextx, oldPwd, newPwd string) (err error) {
	um1, err := ctx.BaseTx.UserModel.QueryUser(model.UserModel{CreatedBy: ctx.UID})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", ctx.UID))
		err = i18n.NewCodeError(module.UserNotFoundErr)
		return
	}
	if um1.Pwd != oldPwd {
		err = i18n.NewCodeError(module.UserPwdErr)
		return
	}
	pwd, err := utils.EncryptPwd(newPwd, constx.PwdSalt)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", ctx.UID), zap.Any("pwd", newPwd))
		err = i18n.NewCodeError(module.EncryptPwdErr)
		return
	}
	err = ctx.BaseTx.UserModel.Update(model.UserModel{CreatedBy: ctx.UID, Pwd: pwd})
	if err != nil {
		err = i18n.NewCodeError(module.UserUpdateErr)
		return
	}
	return
}

func UserUpdate(ctx *contextx.Contextx, req model.UserModel) (err error) {
	_, err = GetUserInfo(ctx, req.CreatedBy)
	if err != nil {
		return
	}
	err = ctx.BaseTx.UserModel.Update(req)
	if err != nil {
		err = i18n.NewCodeError(module.UserUpdateErr)
		return
	}
	return
}
