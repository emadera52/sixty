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

// type BlogController declares a receiver for methods that define
// action common to all blog pages
type BlogController struct {
	beego.Controller
}

// Local method activeContent defines an active page layout. For example
// blog pages are described by comment-layout.tpl
// 	{{.Header}}
// 	{{.LayoutContent}}
// 	{{.Footer}}
// Note {{.LayoutContent}} is replaced by template content as specified by view
func (bc *BlogController) activeContent(view string) {
	bc.Layout = "comment-layout.tpl"
	bc.LayoutSections = make(map[string]string)
	bc.LayoutSections["Header"] = "header.tpl"
	bc.LayoutSections["Footer"] = "footer.tpl"
	bc.TplNames = view + ".tpl"

	// Context Data common to all blog pages
	bc.Data["Website"] = SiteTitle
	bc.Data["IsBlog"] = true
	bc.Data["TagTxt"] = "60+ Adventures Blog"
}

// Blog controls the presentation of blog pages"
func (bc *BlogController) Blog() {
	// Parse the RequestURI
	uParm, err := url.Parse(bc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Blog GET: Error parsing URL") // TODO Error logging
		return
	}
	// Get query map
	qm, _ := url.ParseQuery(uParm.RawQuery)

	// URL params indicate the requested page
	// e.g. a param of "dest=blog" indicates the blog home page
	// and a param of "dest=firstblog" indicates the article firstblog
	if page, ok := qm["dest"]; ok {
		if page[0] == "blog" {
			bc.activeContent("blog")
			return
		} else {
			bc.Data["Article"] = true
			bc.activeContent("articles/" + page[0])
			return
		}
	}
}
