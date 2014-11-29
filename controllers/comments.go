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
	"html/template"
	"net/url"
	"strconv"

	"github.com/astaxie/beego"

	"sixty/models"
)

// HTML value indicating that a checkbox or radio button is selected
const IsChecked = "checked"

// type CommentsController declares a receiver for methods that define
// action common to all "comment pages", i.e. pages that accept user comments
type CommentsController struct {
	beego.Controller
}

// Local method activeContent defines an active page layout. for example
// "comment pages" are described by comment-layout.tpl
// 	{{.Header}}
// 	{{.LayoutContent}}
// 	{{.Footer}}
// Note {{.LayoutContent}} is replaced by template content as specified by view
func (cc *CommentsController) activeContent(view string) {
	cc.Layout = "comment-layout.tpl"
	cc.LayoutSections = make(map[string]string)
	cc.LayoutSections["Header"] = "header.tpl"
	cc.LayoutSections["Footer"] = "footer.tpl"
	cc.TplNames = view + ".tpl"

	cc.Data["Website"] = ProgTitle
	cc.Data["xsrftoken"] = template.HTML(cc.XsrfFormHtml())
}

// Comment controls the presentation and processing of all "comment pages"
func (cc *CommentsController) Comment() {
	// A session will exist if the user has Signed in or Registered
	// If found, session values are used to initialize controller Data,
	// otherwise user may browse as a guest without comment privileges
	sess := cc.GetSession("sixty")
	if sess != nil {
		cc.Data["InSession"] = 1            // indicate that user has registered
		sm := sess.(map[string]interface{}) // and init controller Data values
		cc.Data["Username"] = sm["username"]
		// Set appropriate Data variable to indicate user's primary interest
		// see comment.tpl for usage.
		switch sm["interest"].(uint8) {
		case 1:
			cc.Data["IsChk1"] = IsChecked
		case 2:
			cc.Data["IsChk2"] = IsChecked
		case 3:
			cc.Data["IsChk3"] = IsChecked
		case 4:
			cc.Data["IsChk4"] = IsChecked
		default:
			cc.Data["IsChk0"] = IsChecked
		}
	} else {
		cc.Data["InSession"] = "" // Don't allow comments (e.g. see adventurer.tpl)
	}

	// Parse the RequestURI
	uParm, err := url.Parse(cc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Comment GET: Error parsing URL") // TODO Error logging
		return
	}
	// Get query map
	qm, _ := url.ParseQuery(uParm.RawQuery)

	// URL params indicate which "comment page" was requested by the user
	// e.g. a param of "adventurer=y" should load the Adventurer page
	// Parse queries looking for one of the menu key strings
	// If found, init controller Data values and call activeContent
	if _, ok := qm["adventurer"]; ok {
		cc.Data["IsAdventurer"] = true
		cc.Data["TagTxt"] = "Information for 60+ Adventurers"
		cc.Data["CmntCat"] = Adventurer
		cc.activeContent("adventurer")
		return
	}
	if _, ok := qm["advocate"]; ok {
		cc.Data["IsAdvocate"] = true
		cc.Data["TagTxt"] = "Information for 60+ Advocates"
		cc.Data["CmntCat"] = Advocate
		cc.activeContent("advocate")
		return
	}
	if _, ok := qm["provider"]; ok {
		cc.Data["IsProvider"] = true
		cc.Data["TagTxt"] = "Information for 60+ Providers"
		cc.Data["CmntCat"] = Provider
		cc.activeContent("provider")
		return
	}
	if _, ok := qm["supporter"]; ok {
		cc.Data["IsSupporter"] = true
		cc.Data["TagTxt"] = "Information for 60+ Supporters"
		cc.Data["CmntCat"] = Supporter
		cc.activeContent("supporter")
		return
	}

	// Refresh flash content displayed by active content tpl after redirect
	flash := beego.ReadFromRequest(&cc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		cc.Data["notice"] = fn
	}

	// For POST requests process values submitted by user
	if cc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		// declare vars with func scope
		var (
			usrName string
			ptEmail string
			usrCat  uint8
		)

		sess := cc.GetSession("sixty")
		if sess != nil {
			// create session map
			sm := sess.(map[string]interface{})
			// save session values needed to process comment form input
			usrName = sm["username"].(string) // to read DB record
			ptEmail = sm["email"].(string)    // plain text email address
			usrCat = sm["interest"].(uint8)   // user's current category code
		} else {
			// This should never happen since POST not allowed if no session
			flash.Notice("You must be signed in to comment")
			flash.Store(&cc.Controller)
			cc.Redirect("/", 302)
			return
		}

		if err := cc.Ctx.Request.ParseForm(); err != nil {
			flash.Notice("Error processing Comment form")
			flash.Store(&cc.Controller)
			fmt.Println(err) // TODO Ignore or log based on possible errors
			cc.Redirect("/", 302)
			return
		}

		// Get user's DB record using session username
		user := models.User{Username: usrName}
		if err := user.Read("Username"); err != nil {
			flash.Notice("Database error: Screen name not found. Please try again.")
			flash.Store(&cc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			cc.Redirect("/", 302)
			return
		}

		usrCategory, err := strconv.Atoi(cc.Ctx.Request.Form.Get("category"))
		if err != nil {
			fmt.Println(err) // TODO Ignore or log based on possible errors
			usrCategory = 0
		}

		// If the user selected a different category update DB record
		if uint8(usrCategory) != usrCat {
			user.Interest = uint8(usrCategory)
			if err := user.Update(); err != nil {
				flash.Notice("Database error: Unable to update " + usrName)
				flash.Store(&cc.Controller)
				fmt.Println(err) // TODO Review DB error handling
				cc.Redirect("/profile", 302)
				return
			}
			// Delete existing session and create one with updated values
			cc.DelSession("sixty")
			sm := make(map[string]interface{})
			sm["username"] = user.Username
			sm["fullname"] = user.Fullname
			sm["email"] = ptEmail // unencrypted email address
			sm["autolog"] = user.AutoLog
			sm["interest"] = user.Interest
			cc.SetSession("sixty", sm)
		}

		// Get comment and add it to comments DB table if length > 0
		usrComment := cc.Ctx.Request.Form.Get("cmnttext")
		cmntCategory, err2 := strconv.Atoi(cc.Ctx.Request.Form.Get("cmntcat"))
		if err2 != nil {
			fmt.Println(err2) // TODO Ignore or log based on possible errors
			cmntCategory = 1
		}

		if len(usrComment) > 0 {
			comment := models.Comment{Txt: usrComment,
				Category: uint8(cmntCategory), Archived: false}
			comment.User = &user
			if err := comment.Insert(); err != nil {
				flash.Notice("Data base error: Unable to create Comment for " +
					usrName)
				flash.Store(&cc.Controller)
				fmt.Println(err) // TODO Review DB error handling
				cc.Redirect("/", 302)
				return
			}
		}

		// Create a flash notice indicating success.
		flash.Notice("Comment Received from: " +
			user.Username + ". Thank you for participating.")
		flash.Store(&cc.Controller)

		cc.Redirect("/", 302)
		return
	}
}
