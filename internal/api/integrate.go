/*
Copyright 2021 The AtomCI Group Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"github.com/go-atomci/atomci/internal/core/settings"
	"github.com/go-atomci/atomci/internal/middleware/log"
)

// IntegrateController ...
type IntegrateController struct {
	BaseController
}

func (p *IntegrateController) GetClusterIntegrateSettings() {
	pm := settings.NewSettingManager()
	rsp, err := pm.GetIntegrateSettings("kubernetes")
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("Get integrate settings occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, rsp, "")
	p.ServeJSON()
}

// GetIntegrateSettings ..
func (p *IntegrateController) GetIntegrateSettings() {
	pm := settings.NewSettingManager()
	rsp, err := pm.GetIntegrateSettings("")
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("Get integrate settings occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, rsp, "")
	p.ServeJSON()
}

// GetIntegrateSettingsByPagination ..
func (p *IntegrateController) GetIntegrateSettingsByPagination() {
	filterQuery := p.GetFilterQuery()
	pm := settings.NewSettingManager()
	rsp, err := pm.GetIntegrateSettingsByPagination(filterQuery)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("Get integrate settings occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, rsp, "")
	p.ServeJSON()
}

// CreateIntegrateSetting ..
func (p *IntegrateController) CreateIntegrateSetting() {
	request := settings.IntegrateSettingReq{}
	creator := p.User
	p.DecodeJSONReq(&request)
	pm := settings.NewSettingManager()
	err := pm.CreateIntegrateSetting(&request, creator)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("Create integrate setting occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}

// VerifyIntegrateSetting ..
func (p *IntegrateController) VerifyIntegrateSetting() {
	request := settings.IntegrateSettingReq{}
	p.DecodeJSONReq(&request)
	pm := settings.NewSettingManager()
	resp := pm.VerifyIntegrateSetting(&request)
	if resp.Error != nil {
		p.HandleInternalServerError(resp.Error.Error())
		log.Log.Error("verify integrate setting occur error: %s", resp.Error.Error())
		return
	}
	p.Data["json"] = NewResult(true, resp.Msg, "")
	p.ServeJSON()
}

// UpdateIntegrateSetting ..
func (p *IntegrateController) UpdateIntegrateSetting() {
	stageID, _ := p.GetInt64FromPath(":id")
	request := settings.IntegrateSettingReq{}
	p.DecodeJSONReq(&request)
	pm := settings.NewSettingManager()
	err := pm.UpdateIntegrateSetting(&request, stageID)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("update integrate setting occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}

// DeleteIntegrateSetting ..
func (p *IntegrateController) DeleteIntegrateSetting() {
	itemID, _ := p.GetInt64FromPath(":id")
	pm := settings.NewSettingManager()
	err := pm.DeleteIntegrateSetting(itemID)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("delete integrate setting occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}

// GetCompileEnvs ..
func (p *IntegrateController) GetCompileEnvs() {
	pm := settings.NewSettingManager()
	rsp, err := pm.GetCompileEnvs("")
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("Get compile envs occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, rsp, "")
	p.ServeJSON()
}

// GetCompileEnvsByPagination ..
func (p *IntegrateController) GetCompileEnvsByPagination() {
	filterQuery := p.GetFilterQuery()
	pm := settings.NewSettingManager()
	rsp, err := pm.GetCompileEnvsByPagination(filterQuery)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("Get compile envs occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, rsp, "")
	p.ServeJSON()
}

// CreateCompileEnv ..
func (p *IntegrateController) CreateCompileEnv() {
	request := settings.CompileEnvReq{}
	creator := p.User
	p.DecodeJSONReq(&request)
	pm := settings.NewSettingManager()
	err := pm.CreateCompileEnv(&request, creator)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("Create compile env occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}

// UpdateCompileEnv ..
func (p *IntegrateController) UpdateCompileEnv() {
	stageID, _ := p.GetInt64FromPath(":id")
	request := settings.CompileEnvReq{}
	p.DecodeJSONReq(&request)
	pm := settings.NewSettingManager()
	err := pm.UpdateCompileEnv(&request, stageID)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("update compile env occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}

// DeleteCompileEnv ..
func (p *IntegrateController) DeleteCompileEnv() {
	itemID, _ := p.GetInt64FromPath(":id")
	pm := settings.NewSettingManager()
	err := pm.DeleteCompileEnv(itemID)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("delete compile env occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}
