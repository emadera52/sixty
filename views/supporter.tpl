	<div class="container" >
		<div class="sign-up">
			<a href="#auth">&nbsp;Comment&nbsp;</a><br><br>
		</div>
		<div class="row home-box">
			<div class="col-md-5">
				<p class="title">What We're Up To</p>
				<p class="content">Check this page to keep up to date on features and activity of interest to corporate Supporters. As an <em>Early Bird</em>, you'll have access to "page previews" as features of interest are added. Special pre-launch promotions are now in the design phase. This is a great opportunity to influence decisions that will impact you and other <em>60+</em> Supporters.<br><br>We are working to complete pages for all of the "Under Construction" links and menu options. Next we'll go to work on revising the site's main home page.<br><br>Content on the new home page will be directed toward Adventurers. Menu options will change. A new column will be added as illustrated on the right.
			    </p>
			</div>	
			<div class="col-md-5">
				<p class="title">How You Can Participate</p>
				<p class="content">When a destination is selected, the main page for that destination will open. As a corporate Supporter or an agency working on their behalf, you may find the new home page and/or selected destination pages of interest.<br><br>Creating an example destination page is high on our priority list. As a <em>60+</em> corporate Supporter what would you like to see on a destination's main page?<br><br>A custom web page for each active Supporter is under consideration. Promotional links on selected pages would open your custom page. It could also be cross linked to your existing web site(s). What are your thoughts on this idea?</p>
			</div>
			{{template "destinations.tpl" .}}
			<div class="col-md-10">
				<a id="auth"></a>
				<div class="sign-up"><br>Return to &nbsp;<a href="/comments?supporter=y">&nbsp;Top&nbsp;</a></div>
			</div>
			{{if .InSession}} {{template "comment.tpl" .}} {{else}}
			{{template "reqlogin.tpl"}}
			{{end}}
		</div>
	</div> <!-- end container -->
	<div id="wrapper-push"></div>
</div> <!-- end wrapper started in Header -->
