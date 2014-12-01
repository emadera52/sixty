package controllers

import (
	"fmt"
	"html/template"
	"net/url"

	"github.com/astaxie/beego"

	"sixty/crypto"
	"sixty/models"
)

const (
	// Program name for title
	ProgTitle     = "60+ Adventures"
	DenyAutoLogin = "Another user has already enabled auto-login on this device"
)

// Define category code constants
const (
	Unknown = iota
	Adventurer
	Advocate
	Provider
	Supporter
	Home
	Faq
	Blog
	About
)

// type Home Controller declares a receiver for methods that define
// action on the home page. This includes all user auth actions
type HomeController struct {
	beego.Controller
}

/*  Local method activeContent defines an active page layout.
 *  For example the home page as described by landing-layout.tpl
 * 	{{.Header}}
 * 	{{.LayoutContent}}
 * 	{{.Home}}
 * 	{{.Footer}}
 * Note {{.LayoutContent}} is replaced by template content specified by view
 */
func (hc *HomeController) activeContent(view string) {
	hc.Layout = "landing-layout.tpl"
	hc.LayoutSections = make(map[string]string)
	hc.LayoutSections["Header"] = "header.tpl"
	hc.LayoutSections["Home"] = "home.tpl"
	hc.LayoutSections["Footer"] = "footer.tpl"
	hc.TplNames = view + ".tpl"

	hc.Data["Website"] = ProgTitle
	hc.Data["xsrftoken"] = template.HTML(hc.XsrfFormHtml())
	hc.Data["IsHome"] = true
}

// Make user values available between HomeController functions
var usr models.RegFrm

/* Get selects the active content handler for "/" (the landing page)
 * it presents and processes the selected template
 */
func (hc *HomeController) Get() {
	// A session will exist if the user has Signed in or Registered
	// If a session exists, session values are used to initialize
	// controler context Data needed to render the Welcome view
	// Otherwise the Unknown view is rendered unless an auto login cookie exists
	// in which case we login the user as if they had signed in
	sess := hc.GetSession("sixty")
	if sess != nil {
		hc.Data["InSession"] = 1            // indicate that user has registered
		sm := sess.(map[string]interface{}) // and init controller Data values
		hc.Data["Username"] = sm["username"]
		hc.Data["Interest"] = CatCodeToString(sm["interest"].(uint8))
		hc.activeContent("welcome") // landing page for logged in user
	} else {
		hc.Data["InSession"] = ""   // indicate that user has not registered
		hc.activeContent("unknown") // landing page for unknown user
		UsrClear(&usr)              // clear Registration input values
		// check for auto login cookie
		usrNm, _ := crypto.DecodedCookieValue(hc.Ctx.ResponseWriter,
			hc.Ctx.Request, "alTkn")

		// If cookie found with a valid username, read user's record from DB
		// If successful, create a session and set activeContent to welcome
		// Otherwise continue with the "unknown" template
		if usrNm != nil {
			user := models.User{Username: string(usrNm)}
			if err := user.Read("Username"); err == nil {
				// decode email and create session
				var emAddr string // used for plain text email address
				// Decrypt email address
				if ptAddr, err := crypto.DecryptEmailAddr(user.Email); err != nil {
					fmt.Println("Database error: ", err.Error()) // TODO ???
				} else {
					emAddr = ptAddr
				}

				// Create "sixty" session initialized with user DB values
				sm := make(map[string]interface{})
				sm["email"] = emAddr // unencrypted email address
				sm["username"] = user.Username
				sm["fullname"] = user.Fullname
				sm["autolog"] = user.AutoLog
				sm["interest"] = user.Interest
				hc.SetSession("sixty", sm)
				hc.Data["InSession"] = 1 // indicate that user has registered
				hc.Data["Username"] = sm["username"]
				hc.Data["Interest"] = CatCodeToString(sm["interest"].(uint8))
				hc.activeContent("welcome") // landing page for logged in user
			}
		}
	}

	// Refresh flash content displayed by active content tpl after redirect
	flash := beego.ReadFromRequest(&hc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		hc.Data["notice"] = fn
	}
}

/* Register() is the active content handler for "/register"
 * it presents and processes the registration form
 */
