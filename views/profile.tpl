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
			<h2>Profile for: <em>{{.Username}}</em></h2>
			<form id="signup" action="/profile" method="POST">
				<div><input class="field" name="email"
					type="email" value={{.Email}}
					maxlength="128"/>
					<div class="img-info">
					  <img src="static/img/base/info.png" alt="info" />
					  <p class="img-text">Changing your email address will reset your status to unverified.</p>
					</div>  
				</div>  
				<div><input class="field" name="username"
				  type="text" value={{.Username}}
				  maxlength="30"/>
					<div class="img-info">
					  <img src="static/img/base/info.png" alt="info" />
					  <p class="img-text">Change the Screen Name that others use to identify you. Max length 30.</p>
					</div>  
				</div>  
				<div><input class="field" name="fullname"
				  type="text" value={{.Fullname}}
				  maxlength="60"/>
					<div class="img-info">
					  <img src="static/img/base/info.png" alt="info" />
					  <p class="img-text">Used only in emails sent to you. If blank your Screen Name will be used.</p>
					</div>  
				</div>
				<hr>
				<p>To CHANGE your password enter the NEW password in both boxes</p>
				<div><input class="field" name="pw1"
				  type="password" placeholder= "Use only to CHANGE your password" maxlength="60"/>
					<div class="img-info">
					  <img src="static/img/base/info.png" alt="info" />
					  <p class="img-text">Optional: Enter a NEW password. 8 to 60 letters, numbers and symbols.</p>
					</div>  
				</div>
				<div><input class="field" name="pw2"
				  type="password" placeholder= "Confirm NEW password" maxlength="60"/>
					<div class="img-info">
					  <img src="static/img/base/info.png" alt="info" />
					  <p class="img-text">Re-enter the desired NEW password</p>
					</div>
				</div>
				<div>
				  <input type="checkbox" name="autolog"
					{{if .AutoLog}}checked{{end}}/><b> Remember me!</b>
					<div class="img-info">
					  <img src="static/img/base/info.png" alt="info" />
					  <p class="img-text">You may change you auto-login status by checking or un-checking the box.</p>
					</div>
				</div>  
				{{ .xsrftoken }}
				<a href="/profile?delete=y">&nbsp;Delete&nbsp;</a>	
				<div class="img-info">
					<img src="static/img/base/info.png" alt="info" />
					<p class="img-text">Caution: This will remove all of your information from 60+ Adventures.</p>
				</div>&nbsp;&nbsp;&nbsp;
				<input class="button" type="submit"
				value="Update Profile"/>&nbsp;&nbsp;&nbsp;
				<a href="/profile?cancel=y">&nbsp;Cancel&nbsp;</a>
			</form>
		</div>
	</div>
</div>
