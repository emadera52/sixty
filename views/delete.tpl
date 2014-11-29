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
				<h2>Delete <em>{{.Username}}</em> ?</h2>
				<h3>Are you certain you want to <span style="color:#6a3705">DELETE</span> 
				<br>your <em>60+</em> Adventures Profile and Comments?</h3>
				<br><h3>Click <span style="color:#6a3705"> Yes </span> to CONFIRM DELETE or
				<br>click <span style="color:#6a3705"> Cancel </span> to keep your Early Bird status.</h3><br>
				<a href="/delete?delete=y">&nbsp;Yes&nbsp;</a>
				&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	
				<a href="/delete?cancel=y">&nbsp;Cancel&nbsp;</a>
			</div>
		</div>	
	</div>
