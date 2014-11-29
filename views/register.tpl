<div class="container" id="beta-signup">
	<div class="row col-md-12">
		{{if .flash.notice}}
		<div id="alert60" class="alert alert-info">
			<h4>{{.flash.notice}}</h4>
		</div>
		<script type="text/javascript">
		<!--
		$(document).ready(function () {
		window.setTimeout(function() {
			$(".alert").fadeTo(1500, 0).slideUp(500, function(){
				$(this).remove(); 
			});
		}, 3500);
		});
		//-->
		</script>
		{{end}}
		<div class="sign-up">
			<h3>Tell us how you want <em>60+</em> Adventures to work for you.
			<br><span style="color:#6a3705">Register as a <em>60+</em> Early Bird to participate.</span>
			<br>No cost... No obligation
			</h3><br>
			<form id="signup" action="/register" method="POST">
			<div><input class="field" name="email"
			  type="email" {{if .IsInput}}value={{.Email}}
			  {{else}}{{.PhEmail}}{{end}}
			  maxlength="128" required autofocus/>
				<div class="img-info">
				  <img src="static/img/base/info.png" alt="info" />
				  <p class="img-text">Email Address used to Sign Up or Log In. This will allow you to leave comments</p>
				</div>  
			</div>  
			<div><input class="field" name="username"
			  type="text" {{if .IsInput}}value={{.Username}}
			  {{else}}{{.PhUsername}}{{end}} 
			  maxlength="30" required autocomplete="off"/>
				<div class="img-info">
				  <img src="static/img/base/info.png" alt="info" />
				  <p class="img-text">Public screen name that others will use to identify you. Max length 30</p>
				</div>  
			</div>  
			<div><input class="field" name="fullname"
			  type="text" {{if .IsInput}}value={{.Fullname}}
			  {{else}}{{.PhFullname}}{{end}} 
			  maxlength="60" autocomplete="off"/>
				<div class="img-info">
				  <img src="static/img/base/info.png" alt="info" />
				  <p class="img-text">Used only in emails sent to you. If blank, Screen Name is used. Max length 60</p>
				</div>  
			</div>  
			<div><input class="field" name="pw1"
			  type="password" placeholder= "Required: Enter a password for this site" maxlength="60" required/>
				<div class="img-info">
				  <img src="static/img/base/info.png" alt="info" />
				  <p class="img-text">Your password is used to authorize access.<br>8 to 60 letters, numbers, symbols</p>
				</div>  
			</div>
			<div><input class="field" name="pw2"
			  type="password" placeholder= "Required: Confirm password" minlength="8" maxlength="60" required/>
				<div class="img-info">
				  <img src="static/img/base/info.png" alt="info" />
				  <p class="img-text">Re-enter your desired password</p>
				</div>
			</div>
			<div>
			  <input type="checkbox" name="autolog" 
			  {{if .AutoLog}}checked{{end}}/><b> Remember me!</b>
				<div class="img-info">
				  <img src="static/img/base/info.png" alt="info" />
				  <p class="img-text">Check this box to enable automatic login. Not recommended unless this is your personal device.<br>(i.e. not shared).</p>
				</div>
			</div>
			{{.xsrftoken}}
			<div><a href="/register?clear=y">&nbsp;Clear Form&nbsp;</a>&nbsp;&nbsp;&nbsp;
			<input class="button" type="submit"
			  value="Register"/>&nbsp;&nbsp;&nbsp;
			  <a href="/register?cancel=y">&nbsp;Cancel&nbsp;</a>
			</div>  
			</form>  
			<p class="info-text"><span style="color:#f0f8ff">
				We do <b>not</b> share email addresses.</span>
			</p>
		</div>
	</div>
</div>