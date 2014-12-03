	<div class="container" >
		<div class="row home-box">
			<div class="col-md-5">
				<p class="title">What We're Up To</p>
				<p class="content">We'll use this page to keep you up to date on progress. As an <em>Early Bird</em>, you'll have access to "page previews" as features for Providers are added. Support for Providers and special pre-launch promotions are now in the design phase. This is a great opportunity to influence decisions that will impact you and other <em>60+</em> Providers.<br><br>We are working to complete pages for all of the "Under Construction" links and menu options. Next we'll go to work on revising the site's main home page.<br><br>Differences will include content specifically directed toward Adventurers. Menu options will change and there will be a column showing popular destinations as illustrated on the right.
			    </p>
			</div>	
			<div class="col-md-5">
				<p class="title">How You Can Participate</p>
				<p class="content">As a Provider of activities, services and accommodations for retired travellers, your focus will be on the main page of the destination(s) you serve. When a destination is selected, the main page for that destination will open.<br><br>Creating an example is high on our priority list. As a <em>60+</em> Provider what would you like to see on a destination's main page?<br><br>A custom web page for each active Provider is under consideration. A promotional link on the destination's main page would open your custom page. It could also be cross linked to your existing web site, if any. What are your thoughts on this idea?</p>
			</div>
			{{template "destinations.tpl" .}}
			{{if .InSession}} {{template "comment.tpl" .}} {{else}}
			{{template "reqlogin.tpl"}}
			{{end}}
		</div>
	</div> <!-- end container -->
	<div id="wrapper-push"></div>
</div> <!-- end wrapper started in Header -->
