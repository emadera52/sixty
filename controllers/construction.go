/* Copyright 2014 60plusadventures Author. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

// Package controllers defines controllers for each router path
package controllers

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
)

// type ConstructionController declares a receiver for methods that define
// action common to all pages that are under construction
type ConstructionController struct {
	beego.Controller
}

// Local method activeContent defines an active page layout. For example
// pages under construction are described by comment-layout.tpl
// 	{{.Header}}
// 	{{.LayoutContent}}
// 	{{.Footer}}
// Note {{.LayoutContent}} is replaced by template content as specified by view
func (cp *ConstructionController) activeContent(view string) {
	cp.Layout = "comment-layout.tpl"
	cp.LayoutSections = make(map[string]string)
	cp.LayoutSections["Header"] = "header.tpl"
	cp.LayoutSections["Footer"] = "footer.tpl"
	cp.TplNames = view + ".tpl"

	cp.Data["Website"] = SiteTitle
}

// Construction controls the presentation of pages under construction"
func (cp *ConstructionController) Construction() {
	// Parse the RequestURI
	uParm, err := url.Parse(cp.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Construction GET: Error parsing URL") // TODO Error logging
		return
	}
	// Getquery map
	qm, _ := url.ParseQuery(uParm.RawQuery)

	// Same active content template for all pages
	cp.activeContent("construction")

	// URL params indicate the requested page
	// e.g. a param of "about=y" indicates the About Us page
	// If key "about" is found, init TagTxt and display the page
	if _, ok := qm["about"]; ok {
		cp.Data["TagTxt"] = "The About Us page is not quite ready yet"
		return
	}
	if _, ok := qm["blog"]; ok {
		cp.Data["TagTxt"] = "The 60+ Blog isn't ready yet"
		return
	}
	if _, ok := qm["faq"]; ok {
		cp.Data["TagTxt"] = "The FAQ page will be ready soon"
		return
	}
	if _, ok := qm["facebook"]; ok {
		cp.Data["TagTxt"] = "Our Facebook page is not ready yet"
		return
	}
	if _, ok := qm["twitter"]; ok {
		cp.Data["TagTxt"] = "Oops. Nothing to see yet on Twitter"
		return
	}
	if _, ok := qm["email"]; ok {
		cp.Data["TagTxt"] = "We're working hard on the Email system"
		return
	}
}
