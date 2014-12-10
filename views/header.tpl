<!DOCTYPE html>

<html>
	<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	<meta name="author" content="emadera52">
    <meta name="description" content="Sixty Plus Adventures helps retired travellers find destinations and activities tailored to their lifestyle." />
    <meta name="keywords" content="60Plus, retired, senior, boomer, vacation, travel, activities, lifestyle, 60+" />

    <title>{{.Website}}</title>
	<link rel="shortcut icon" href="/static/img/base/favicon.png">

	<!-- Stylesheets -->
	<link rel="stylesheet" href="/static/css/base/bootstrap.min.css">
	<link rel="stylesheet" href="/static/css/base/bootstrap-theme.min.css">
	<link rel="stylesheet" href="/static/css/base/base.css">
	<link rel="stylesheet" href="/static/css/sixty.css"/>

	<!-- Scripts -->
	<script type="text/javascript" src="/static/js/jquery.min.js"></script>
	<script type="text/javascript" src="/static/js/jquery.extend.js"></script>
	<script type="text/javascript" src="/static/js/bootstrap.js"></script>
	<script type="text/javascript" src="/static/js/jquery.clingify.min.js"></script>
	<!--[if lt IE 9]>
		<script type="text/javascript" src="/static/js/html5shiv.js"></script>
		<script type="text/javascript" src="/static/js/respond.min.js"></script>
	<![endif]-->
	</head>
	<body id="front">
		<noscript>Please enable Javascript in your browser</noscript> 
		<div id="wrapper">
			<nav class="navbar navbar-default navbar-fixed-top" >
			  <div class="container" data-toggle="clingify">
				<div class="row">
					<div class="navbar-header">
					  <a type="button" class="navbar-toggle collapsed" 
					  data-toggle="collapse" data-target="#navbar-collapse">
						<span class="sr-only">Toggle navigation</span>
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
					  </a>
					  <a class="hidden-xs hidden-sm logo" href="/">
						<img style="height:32px;width:60px;" 
						src="static/img/base/nav-logo.png">
					  </a>
					  <div class="visible-xs text-center">
						<a class="navbar-brand" href="/">
							<span color="white"><em>60+ !!!</em></span>
						</a>
					  </div>  
					</div>
					<div class="collapse navbar-collapse" role="navigation" 
						id="navbar-collapse">
						<ul class="nav navbar-nav">
							<li {{if .IsHome}}class="active"{{end}}><a href="/">Home</a></li>
							<li {{if .IsAdventurer}}class="active"{{end}}><a href="/comments?adventurer=y">Adventurer</a></li>
							<li {{if .IsAdvocate}}class="active"{{end}}><a href="/comments?advocate=y">Advocate</a></li>
							<li {{if .IsProvider}}class="active"{{end}}><a href="/comments?provider=y">Provider</a></li>
							<li {{if .IsSupporter}}class="active"{{end}}><a href="/comments?supporter=y">Supporter</a></li>
							<li {{if .IsFAQ}}class="active"{{end}}><a href="/comments?faq=y">FAQ</a></li>
							<li {{if .IsBlog}}class="active"{{end}}><a href="/blog?dest=blog">Blog</a></li>
							<li {{if .IsAbout}}class="active"{{end}}><a href="/comments?about=y">About Us</a></li>
						</ul>
					</div>
				</div>
			  </div><!-- /.container -->
			</nav>
			<header class="hero-unit">
				<h1><em>60+ </em>Adventures!</h1>
				<h2>{{if .TagTxt}}{{.TagTxt}}{{else}}Your guide to retirement vacation fun{{end}}</h2>
			</header>
			<!-- start page content -->
