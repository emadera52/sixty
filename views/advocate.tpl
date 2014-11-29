	<div class="container" >
		<div class="row home-box">
			<div class="col-md-5">
				<p class="title">What We're Up To</p>
				<p class="content">We'll use this page to keep you up to date on progress. As an <em>Early Bird</em>, you'll have access to "page previews" as features for Advocates are added. Programs and support for Advocates are now in the design phase. This is a great opportunity to influence decisions that will impact you and other <em>60+</em> Advocates.<br><br>The <em>60+</em> Adventures home page will be different from the current home page. Differences will include content specifically directed toward Adventurers. Menu options will change and there will be a column showing popular destinations as illustrated on the right.
			    </p>
			</div>	
			<div class="col-md-5">
				<p class="title">How You Can Participate</p>
				<p class="content">The ideal Advocate is active in promoting the tourist destination they call home. Experience helping local businesses develop promotional campaigns is a plus.<br><br>Your input on changes to the home page are appreciated. Also, when a destination is selected, a new page (the main page for that destination) will open. Creating an example is high on our priority list. As an Advocate, what would you like to see on a destination's main page?<br><br>What do you think of including a custom web page for each advertiser? Links on the destination's main page would open that Provider's custom page.</p>
			</div>
			{{template "destinations.tpl" .}}
			{{if .InSession}} {{template "comment.tpl" .}} {{else}}
			{{template "reqlogin.tpl"}}
			{{end}}
		</div>
	</div> <!-- end container -->
</div> <!-- end wrapper started in Header -->
