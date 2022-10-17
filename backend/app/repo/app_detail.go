package repo

import (
	"context"
	"github.com/1Panel-dev/1Panel/backend/app/model"
	"gorm.io/gorm"
)

type AppDetailRepo struct {
}

func (a AppDetailRepo) WithVersion(version string) DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("version = ?", version)
	}
}
func (a AppDetailRepo) WithAppId(id uint) DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("app_id = ?", id)
	}
}

func (a AppDetailRepo) GetFirst(opts ...DBOption) (model.AppDetail, error) {
	var detail model.AppDetail
	err := getDb(opts...).Model(&model.AppDetail{}).Find(&detail).Error
	return detail, err
}

func (a AppDetailRepo) Update(ctx context.Context, detail model.AppDetail) error {
	return getTx(ctx).Save(&detail).Error
}

func (a AppDetailRepo) BatchCreate(ctx context.Context, details []model.AppDetail) error {
	return getTx(ctx).Model(&model.AppDetail{}).Create(&details).Error
}

func (a AppDetailRepo) DeleteByAppIds(ctx context.Context, appIds []uint) error {
	return getTx(ctx).Where("app_id in (?)", appIds).Delete(&model.AppDetail{}).Error
}

func (a AppDetailRepo) GetBy(opts ...DBOption) ([]model.AppDetail, error) {
	var details []model.AppDetail
	err := getDb(opts...).Find(&details).Error
	return details, err
}

func (a AppDetailRepo) BatchUpdateBy(maps map[string]interface{}, opts ...DBOption) error {
	db := getDb(opts...).Model(&model.AppDetail{})
	if len(opts) == 0 {
		db = db.Where("1=1")
	}
	return db.Updates(&maps).Error
}