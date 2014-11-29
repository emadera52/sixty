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
				<h3>Please <em>Sign In</em> to enable your  
				<br>
				<em>60+</em> Early Bird Comment pivileges.
				</h3><br>
				<form id="signup" action="/login" method="POST">
				<div>
					<input class="field" name="username" type="text"
					placeholder="Screen name (user name)"
					Maxlength="30" required autofocus/>
					<div class="img-info">
						<img src="static/img/base/info.png" alt="info" />
						<p class="img-text">Enter the Screen Name you used when you registered</p>
					</div>  
				</div>  
				<div>
					<input class="field" name="password" type="password" placeholder= "Enter your 60+ Adventures password" 
					maxlength="60" required/>
					<div class="img-info">
						<img src="static/img/base/info.png" alt="info" />
						<p class="img-text">Required. This must match the password provided when you registered</p>
					</div>
				</div>	
				{{.xsrftoken}}
				<div>
				<input class="button" type="submit"
				  value="Sign In"/>&nbsp;&nbsp;&nbsp;
				  <a href="/login?cancel=y">&nbsp;Cancel&nbsp;</a>
				</div>  
				</form>  
			</div>
		</div>	
	</div>
