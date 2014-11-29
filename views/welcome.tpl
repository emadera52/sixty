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
			<h3><b>Welcome <span style="color:#6a3705">{{if .Interest}}&nbsp;{{.Interest}}&nbsp;{{end}}<em>{{.Username}}</em></span>
			<br>Use the Comment area at the bottom of a page
			<br>to submit questions, suggestions and ideas.</b>
			</h3><br>
			<form id="signup" action="/profile" method="GET">
				{{.xsrftoken}}
				<b>You may also&nbsp;&nbsp;&nbsp;</b>
				<input class="button" type="submit"
				  value="View/Update Your Profile"/><b>&nbsp;&nbsp;&nbsp;Or&nbsp;&nbsp;</b>
			  <a href="/logout">&nbsp;Sign Out&nbsp;</a>
			</form>
		</div>	
	</div>	
</div>