func (hc *HomeController) Register() {
	hc.activeContent("register")

	// Parse the RequestURI
	u, err := url.Parse(hc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Register GET: Error parsing URL") // TODO Add error logging
		return
	}
	// If user clicked "Clear Form" Request params include query "clear=y"
	// Parse queries looking for key of "clear". If found, clear input values
	// Note that if key exists, value will always be "y"
	qm, _ := url.ParseQuery(u.RawQuery)
	if _, ok := qm["clear"]; ok {
		UsrClear(&usr) // clear Registration input values
	}

	// If user clicked "Cancel" Request params include query "cancel=y"
	// Parse queries looking for key of "cancel". If found, redirect to home
	// Note that if key exists, value will always be "y"
	if _, ok := qm["cancel"]; ok {
		hc.Redirect("/", 302)
		return
	}

	// If Register input form cleared or not posted,
	// update template Data with placeholder text
	// Otherwise update template Data with known input values
	if usr.Username == "" {
		hc.Data["IsInput"] = ""
		hc.Data["PhEmail"] = template.HTMLAttr(
			`placeholder="Required: Email address: e.g. youremail@anymail.com"`)
		hc.Data["PhUsername"] = template.HTMLAttr(
			`placeholder="Required: Public screen name. e.g. LuckyLou42"`)
		hc.Data["PhFullname"] = template.HTMLAttr(
			`placeholder="Optional: Enter your full name"`)
	} else {
		hc.Data["IsInput"] = 1
		hc.Data["Email"] = template.HTMLAttr(usr.Email)
		hc.Data["Username"] = template.HTMLAttr(usr.Username)
		hc.Data["Fullname"] = template.HTMLAttr(usr.Fullname)
		hc.Data["AutoLog"] = usr.AutoLog
	}

	// Refresh flash content displayed by register.tpl after redirect
	flash := beego.ReadFromRequest(&hc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		hc.Data["notice"] = fn
	}

	if hc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		// Parse Register form input values into models.RegFrm struct
		if err := hc.ParseForm(&usr); err != nil {
			flash.Notice("Error processing input form. Registration failed")
			flash.Store(&hc.Controller)
			fmt.Println("Error parsing Register form") // TODO Add error logging
			hc.Redirect("/register", 302)
			return
		}

		// Verify that Password and Confirm match with length > 7
		if len(usr.Password) < 8 || usr.Password != usr.Confirm {
			if len(usr.Password) < 8 {
				flash.Notice("Minimum password length is 8. Please try again")
			} else {
				flash.Notice("Passwords don't match. Please try again")
			}
			flash.Store(&hc.Controller)
			hc.Redirect("/register", 302)
			return
		}

		// Create new User struct and initialize known attributes
		user := new(models.User)
		user.Username = usr.Username
		if len(usr.Fullname) > 0 {
			user.Fullname = usr.Fullname
		} else {
			user.Fullname = usr.Username
		}

		if hxEncEm, err := crypto.EncryptEmailAddr([]byte(usr.Email)); err != nil {
			flash.Notice(err.Error() + "Please try again or Clear Form")
			flash.Store(&hc.Controller)
			fmt.Println(err) // TODO Review possible errors
			hc.Redirect("/register", 302)
			return
		} else {
			user.Email = hxEncEm
		}

		// Create secure password hash value
		salt := crypto.GetRandomString(10)
		pwHash := crypto.EncryptPassword(usr.Password, salt)
		// Password stored as salt+'$'+pwHash
		user.Password = fmt.Sprintf("%s$%s", salt, pwHash)

		user.AutoLog = true
		if usr.AutoLog == "" {
			user.AutoLog = false
		}

		if user.AutoLog {
			// See if an autolog cookie already exists.
			// If NOT, create one for this user
			c, _ := hc.Ctx.Request.Cookie("alTkn")
			if c == nil {
				// Create autolog cookie good for two years
				crypto.CreateEncodedCookie(hc.Ctx.ResponseWriter, "alTkn",
					user.Username, 17520, true)
			} else {
				user.AutoLog = false
				flash.Notice(DenyAutoLogin)
				flash.Store(&hc.Controller)
			}
		}

		// Create UUID value for confirmation email & set PW reset key to blank
		user.RegKey = crypto.NewV4UUID()
		user.ResetKey = ""

		// Insert new record in user table. If error assume record already exists.
		if err := user.Insert(); err != nil {
			flash.Notice("The Screen Name " + user.Username +
				" has been taken. Please try again")
			flash.Store(&hc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			hc.Redirect("/register", 302)
			return
		}

		// Insert dummy Comment record to avoid error on delete if no comments
		comment := models.Comment{Txt: `Registered`,
			Category: 0, Archived: false}
		comment.User = user
		if err := comment.Insert(); err != nil {
			flash.Notice("Data base error: unable to create Comment for " +
				user.Username)
			flash.Store(&hc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			hc.Redirect("/register", 302)
			return
		}

		// Create "sixty" session and init with shared values
		sm := make(map[string]interface{})
		sm["username"] = user.Username
		sm["fullname"] = user.Fullname
		sm["email"] = usr.Email // unencrypted email address
		sm["autolog"] = user.AutoLog
		sm["interest"] = uint8(Unknown)
		hc.SetSession("sixty", sm)
		hc.Redirect("/", 302)
	}
}

/* Profile() is the active content handler for "/profile"
 * it presents and processes the profile form
 */
func (hc *HomeController) Profile() {
	hc.activeContent("profile")

	// Active session required
	sess := hc.GetSession("sixty")
	if sess == nil {
		hc.Redirect("/restricted?profile=n", 302)
		return
	}
	sm := sess.(map[string]interface{})

	// init local Data values used by profile.tpl
	hc.Data["Email"] = sm["email"]
	hc.Data["Username"] = sm["username"]
	hc.Data["Fullname"] = sm["fullname"]
	hc.Data["AutoLog"] = sm["autolog"]

	// Parse the RequestURI
	uparm, err := url.Parse(hc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Profile GET: Error parsing URL") // TODO Add error logging
		return
	}

	// If user clicked "Delete" Request params include "delete=y"
	// Parse queries for key = "delete". If found, redirect to confirm delete
	qm, _ := url.ParseQuery(uparm.RawQuery)
	if _, ok := qm["delete"]; ok {
		hc.Redirect("/delete?display=y", 302)
		return
	}

	// If user clicked "Cancel" Request params include "cancel=y"
	// Parse queries key = "cancel". If found, redirect to home
	if _, ok := qm["cancel"]; ok {
		hc.Redirect("/", 302)
		return
	}

	// Refresh flash content displayed by profile.tpl after redirect
	flash := beego.ReadFromRequest(&hc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		hc.Data["notice"] = fn
	}

	// Process Profile values submitted by user
	if hc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		req := hc.Ctx.Request
		// Parse Profile form input values
		if err := req.ParseForm(); err != nil {
			flash.Notice("Error processing Profile form. Please try again.")
			flash.Store(&hc.Controller)
			fmt.Println(err) // TODO Add error logging
			hc.Redirect("/profile", 302)
			return
		}

		// Get form input values to be used for comparison to prev values
		// Note: The registration form struct matches the profile form struct
		var prof models.RegFrm
		prof.Email = req.FormValue("email")
		prof.Username = req.FormValue("username")
		prof.Fullname = req.FormValue("fullname")
		prof.Password = req.FormValue("pw1")
		prof.Confirm = req.FormValue("pw2")
		prof.AutoLog = req.FormValue("autolog")

		// Load user's current data from DB using session username
		// continue on success otherwise fail with flash message
		user := models.User{Username: sm["username"].(string)}
		if err := user.Read("Username"); err != nil {
			flash.Notice("Database error: Profile data not found. Please try again.")
			flash.Store(&hc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			hc.Redirect("/profile", 302)
			return
		}
		needUpdate := false

		// if Username changed, update user data
		if prof.Username != user.Username {
			user.Username = prof.Username
			needUpdate = true
		}
		// if Fullname changed, update user data
		if prof.Fullname != user.Fullname {
			if len(prof.Fullname) > 0 {
				user.Fullname = prof.Fullname
			} else {
				user.Fullname = prof.Username
			}
			needUpdate = true
		}

		// If a new email address was entered, encrypt and update user data
		if prof.Email != sm["email"].(string) {
			if hxEncEm, err := crypto.EncryptEmailAddr([]byte(prof.Email)); err != nil {
				flash.Notice(err.Error() + "Please try again")
				flash.Store(&hc.Controller)
				fmt.Println(err) // TODO Review possible errors
				hc.Redirect("/profile", 302)
				return
			} else {
				user.Email = hxEncEm
				needUpdate = true
			}
		}

		// check to see if a new, valid password was entered and confirmed
		// If new password entered, verify match with length > 7
		if len(prof.Password) > 0 {
			if len(prof.Password) < 8 || prof.Password != prof.Confirm {
				if len(prof.Password) < 8 {
					flash.Notice("Minimum password length is 8. Please try again")
				} else {
					flash.Notice("Passwords don't match. Please try again")
				}
				flash.Store(&hc.Controller)
				hc.Redirect("/profile", 302)
				return
			}
			// Create secure password hash value and update user data
			salt := crypto.GetRandomString(10)
			pwHash := crypto.EncryptPassword(prof.Password, salt)
			// Password stored as salt+'$'+pwHash
			user.Password = fmt.Sprintf("%s$%s", salt, pwHash)
			needUpdate = true
		}
		// See if an autolog cookie already exists. If NOT, this user
		// may change autolog status from false to true.
		// Otherwise, this user may change autolog status from true to false.
		c, _ := req.Cookie("alTkn")
		if c == nil {
			// No cookie. Confirm user status changed from false to true.
			// If so, create cookie and mark for update
			if prof.AutoLog != "" {
				if user.AutoLog == false {
					user.AutoLog = true
					// Create autolog cookie good for two years
					crypto.CreateEncodedCookie(hc.Ctx.ResponseWriter, "alTkn",
						user.Username, 17520, true)
					needUpdate = true
				}
			}
		} else {
			// Cookie found. Confirm user status changed from true to false.
			// If so, delete cookie and mark for update
			if prof.AutoLog == "" {
				if user.AutoLog == true {
					user.AutoLog = false
					crypto.DeleteCookie(hc.Ctx.ResponseWriter, "alTkn")
					needUpdate = true
				}
			} else {
				// A second user tried to enable auto-log. Create flash message
				flash.Notice(DenyAutoLogin)
				flash.Store(&hc.Controller)
			}
		}

		// If user changed anything, update user's DB record
		if needUpdate {
			if err := user.Update(); err != nil {
				flash.Notice("Database error: Unable to update Profile for " +
					user.Username)
				flash.Store(&hc.Controller)
				fmt.Println(err) // TODO Review DB error handling
				hc.Redirect("/profile", 302)
				return
			}
			// Delete existing session and create one with updated values
			hc.DelSession("sixty")
			sm := make(map[string]interface{})
			sm["username"] = user.Username
			sm["fullname"] = user.Fullname
			sm["email"] = prof.Email // unencrypted email address
			sm["autolog"] = user.AutoLog
			sm["interest"] = user.Interest
			hc.SetSession("sixty", sm)
		}
		hc.Redirect("/", 302)
	}
}

/* Login() is the active content handler for "/login"
 * it presents and processes the login form
 */
func (hc *HomeController) Login() {
	hc.activeContent("login")

	// Parse the RequestURI
	u, err := url.Parse(hc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Register GET: Error parsing URL") // TODO Review possible errors
		return
	}
	// If user clicked "Cancel" Request params include query "cancel=y"
	// Parse queries looking for key of "cancel". If found, redirect to home
	qm, _ := url.ParseQuery(u.RawQuery)
	if _, ok := qm["cancel"]; ok {
		hc.Redirect("/", 302)
		return
	}

	// Refresh flash content displayed by login.tpl
	flash := beego.ReadFromRequest(&hc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		hc.Data["notice"] = fn
	}

	// Process Login values submitted by user when method is POST
	if hc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()

		if err := hc.Ctx.Request.ParseForm(); err != nil {
			flash.Notice("Error processing Login form")
			flash.Store(&hc.Controller)
			fmt.Println(err) // TODO Review possible errors
			hc.Redirect("/", 302)
			return
		}

		//	Try to get user's DB record using input username
		//	continue on success otherwise fail with flash message
		logUsr := hc.Ctx.Request.Form.Get("username")
		user := models.User{Username: logUsr}
		if err := user.Read("Username"); err != nil {
			flash.Notice("Database error: Screen name not found. Please try again.")
			flash.Store(&hc.Controller)
			fmt.Println(err) // TODO Review DB error handling
			hc.Redirect("/login", 302)
			return
		}

		// Confirm that input password hashes to stored pw hash value
		logPw := hc.Ctx.Request.Form.Get("password")
		pwHash := crypto.EncryptPassword(logPw, user.Password[:10])
		if pwHash != user.Password[11:] {
			flash.Notice("Incorrect password entered. Please try again.")
			flash.Store(&hc.Controller)
			hc.Redirect("/login", 302)
			return
		}

		var emAddr string // to be used for plain text email address
		// Decrypt email address
		if ptAddr, err := crypto.DecryptEmailAddr(user.Email); err != nil {
			flash.Notice("Database error: " + err.Error())
			flash.Store(&hc.Controller)
			hc.Redirect("/login", 302)
		} else {
			emAddr = ptAddr
		}

		// Create "sixty" session, init with user session values, return to home
		sm := make(map[string]interface{})
		sm["email"] = emAddr // unencrypted email address
		sm["username"] = user.Username
		sm["fullname"] = user.Fullname
		sm["autolog"] = user.AutoLog
		sm["interest"] = user.Interest
		hc.SetSession("sixty", sm)
		hc.Redirect("/", 302)
		return
	}
}

