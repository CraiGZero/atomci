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

package initialize

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/go-atomci/atomci/internal/core/pipelinemgr"
	"github.com/go-atomci/atomci/internal/middleware/log"
)

func initTaskTemplates() error {

	taskTmpls := []pipelinemgr.TaskTmplReq{
		{
			Name:        "应用构建",
			Type:        "build",
			Description: "用于应用构建",
			SubTask: []pipelinemgr.SubTask{
				{
					Index: 1,
					Type:  "checkout",
					Name:  "检出代码",
				},
				{
					Index: 2,
					Type:  "compile",
					Name:  "编译",
				},
				{
					Index: 3,
					Type:  "build-image",
					Name:  "制作镜像",
				},
			},
		},
		{
			Name:        "应用部署",
			Type:        "deploy",
			Description: "用于应用部署健康检查",
		},
		{
			Name:        "人工卡点",
			Type:        "manual",
			Description: "人工卡点",
		},
	}

	pipeline := pipelinemgr.NewPipelineManager()

	for _, item := range taskTmpls {
		_, err := pipeline.GetTaskTmplByName(item.Name)
		if err != nil {
			if err == orm.ErrNoRows {
				if err := pipeline.CreateTaskTmpl(&item, "admin"); err != nil {
					log.Log.Error("when init task template, occur error: %s", err.Error())
					continue
				}
			} else {
				logs.Warn("init task template occur error: %s", err.Error())
				continue
			}
		} else {
			log.Log.Debug("component type `%s` already exists, skip", item.Name)
			continue
		}
	}
	return nil
}
