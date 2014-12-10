	<div class="container" >
		<div class="sign-up">
			<a href="#auth">&nbsp;Comment&nbsp;</a><br><br>
		</div>
		<div class="row home-box">
			<div class="col-md-10">
				<p class="title">What is the goal for <em>60+</em>&nbsp;Adventures?</p>
				<p class="content">The goal of the site you see today is to guage public interest in a web site focused on providing accurate, timely and relevant information to retired travelers. Visitors like you can demonstrate an interest by registering and leaving comments. You can ask questions, make suggestions, leave constructive criticism. You can tell us that you like the idea or that you think it's a bad idea.<br>
				<br>If enough interest is shown, we'll go to work choosing a few <em>Destinations&nbsp;&nbsp;</em>to focus on. Each will have its own main page. There will be <em>photos, videos and text&nbsp;&nbsp;</em>highlighting its appeal to retired visitors. The key to selecting destinations will be finding an enthusiastic <em>Advocate&nbsp;&nbsp;</em>to work with us in promoting the destination and the <em>Providers&nbsp;&nbsp;</em>that will advertise their services on the site.<br>
				<br>In the right column will be a list of <em>attractions, activities, accommodations, cuisine&nbsp;&nbsp;</em>and anything else that includes features designed to appeal to <em>60+ Adventurers.&nbsp;&nbsp;Providers</em>&nbsp;&nbsp;will get special attention, including a link to their own web page hosted on the site.<br>
				<br>All of this is subject to change based on feedback we get from <em>Early Bird&nbsp;&nbsp;</em>visitors like you!&nbsp;&nbsp;Leave <em>Comments</em>&nbsp;&nbsp;letting us know what is important to you as an <em>Adventurer, Advocate, Provider or Supporter</em></p>
			</div>
			{{template "destinations.tpl" .}}
			<div class="col-md-10">
				<br><p class="title">Who is working on <em>60+</em> Adventures and what future plans exist?</p>
				<p class="content">This site began as the final assignment for an online college course taken by a 68 year old retired software developer living in Puerto Vallarta, Mexico. After a half dozen years of retirement he decided to learn something new. His previous career did not involve web pages or anything else related to the Interet. He started the class in early June of 2014. Work began on the final assignment (this website) in early September. It will have been submitted by December 20th.<br>
				<br>Others have shown an interest and may join the project after December 20th. The current plan is to move out of the <em>Early Bird</em> phase in early January of 2015. From that point promotional efforts will continue through March. At that point interest will be evaluated, followed by a decision regarding moving forward as described above.</p>
				<br>
				<a id="auth"></a>
				<div class="sign-up"><br>Return to &nbsp;<a href="/comments?about=y">&nbsp;Top&nbsp;</a></div>
			</div>
			{{if .InSession}} {{template "comment.tpl" .}} {{else}}
			{{template "reqlogin.tpl"}}
			{{end}}
		</div>
	</div> <!-- end container -->
	<div id="wrapper-push"></div>
</div> <!-- end wrapper started in Header -->
