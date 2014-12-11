## Sixty

**sixty** is an example of a "test the waters" website built using
Go, HTML5, CSS3, Bootstrap, MySQL and the *Beego* web framework.
The only JS (not included in Bootstrap) was added to handle flash messages.

Visit the *60+* Adventures website at http://60plusadventures.com

## Installation

### Requirements
Everything below assumes you have installed Go and defined **$GOPATH** (linux) or **%GOPATH%** (windows). [This document](https://golang.org/doc/code.html#GOPATH) explains GOPATH setup. [This site](http://www.computerhope.com/issues/ch000549.htm) explains how windows users can create **%GOPATH%** and edit **PATH**. 

### Database Setup
The file *setup.sql* contains SQL to create the database and the required tables (assuming you use MySQL). Beego Also supports *SQL Lite* and *Postgres*, however you will need to modify the SQL as needed.

See comments in **main.go** regarding correct database registration.

### Installation

	go get github.com/emadera52/sixty

This will install the application executable **sixty** (linux) or **sixty.exe** (windows) in *GOPATH/bin* which is why it's handy to add that to your *PATH*.

## Get the Source Code

From the **master** branch, click **Clone in Desktop** if you don't plan to submit updates to the project.

To **Fork** this project with the intent of contributing bug fixes,
modification or new stuff follow the [instructions here](https://help.github.com/articles/fork-a-repo/)

Click **Download ZIP** from master (more stable) or dev60 (not stable) to get a completely independent copy to do with as you please within the limits of the *LICENSE* (see below).

**Note:** I've made the following  modification to CheckXsrfCookie() in my vendored copy of *beego/context/context.go*

**changed**

    if token == "" {
    	ctx.Abort(403, "'_xsrf' argument missing from POST")
    } else if ctx._xsrf_token != token {
    	ctx.Abort(403, "XSRF cookie does not match POST argument")
    }
	
**to**

    if token == "" {
    	ctx.Abort(403, "403")
    } else if ctx._xsrf_token != token {
    	ctx.Abort(403, "403")
    }

Without this change, a user with cookies disabled may see a white screen, or a panic message depending on system/previous activity. With the change, a 403 error is displayed.

A more user friendly solution is to change "403" to something like "601" and modify *beego/middleware/error.go* by adding a 601 "CookiesDisabled" function and registering it in RegisterErrorHandler().

I've reported the problem and both solutions as an issue on the beego project.

## Features

* Non-SSL User Authentication
* CSRF protection
* Extensive use of **Go** templates to avoid JS
* Encoding, Encrypting and Hashing examples
* Based on Beego's MVC architecture: http://beego.me/docs/mvc/
* Uses Beego's ORM for Database access
* Demonstrates a simple 1:many Database relationship
* Uses Beego's per request *context* along with persistent *sessions*
* Demonstrates *bootstap's* responsive grid. Usable smart phone > desktop 
* Can be used as a template for gaging public interest in other ideas

## Documentation (Technical)

* http://godoc.org/github.com/emadera52/sixty

## Fair Warning

* This is my first **Go** project
* This is the first project I've published on GitHub
* Constructive suggestions and criticism solicited
* Use **Issues** to report problems, ask questions or make suggestions 

## TODO

* Make comments viewable
* Add email handling for validation, password reset and user contact
* Create a facebook page and twitter account for the project
* Create a demo *destination* site with video, ads, etc.

## LICENSE

**sixty** is licensed under the Apache Licence, Version 2.0

(http://www.apache.org/licenses/LICENSE-2.0.html).

Individual source files may contain additional license
information regarding included third party code
