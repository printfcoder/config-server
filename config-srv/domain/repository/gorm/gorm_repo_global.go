package gorm

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/micro-in-cn/config-server/config-srv/domain/model"
	"github.com/micro/go-micro/util/log"
)

type globalRepo struct{}

var ErrIllegalID = errors.New("illegal id")

func (g *globalRepo) CreateApp(appName string) (*model.App, error) {
	app := &model.App{Name: appName}

	err := db.Table("app").Create(app).Error
	if err != nil {
		log.Error("[CreateApp] create app:%s error: %s", appName, err.Error())
	}
	return app, err
}

func (g *globalRepo) DeleteApp(appName string) error {
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			log.Error("[DeleteApp] delete app:%s error: %s", appName, err.Error())
			tx.Rollback()
		}
	}()

	if err = tx.Table("instance").Where("app_name = ?", appName).Delete(model.Instance{}).Error; err != nil {
		return err
	}
	if err = tx.Table("item").Where("app_name = ?", appName).Delete(model.Item{}).Error; err != nil {
		return err
	}
	if err = tx.Table("namespace").Where("app_name = ?", appName).Delete(model.Namespace{}).Error; err != nil {
		return err
	}
	if err = tx.Table("cluster").Where("app_name = ?", appName).Delete(model.Cluster{}).Error; err != nil {
		return err
	}
	if err = tx.Table("app").Where("name = ?", appName).Delete(model.App{}).Error; err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (g *globalRepo) ListApps(appNames ...string) (apps []*model.App, err error) {
	for _, v := range appNames {
		var app model.App
		if err = db.Table("app").Where("name = ?", v).First(&app).Error; err != nil {
			log.Error("[ListApps] list apps error: %s", err.Error())
			return nil, err
		}

		apps = append(apps, &app)
	}

	return
}

func (g *globalRepo) CreateCluster(appName, clusterName string) (*model.Cluster, error) {
	cluster := &model.Cluster{
		AppName: appName,
		Name:    clusterName,
	}

	err := db.Table("cluster").Create(cluster).Error
	if err != nil {
		log.Error("[CreateCluster] create cluster:%s-%s error: %s", appName, clusterName, err.Error())
	}
	return cluster, err
}

func (g *globalRepo) DeleteCluster(appName, clusterName string) error {
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			log.Error("[DeleteCluster] delete cluster:%s-%s error: %s", appName, clusterName, err.Error())
			tx.Rollback()
		}
	}()

	if err = tx.Table("instance").Where("app_name = ? and cluster_name = ?", appName, clusterName).
		Delete(model.Instance{}).Error; err != nil {
		return err
	}
	if err = tx.Table("item").Where("app_name = ? and cluster_name = ?", appName, clusterName).
		Delete(model.Item{}).Error; err != nil {
		return err
	}
	if err = tx.Table("namespace").Where("app_name = ? and cluster_name = ?", appName, clusterName).
		Delete(model.Namespace{}).Error; err != nil {
		return err
	}
	if err = tx.Table("cluster").Where("app_name = ? and name = ?", appName, clusterName).
		Delete(model.Cluster{}).Error; err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (g *globalRepo) ListClusters(appName string) (clusters []*model.Cluster, err error) {
	err = db.Table("cluster").Where("app_name = ?", appName).Find(&clusters).Error
	if err != nil {
		log.Error("[ListClusters] list cluster:%s error: %s", appName, err.Error())
	}
	return
}

func (g *globalRepo) CreateNamespace(appName, clusterName, namespaceName string) (*model.Namespace, error) {
	namespace := &model.Namespace{
		AppName:     appName,
		ClusterName: clusterName,
		Name:        namespaceName,
	}

	err := db.Table("namespace").Create(namespace).Error
	if err != nil {
		log.Error("[CreateNamespace] create namespace:%s-%s-%s error: %s", appName, clusterName, namespaceName, err.Error())
	}
	return namespace, err
}

func (g *globalRepo) DeleteNamespace(appName, clusterName, namespaceName string) error {
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			log.Error("[DeleteNamespace] delete namespace:%s-%s-%s error: %s", appName, clusterName, namespaceName, err.Error())
			tx.Rollback()
		}
	}()

	if err = tx.Table("instance").Where("app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).Delete(model.Instance{}).Error; err != nil {
		return err
	}
	if err = tx.Table("item").Where("app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).Delete(model.Item{}).Error; err != nil {
		return err
	}
	if err = tx.Table("namespace").Where("app_name = ? and cluster_name = ? and name = ?",
		appName, clusterName, namespaceName).Delete(model.Namespace{}).Error; err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (g *globalRepo) ListNamespace(appName, clusterName string) (namespaces []*model.Namespace, err error) {
	err = db.Table("namespace").Where("app_name = ? and cluster_name = ?", appName, clusterName).Find(&namespaces).Error
	if err != nil {
		log.Error("[ListNamespace] list namespace:%s-%s error: %s", appName, clusterName, err.Error())
	}
	return
}

func (g *globalRepo) ListItems(appName, clusterName, namespaceName string) (items []*model.Item, err error) {
	err = db.Select("item").Where("app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).Find(&items).Error
	if err != nil {
		log.Error("[ListItems] list items:%s-%s-%s error: %s", appName, clusterName, namespaceName, err.Error())
		return
	}
	return
}

func (g *globalRepo) UpdateItems(appName, clusterName, namespaceName string, del []*model.Item, update []*model.Item, insert []*model.Item) error {
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			log.Error("[UpdateItems] update items:%s-%s-%s error: %s", appName, clusterName, namespaceName, err.Error())
			tx.Rollback()
		}
	}()

	if len(del) != 0 {
		if err = deleteItems(tx, appName, clusterName, namespaceName, del); err != nil {
			return err
		}
	}

	if len(update) != 0 {
		if err = updateItems(tx, appName, clusterName, namespaceName, update); err != nil {
			return err
		}
	}

	if len(insert) != 0 {
		if err = insertItems(tx, appName, clusterName, namespaceName, insert); err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

func deleteItems(tx *gorm.DB, appName, clusterName, namespaceName string, del []*model.Item) error {
	for _, v := range del {
		if v.ID == 0 {
			return ErrIllegalID
		}

		if err := tx.Table("item").Where("app_name = ? and cluster_name = ? and namespace_name = ? and key = ?",
			appName, clusterName, namespaceName, v.Key).Delete(v).Error; err != nil {
			return err
		}
	}
	return nil
}

func updateItems(tx *gorm.DB, appName, clusterName, namespaceName string, update []*model.Item) error {
	for _, v := range update {
		if v.ID == 0 {
			return ErrIllegalID
		}

		err := tx.Table("item").Where("app_name = ? and cluster_name = ? and namespace_name = ? and key = ?",
			appName, clusterName, namespaceName, v.Key).Updates(
			map[string]string{
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
