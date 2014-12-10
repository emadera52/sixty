	<div class="container" >
		<div class="sign-up">
			<a href="#auth">&nbsp;Comment&nbsp;</a><br><br>
		</div>
		<div class="row home-box">
			<div class="col-md-4">
				<p class="title">Questions</p>
				<p class="question">What are <em>60+</em> Early Bird privileges?
			    </p>
			</div>	
			<div class="col-md-6">
				<p class="title">Answers</p>
				<p class="content">Being an <b>Early Bird</b> gives you the opportunity to let us know what features and options you'd like to see in the final product. We'll also offer previews as new destinations are added. You will qualify for special offers and discounts at new destinations. There will be other privileges as well.<br><em>Leave a comment telling us what you'd like to see!</em></p>
			</div>
			{{template "destinations.tpl" .}}
			<div class="col-md-4">
				<p class="question">What happens to the information I give you when I register?
			    </p>
			</div>	
			<div class="col-md-6">
				<p class="content">It is stored in a database located on a secure server. We take extra precautions with your <b>email address</b> and <b>password</b>. The email address is encrypted using a secret key. Every user's key is different. We use the key to retrieve an email address when it is needed. Your password is processed in a way that no one, not even us, can decipher it.<br><em>We will not share your information with anyone else.</em></p>
			</div>
			<div class="col-md-4">
				<p class="question">Why must my web browser have "Cookies" enabled before I can Sign In?
			    </p>
			</div>	
			<div class="col-md-6">
				<p class="content">We use cookies as part of our user authentication procedures. They help us protect our legitimate users from hackers. On the other hand, cookies themselves cannot carry viruses, nor can they install malware on you computer.<br><br>If you have privacy concerns you may want to block <em>third party tracking cookes</em>. Doing so won't limit your access on <em>60+</em> Adventures.</p>
			</div>
			<div class="col-md-4">
				<p class="question"><em>Do you have a question that's not answered here?</em>
			    </p>
			</div>	
			<div class="col-md-6">
				<p class="content"><em>Use the comment feature below to ask it.</em> We'll get back to you by email with a response</p>
			</div>
			<div class="col-md-10">
				<a id="auth"></a>
				<div class="sign-up"><br>Return to &nbsp;<a href="/comments?faq=y">&nbsp;Top&nbsp;</a></div>
			</div>
			{{if .InSession}} {{template "comment.tpl" .}} {{else}}
			{{template "reqlogin.tpl"}}
			{{end}}
		</div>
	</div> <!-- end container -->
	<div id="wrapper-push"></div>
</div> <!-- end wrapper started in Header -->
