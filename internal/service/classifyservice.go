package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
)

func CreateClassify(tx model.IClassifyModel, cm *model.ClassifyModel) (id uint, err error) {
	id, err = tx.Insert(cm)
	return
}

func QueryAllClassify(tx model.IClassifyModel) (cms []model.ClassifyModel, err error) {
	cms, err = tx.QueryAll()
	return
}

func InsertClassify(tx model.IClassifyModel, cm *model.ClassifyModel) (id uint, err error) {
	id, err = tx.Insert(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(module.ClassifyInsertErr)
	}
	return
}

func UpdateClassify(tx model.IClassifyModel, cm model.ClassifyModel) (err error) {
	err = tx.Update(cm)
	return
}

func DeleteClassify(tx model.IClassifyModel, id uint) (err error) {
	err = tx.Delete(id)
	return
}
