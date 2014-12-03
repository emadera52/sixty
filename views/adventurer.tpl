	<div class="container" >
		<div class="row home-box">
			<div class="col-md-5">
				<p class="title">What We're Up To</p>
				<p class="content">We'll use this page to keep you up to date on progress. As an <em>Early Bird</em>, you'll have access to "page previews" as features are added. New pages and page updates will include ideas and suggestions based on comments submitted on this page.<br><br>We are currently working to complete pages for all of the "Under Construction" links and menu options. Next we'll go to work on the site's main home page. Eventually it will replace the one you see when you click "Home".<br><br>Check out the column to the right to see areas where your input is especially important. Feel free to include other ideas, suggestions and constructive criticism. Thanks for your input.
			    </p>
			</div>	
			<div class="col-md-5">
				<p class="title">How You Can Participate</p>
				<p class="content">We'd like your input as work begins on the <em>60+</em> Adventures home page. It will be different from the current home page.<br><br>Differences will include content specifically directed toward Adventurers. Menu options will change and there will be a column showing popular destinations as illustrated on the right. How would you like to see destinations organized? What filter options would you like to see, if any?<br><br>When a destination is selected, a new page will open. It will be the main page for that destination. Creating an example is high on our priority list. As a retired traveller, what would you like to see on a destination's main page?</p>
			</div>
			{{template "destinations.tpl" .}}
			{{if .InSession}} {{template "comment.tpl" .}} {{else}}
			{{template "reqlogin.tpl"}}
			{{end}}
		</div>
	</div> <!-- end container -->
	<div id="wrapper-push"></div>
</div> <!-- end wrapper started in Header -->
