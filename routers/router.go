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

// Package routers associates defined routes with controller functions
// that can handle those routes. Note that the Beego framework can handle
// complex routes, but so far this app uses only simple routes.
package routers

import (
	"github.com/astaxie/beego"
	"github.com/emadera52/sixty/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.HomeController{}, "get,post:Register")
	beego.Router("/login", &controllers.HomeController{}, "get,post:Login")
	beego.Router("/profile", &controllers.HomeController{}, "get,post:Profile")
	beego.Router("/logout", &controllers.HomeController{}, "get:Logout")
	beego.Router("/delete", &controllers.HomeController{}, "get:Delete")
	beego.Router("/restricted", &controllers.HomeController{}, "get:Restricted")
	beego.Router("/comments", &controllers.CommentsController{}, "get,post:Comment")
	beego.Router("/blog", &controllers.BlogController{}, "get:Blog")
	beego.Router("/construction", &controllers.ConstructionController{}, "get:Construction")
	// TODO add email service in order to provide email verification
	//	and password reset functionality
	//	beego.Router("/verify/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", &controllers.HomeController{}, "get:Verify")
	//	beego.Router("/forgot", &controllers.HomeController{}, "get,post:Forgot")
	//	beego.Router("/reset/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", &controllers.HomeController{}, "get,post:Reset")
}
