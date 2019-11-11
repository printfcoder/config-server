package service

import (
	"fmt"

	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
	"github.com/micro-in-cn/config-server/proto/entry"
)

type App struct {
}

func (s *App) CreateApp(appId, appName string) (int64, error) {
	db := repository.GetDB()

	row, err := db.Exec("INSERT INTO app (app_id, app_name) VALUE (?, ?)", appId, appName)
	if err != nil {
		return 0, fmt.Errorf("[CreateApp] create App error: %s", err)
	}

	return row.LastInsertId()
}

func (s *App) ListApps(appIds ...string) ([]*entry.App, error) {
	db := repository.GetDB()

	rows, err := db.Query("SELECT app_id, app_name, created_time, updated_time FROM "+
		" app WHERE app_id IN (?) AND deleted = 0", appIds)
	if err != nil {
		return nil, fmt.Errorf("[ListApps] query App error: %s", err)
	}

	var apps []*entry.App
	for rows.Next() {
		item := &entry.App{}
		if err = rows.Scan(&item.AppId, &item.AppName, &item.CreatedTime, &item.UpdatedTime);
			err != nil {
			err = fmt.Errorf("[ListApps] scan err: %s", err)
			return nil, err
		}
		apps = append(apps, item)
	}

	return apps, nil
}
