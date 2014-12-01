## Sixty

**sixty** is an example of a "test the waters" website built using
GO, HTML5, CSS3, Bootstrap, MySQL and the *Beego* web framework.
The only JS (not included in Bootstrap) was added to handle flash messages.

Visit the *60+* Adventures website at http://60plusadventures.com

## Installation

First you will need to set up a database. The file *setup.sql* contains SQL to create the database and the required tables assuming you use MySQL. Beego Also supports *SQL Lite* and *Postgres*, however you will need to modify the SQL as needed.

See comments in **main.go** regarding correct database registration.

Then you can...

1. Download the .zip file.
2. Confirm that the directory below GOPATH is **src**
3. Unzip to *$GOPATH/src* (linux) or *%GOPATH%\src* (windows)
4. Rename the unzipped directory from *sixty-master* to **sixty**
5. Make required changes to *main.go* for correct database access
6. From the **sixty** directory enter **go build**
7. Assuming no errors, enter **./sixty** (linux) or **sixty** (windows)
8. In a browser go to localhost:12491/ to see the site
9. If necessary, you can change the port in *sixty/conf/app.conf*

## Features

* Non-SSL User Authentication
* CSRF protection
* Extensive use of **Go** templates to avoid JS
* Encoding, Encrypting and Hashing examples
* Based on Beego's MVC architecture: http://beego.me/docs/mvc/
* Uses Beego's ORM for Database access
* Demonstrates a simple 1:many Database relationship
* Uses Beego's per request *context* along with persistent *sessions*
* Demonstrates *bootstap's* responsive grid. Usable phone to desktop 
* App can be used as a template for gaging public interest in any idea

## Documentation (Technical)

* http://godoc.org/github.com/emadera52/sixty

Note that all code is vendored. Using Beego version 1.4.1

## Fair Warning

* This is my first **Go** project
* This is the first project I've published on GitHub
* Constructive suggestions and criticism solicited.

## TODO

* Make project go gettable (per on feedback from eliot5)
* Layout changes etc (per feedback from egonelbre on reddit)
* Finish all *under construction* pages
* Create a demo *destination* site with video, ads, etc.

## LICENSE

**sixty** is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).

Individual source files may contain additional license
information regarding included third party code