/* Logout() is the handler for dummy template "/logout"
 * it deletes the current user session and sets activeContent to "/unknown"
 * bypassing the auto-login check on the "/" (home) landing page.
 */
func (hc *HomeController) Logout() {
	hc.DelSession("sixty")      // delete current session
	hc.Data["InSession"] = ""   // indicate that user has not registered
	hc.activeContent("unknown") // set up landing page for unknown user
	UsrClear(&usr)              // clear Registration input values
}

/* Restricted() is the active content handler for dummy template "/restricted"
 * it creates a flash notice and redirects to home ("/"
 */
func (hc *HomeController) Restricted() {
	var rPage string = "" // string identifying restricted page

	// Parse the RequestURI
	u, err := url.Parse(hc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Restricted GET: Error parsing URL") // TODO Review possible errors
		return
	}

	// If user tried to access /profile params include query "profile=n"
	// Parse queries looking for key of "profile". If found, update rPage
	qm, _ := url.ParseQuery(u.RawQuery)
	if _, ok := qm["profile"]; ok {
		rPage = "the Profile page."
	}

	// If user tried to access /delete params include query "delete=n"
	// Parse queries looking for key of "delete". If found, update rPage
	if _, ok := qm["delete"]; ok {
		rPage = "the Delete page."
	}

	flash := beego.NewFlash()
	flash.Notice("You must be signed in to access " + rPage)
	flash.Store(&hc.Controller)
	hc.Redirect("/", 302)
}

