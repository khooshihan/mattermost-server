// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

func (a *App) TestElasticsearch(cfg *model.Config) *model.AppError {
	if *cfg.ElasticsearchSettings.Password == model.FAKE_SETTING {
		if *cfg.ElasticsearchSettings.ConnectionUrl == *a.Config().ElasticsearchSettings.ConnectionUrl && *cfg.ElasticsearchSettings.Username == *a.Config().ElasticsearchSettings.Username {
			*cfg.ElasticsearchSettings.Password = *a.Config().ElasticsearchSettings.Password
		} else {
			return model.NewAppError("TestElasticsearch", "ent.elasticsearch.test_config.reenter_password", nil, "", http.StatusBadRequest)
		}
	}

	seI := a.SearchEngine.ElasticsearchEngine
	if seI == nil {
		err := model.NewAppError("TestElasticsearch", "ent.elasticsearch.test_config.license.error", nil, "", http.StatusNotImplemented)
		return err
	}
	if err := seI.TestConfig(cfg); err != nil {
		return err
	}

	return nil
}

func (a *App) PurgeElasticsearchIndexes() *model.AppError {
	seI := a.SearchEngine.ElasticsearchEngine
	if seI == nil {
		err := model.NewAppError("PurgeElasticsearchIndexes", "ent.elasticsearch.test_config.license.error", nil, "", http.StatusNotImplemented)
		return err
	}

	if err := seI.PurgeIndexes(); err != nil {
		return err
	}

	return nil
}

func (a *App) PurgeBleveIndexes() *model.AppError {
	seI := a.SearchEngine.BleveEngine
	if seI == nil {
		err := model.NewAppError("PurgeBleveIndexes", "bleve.not-available.error", nil, "", http.StatusNotImplemented)
		return err
	}

	if err := seI.PurgeIndexes(); err != nil {
		return err
	}

	return nil
}
