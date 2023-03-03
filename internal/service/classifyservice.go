package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func QueryClassifyByUidAndTitle(tx model.IClassifyModel, uid, title string) (result model.ClassifyModel, err error) {
	result, err = tx.QueryByTitle(uid, title)
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.String("uid", uid), zap.String("title", title))
		err = i18n.NewCodeError(module.ClassifyQueryErr)
		return
	}
	return
}

func QueryAllClassify(tx model.IClassifyModel) (cms []model.ClassifyModel, err error) {
	cms, err = tx.QueryAll()
	return
}

func InsertClassify(tx model.IClassifyModel, cm *model.ClassifyModel) (id uint, err error) {
	_, err = QueryClassifyByUidAndTitle(tx, cm.Uid, cm.Title)
	if err != gorm.ErrRecordNotFound {
		err = i18n.NewCodeError(module.ClassifyExistedErr)
		return
	}
	id, err = tx.Insert(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(module.ClassifyInsertErr)
	}
	return
}

func UpdateClassify(tx model.IClassifyModel, cm model.ClassifyModel) (err error) {
	_, err = QueryClassifyByUidAndTitle(tx, cm.Uid, cm.Title)
	if err != nil {
		err = i18n.NewCodeError(module.ClassifyNotfoundErr)
		return
	}
	err = tx.Update(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(module.ClassifyUpdateErr)
	}
	return
}

func DeleteClassify(tx model.IClassifyModel, id uint, uid string) (err error) {
	err = tx.Delete(id, uid)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Uint("id", id), zap.String("uid", uid))
		err = i18n.NewCodeError(module.ClassifyUpdateErr)
	}
	return
}