/* Delete() is the active content handler for "/delete"
 * it presents and processes confirmation of a delete profile request
 */
func (hc *HomeController) Delete() {
	hc.activeContent("delete")

	// Active session required
	sess := hc.GetSession("sixty")
	if sess == nil {
		hc.Redirect("/restricted?delete=n", 302)
		return
	}
	sm := sess.(map[string]interface{}) // init local Data values

	hc.Data["Username"] = sm["username"]
	hc.Data["Autolog"] = sm["autolog"]
	hc.Data["IsDelete"] = false

	// Parse the RequestURI
	uPrms, err := url.Parse(hc.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Delete GET: Error parsing URL") // TODO Review possible errors
		return
	}

	// If redirected from Profile page, Request params include "display=y"
	// If "display" key found, return will display delete.tpl.
	qm, _ := url.ParseQuery(uPrms.RawQuery)
	if _, ok := qm["display"]; ok {
		hc.Data["IsDelete"] = true
		return
	}

	// If user confirmed deletion, params include "delete=y"
	// If delete key found, delete records, otherwise redirect to /profile
	if _, ok := qm["delete"]; ok {
		//Deletion confirmed
		flash := beego.NewFlash()

		//Must read user's DB record as Delete requires primary key
		user := models.User{Username: sm["username"].(string)}
		if err := user.Read("Username"); err != nil {
			flash.Notice("Database error: Screen name not found for " +
				sm["username"].(string))
			flash.Store(&hc.Controller)
			fmt.Println("643 ", err) // TODO Review DB error handling
			hc.Redirect("/profile", 302)
			return
		}
		// Delete this user's record which will cascade to delete all comments
		// Note that there MUST be at least one comment (see Register)
		if err := user.Delete(); err != nil {
			flash.Notice("Database error: Deleting user " + sm["username"].(string))
			flash.Store(&hc.Controller)
			fmt.Println("652 ", err) // TODO Review DB error handling
			hc.Redirect("/profile", 302)
			return
		}

		// If user's autolog was enabled, mark autolog token for deletion
		if hc.Data["Autolog"].(bool) {
			crypto.DeleteCookie(hc.Ctx.ResponseWriter, "alTkn")
		}

		// Delete this user's session, create success message and return home
		hc.DelSession("sixty")
		flash.Notice("All information for " + sm["username"].(string) +
			" has been deleted.")
		flash.Store(&hc.Controller)
		hc.Redirect("/", 302)
		return
	} else {
		//Delete not confirmed
		hc.Redirect("/profile", 302)
	}
}

// UsrClear clears form input data from a models.RegFrm struct
func UsrClear(usrInp *models.RegFrm) {
	usrInp.Email = ""
	usrInp.Username = ""
	usrInp.Fullname = ""
	usrInp.Password = ""
	usrInp.Confirm = ""
	usrInp.AutoLog = ""
}

// CatCodeToString converts a const category code to a string
func CatCodeToString(code uint8) (catStr string) {
	switch code {
	case 1:
		catStr = "Adventurer"
	case 2:
		catStr = "Advocate"
	case 3:
		catStr = "Provider"
	case 4:
		catStr = "Supporter"
	default:
		catStr = ""
	}
	return catStr
}
