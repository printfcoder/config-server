package gorm

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/micro-in-cn/config-server/config-srv/domain/model"
	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
)

type globalRepo struct{}

var _ repository.GlobalRepository = &globalRepo{}

var ErrIllegalID = errors.New("illegal id")

func (g *globalRepo) CreateApp(appName string) (*model.App, error) {
	app := &model.App{Name: appName}

	err := db.Table("app").Create(app).Error
	return app, err
}

func (g *globalRepo) CreateCluster(appName, clusterName string) (*model.Cluster, error) {
	cluster := &model.Cluster{
		AppName: appName,
		Name:    clusterName,
	}

	err := db.Table("cluster").Create(cluster).Error
	return cluster, err
}

func (g *globalRepo) CreateNamespace(appName, clusterName, namespaceName string) (*model.Namespace, error) {
	namespace := &model.Namespace{
		AppName:     appName,
		ClusterName: clusterName,
		Name:        namespaceName,
	}

	err := db.Table("namespace").Create(namespace).Error
	return namespace, err
}

func (g *globalRepo) UpdateItems(appName, clusterName, namespaceName string, del []*model.Item, update []*model.Item, insert []*model.Item) error {
	var err error
	tx := db.Begin()

	if len(del) != 0 {
		if err = deleteItems(tx, del); err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(update) != 0 {
		if err = updateItems(tx, update); err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(insert) != 0 {
		if err = insertItems(tx, appName, clusterName, namespaceName, insert); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (g *globalRepo) ListApps(appNames ...string) (apps []*model.App, err error) {
	for _, v := range appNames {
		var app model.App
		if err = db.Table("app").Where("name = ?", v).First(&app).Error; err != nil {
			return nil, err
		}

		apps = append(apps, &app)
	}

	return
}

func deleteItems(tx *gorm.DB, del []*model.Item) error {
	for _, v := range del {
		if v.ID == 0 {
			return ErrIllegalID
		}

		if err := tx.Table("item").Delete(v).Error; err != nil {
			return err
		}
	}
	return nil
}

func updateItems(tx *gorm.DB, update []*model.Item) error {
	for _, v := range update {
		if v.ID == 0 {
			return ErrIllegalID
		}

		err := tx.Table("item").Where("id = ?", v.ID).Updates(
			map[string]string{
				"key":   v.Key,
				"value": v.Value,
			}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func insertItems(tx *gorm.DB, appName, clusterName, namespaceName string, insert []*model.Item) error {
	for _, v := range insert {
		v.AppName = appName
		v.ClusterName = clusterName
		v.NamespaceName = namespaceName
		if err := tx.Table("item").Create(v).Error; err != nil {
			return err
		}
	}
	return nil
}
