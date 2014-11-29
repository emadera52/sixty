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
			<form id="signup" action="/login" method="GET">
				{{.xsrftoken}}
				<input class="button" type="submit" value="Sign In"/>
			</form>
			<h3><b>If you have previously registered,
			<br><em>Sign In</em> to enable your comment privileges.</b></h3><br>
			<form id="signup" action="/register" method="GET">
				{{.xsrftoken}}
				<div><input class="button" type="submit"
				  value="Register"/>
				</div>  
			</form>
			<h3>Or <em>Register</em> to get <em>60+</em> Early Bird privileges.</h3>
			<br><h3>Or continue as an unregistered guest.
			</h3>
			<h3>Check out the overview below, then choose a
			<br>topic from the menu above to get more details.<br></h3>
			<br><b><span style="color:#f0f8ff">Thank you for your interest.</span></b>
			</p>
		</div>	
	</div>	
</div>
