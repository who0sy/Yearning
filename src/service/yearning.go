// Copyright 2019 HenryYee.
//
// Licensed under the AGPL, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"Yearning-go/src/model"
	_ "Yearning-go/src/model"
	"Yearning-go/src/router"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/cookieY/yee"
	"github.com/cookieY/yee/middleware"
	"net/http"
)

//go:embed dist/*
var f embed.FS

//go:embed dist/index.html
var html string

func StartYearning(port string, host string) {
	model.DB().First(&model.GloPer)
	model.Host = host
	_ = json.Unmarshal(model.GloPer.Message, &model.GloMessage)
	_ = json.Unmarshal(model.GloPer.Ldap, &model.GloLdap)
	_ = json.Unmarshal(model.GloPer.Other, &model.GloOther)
	_ = json.Unmarshal(model.GloPer.AuditRole, &model.GloRole)
	e := yee.New()
	e.Pack("/front", f, "dist")
	e.Use(middleware.Cors())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recovery())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 9,
	}))
	e.SetLogLevel(2)
	e.GET("/", func(c yee.Context) error {
		return c.HTML(http.StatusOK, html)
	})
	router.AddRouter(e)

	e.Run(fmt.Sprintf("%s", port))
}
