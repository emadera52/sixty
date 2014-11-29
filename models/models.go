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

// Packate models defines types and methods used to interact with user
// input data and data associated with a table in the DB
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User struct used to read and write to DB table "users"
// see const category in home.go for Interest codes
type User struct {
	Id       int       // Database primary key. AutoIncrement value
	Username string    `orm:"size(30);unique"`
	Fullname string    `orm:"size(60)"`
	Email    string    `orm:"size(256)"` // encoded email address
	Password string    `orm:"size(128)"` // password hash value
	AutoLog  bool      // true if user selected "Remember Me"
	Interest uint8     // User's primary interest as a category code
	RegKey   string    `orm:"size(60)"` // used to confirm registration email
	ResetKey string    `orm:"size(60)"` // used to confirm password reset
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

// User database CRUD methods include Insert, Read, Update and Delete
func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}

// RegFrm struct to hold Registration page and Profile page form values
type RegFrm struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Fullname string `form:"fullname"`
	Password string `form:"pw1"`
	Confirm  string `form:"pw2"`
	AutoLog  string `form:"autolog"`
}

// Comment struct used to read and write to DB table "comment"
// see const category in home.go for Category codes
type Comment struct {
	Id       int    // Database primary key. AutoIncrement value
	User     *User  `orm:"rel(fk);index"` // Indexed Foreign Key -> User.Id
	Txt      string `orm:"size(4096)"`
	Category uint8  // Indicates page used to post the comment
	Archived bool
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

// Comment database access methods include Insert and Read
// TODO Add Archive and review need for Delete
func (cmnt *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(cmnt); err != nil {
		return err
	}
	return nil
}

func (cmnt *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(cmnt, fields...); err != nil {
		return err
	}
	return nil
}

// Register User and Comment models with the Beego ORM
func init() {
	orm.RegisterModel(new(User), new(Comment))
}
