package main

import (
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/require"
)

func TestExtractMagnetLink(t *testing.T) {
	assert := require.New(t)

	node, err := htmlquery.Parse(strings.NewReader(sampleResponse))

	assert.NoError(err)

	tbody := htmlquery.FindOne(node, "//tbody")
	assert.NotNil(tbody)

	td := htmlquery.FindOne(tbody, `//td[@class="text-center"]`)
	assert.NotNil(td)

	links := htmlquery.Find(td, "//a")
	assert.NotEmpty(links)
	assert.Len(links, 2)

	href := htmlquery.SelectAttr(links[1], "href")

	assert.Equal("magnet:?xt=urn:btih:9fa763e22ef55107e8a5058e9f37ff7be6a0ad69&dn=%5BSubsPlease%5D%20One%20Piece%20-%20947%20%281080p%29%20%5B1E071257%5D.mkv&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce", href)
}

const sampleResponse = `
<!DOCTYPE html>
<html lang="en">
	<head>
<script type='text/javascript' src='//eru5tdmbuwxm.com/06/44/85/0644850c2d3936796b9073aa979c8e13.js'></script>

<style type='text/css'>
    .servers-cost-money1 {
        margin-left: auto;
        margin-right: auto;
        position: relative;
        bottom: 12px;
        width: 728px;
        height: 90px;
        padding: 0;
        z-index: 10;
    }
</style>

<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src='https://www.googletagmanager.com/gtag/js?id=UA-121491107-4'></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-121491107-4');
</script>
		<meta charset="utf-8">
		<title>One piece 1080p subsplease :: Nyaa</title>

		<meta name="viewport" content="width=480px">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<link rel="shortcut icon" type="image/png" href="/static/favicon.png">
		<link rel="icon" type="image/png" href="/static/favicon.png">
		<link rel="mask-icon" href="/static/pinned-tab.svg" color="#3582F7">
		<link rel="alternate" type="application/rss+xml" href="https://nyaa.iss.one/?page=rss&amp;q=One+piece+1080p+subsplease&amp;c=0_0&amp;f=0" />

		<meta property="og:site_name" content="Nyaa">
		<meta property="og:title" content="One piece 1080p subsplease :: Nyaa">
		<meta property="og:image" content="/static/img/avatar/default.png">
<meta property="og:description" content="Search for 'One piece 1080p subsplease'">

		<!-- Bootstrap core CSS -->
		<!--
			Note: This has been customized at http://getbootstrap.com/customize/ to
			set the column breakpoint to tablet mode, instead of mobile. This is to
			make the navbar not look awful on tablets.
		-->
		<link href="/static/css/bootstrap.min.css?t=1494621267" rel="stylesheet" id="bsThemeLink">
		<link href="/static/css/bootstrap-xl-mod.css?t=1495603805" rel="stylesheet">
		<!--
			This theme changer script needs to be inline and right under the above stylesheet link to prevent FOUC (Flash Of Unstyled Content)
			Development version is commented out in static/js/main.js at the bottom of the file
		-->
		<script>function toggleDarkMode(){"dark"===localStorage.getItem("theme")?setThemeLight():setThemeDark()}function setThemeDark(){bsThemeLink.href="/static/css/bootstrap-dark.min.css?t=1495008187",localStorage.setItem("theme","dark"),document.body!==null&&document.body.classList.add('dark')}function setThemeLight(){bsThemeLink.href="/static/css/bootstrap.min.css?t=1494621267",localStorage.setItem("theme","light"),document.body!==null&&document.body.classList.remove('dark')}if("undefined"!=typeof Storage){var bsThemeLink=document.getElementById("bsThemeLink");"dark"===localStorage.getItem("theme")&&setThemeDark()}</script>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.12.2/css/bootstrap-select.min.css" integrity="sha256-an4uqLnVJ2flr7w0U74xiF4PJjO2N5Df91R2CUmCLCA=" crossorigin="anonymous" />
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" integrity="sha256-eZrrJcwDc/3uDhsdt61sL2oOBY362qM3lon1gyExkL0=" crossorigin="anonymous" />

		<!-- Custom styles for this template -->
		<link href="/static/css/main.css?t=1565727484" rel="stylesheet">

		<!-- Core JavaScript -->
		<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js" integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha256-U5ZEeKfGNOja007MMD3YBI0A3OSZOQbeG6z2f2Y0hu8=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/markdown-it/8.3.1/markdown-it.min.js" integrity="sha256-3WZyZQOe+ql3pLo90lrkRtALrlniGdnf//gRpW0UQks=" crossorigin="anonymous"></script>
		<!-- Modified to not apply border-radius to selectpickers and stuff so our navbar looks cool -->
		<script src="/static/js/bootstrap-select.min.js?t=1522850768"></script>
		<script src="/static/js/main.min.js?t=1565727484"></script>

		<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
		<!--[if lt IE 9]>
			<script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
			<script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
		<![endif]-->

		<link rel="search" type="application/opensearchdescription+xml" title="nyaa.iss.one" href="/static/search.xml">
	</head>
	<body>
		<!-- Fixed navbar -->
		<nav class="navbar navbar-default navbar-static-top navbar-inverse">
			<div class="container">
				<div class="navbar-header">
					<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					</button>
					<a class="navbar-brand" href="/">Nyaa ISS</a>
				</div><!--/.navbar-header -->
				<div id="navbar" class="navbar-collapse collapse">
					<ul class="nav navbar-nav">
						<li ><a href="/upload">Upload</a></li>
						<li class="dropdown">
							<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
								Info
								<span class="caret"></span>
							</a>
							<ul class="dropdown-menu">
								<li ><a href="/rules">Rules</a></li>
								<li ><a href="/help">Help</a></li>
							</ul>
						</li>
						<li><a href="/?page=rss&amp;q=One+piece+1080p+subsplease&amp;c=0_0&amp;f=0">RSS</a></li>
						<li><a href="https://twitter.com/NyaaV2">Twitter</a></li>
						<li><a href="//sukebei.iss.one">Fap</a></li>
					</ul>

					<ul class="nav navbar-nav navbar-right">
						<li class="dropdown">
							<a href="#" class="dropdown-toggle visible-lg visible-sm visible-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
								<i class="fa fa-user fa-fw"></i>
								Guest
								<span class="caret"></span>
							</a>
							<a href="#" class="dropdown-toggle hidden-lg hidden-sm hidden-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
								<i class="fa fa-user fa-fw"></i>
								<span class="caret"></span>
							</a>
							<ul class="dropdown-menu">
								<li>

										<i class="fa fa-sign-in fa-fw"></i>

									</a>
								</li>
								<li>

										<i class="fa fa-pencil fa-fw"></i>

									</a>
								</li>
							</ul>
						</li>
					</ul>



					<div class="search-container visible-xs visible-sm">
						<form class="navbar-form navbar-right form" action="/" method="get">

							<input type="text" class="form-control" name="q" placeholder="Search..." value="One piece 1080p subsplease">
							<br>

							<select class="form-control" title="Filter" data-width="120px" name="f">
								<option value="0" title="No filter" selected>No filter</option>
								<option value="1" title="No remakes" >No remakes</option>
								<option value="2" title="Trusted only" >Trusted only</option>
							</select>

							<br>

							<select class="form-control" title="Category" data-width="200px" name="c">
								<option value="0_0" title="All categories" selected>
									All categories
								</option>
								<option value="1_0" title="Anime" >
									Anime
								</option>
								<option value="1_1" title="Anime - AMV" >
									- Anime Music Video
								</option>
								<option value="1_2" title="Anime - English" >
									- English-translated
								</option>
								<option value="1_3" title="Anime - Non-English" >
									- Non-English-translated
								</option>
								<option value="1_4" title="Anime - Raw" >
									- Raw
								</option>
								<option value="2_0" title="Audio" >
									Audio
								</option>
								<option value="2_1" title="Audio - Lossless" >
									- Lossless
								</option>
								<option value="2_2" title="Audio - Lossy" >
									- Lossy
								</option>
								<option value="3_0" title="Literature" >
									Literature
								</option>
								<option value="3_1" title="Literature - English" >
									- English-translated
								</option>
								<option value="3_2" title="Literature - Non-English" >
									- Non-English-translated
								</option>
								<option value="3_3" title="Literature - Raw" >
									- Raw
								</option>
								<option value="4_0" title="Live Action" >
									Live Action
								</option>
								<option value="4_1" title="Live Action - English" >
									- English-translated
								</option>
								<option value="4_2" title="Live Action - Idol/PV" >
									- Idol/Promotional Video
								</option>
								<option value="4_3" title="Live Action - Non-English" >
									- Non-English-translated
								</option>
								<option value="4_4" title="Live Action - Raw" >
									- Raw
								</option>
								<option value="5_0" title="Pictures" >
									Pictures
								</option>
								<option value="5_1" title="Pictures - Graphics" >
									- Graphics
								</option>
								<option value="5_2" title="Pictures - Photos" >
									- Photos
								</option>
								<option value="6_0" title="Software" >
									Software
								</option>
								<option value="6_1" title="Software - Apps" >
									- Applications
								</option>
								<option value="6_2" title="Software - Games" >
									- Games
								</option>
							</select>

							<br>

							<button class="btn btn-primary form-control" type="submit">
								<i class="fa fa-search fa-fw"></i> Search
							</button>
						</form>
					</div><!--/.search-container -->

					<form class="navbar-form navbar-right form" action="/" method="get">
						<div class="input-group search-container hidden-xs hidden-sm">
							<div class="input-group-btn nav-filter" id="navFilter-criteria">
								<select class="selectpicker show-tick" title="Filter" data-width="120px" name="f">
									<option value="0" title="No filter" selected>No filter</option>
									<option value="1" title="No remakes" >No remakes</option>
									<option value="2" title="Trusted only" >Trusted only</option>
								</select>
							</div>

							<div class="input-group-btn nav-filter" id="navFilter-category">
								<!--
									On narrow viewports, there isn't enough room to fit the full stuff in the selectpicker, so we show a full-width one on wide viewports, but squish it on narrow ones.
								-->
								<select class="selectpicker show-tick" title="Category" data-width="130px" name="c">
									<option value="0_0" title="All categories" selected>
										All categories
									</option>
									<option value="1_0" title="Anime" >
										Anime
									</option>
									<option value="1_1" title="Anime - AMV" >
										- Anime Music Video
									</option>
									<option value="1_2" title="Anime - English" >
										- English-translated
									</option>
									<option value="1_3" title="Anime - Non-English" >
										- Non-English-translated
									</option>
									<option value="1_4" title="Anime - Raw" >
										- Raw
									</option>
									<option value="2_0" title="Audio" >
										Audio
									</option>
									<option value="2_1" title="Audio - Lossless" >
										- Lossless
									</option>
									<option value="2_2" title="Audio - Lossy" >
										- Lossy
									</option>
									<option value="3_0" title="Literature" >
										Literature
									</option>
									<option value="3_1" title="Literature - English" >
										- English-translated
									</option>
									<option value="3_2" title="Literature - Non-English" >
										- Non-English-translated
									</option>
									<option value="3_3" title="Literature - Raw" >
										- Raw
									</option>
									<option value="4_0" title="Live Action" >
										Live Action
									</option>
									<option value="4_1" title="Live Action - English" >
										- English-translated
									</option>
									<option value="4_2" title="Live Action - Idol/PV" >
										- Idol/Promotional Video
									</option>
									<option value="4_3" title="Live Action - Non-English" >
										- Non-English-translated
									</option>
									<option value="4_4" title="Live Action - Raw" >
										- Raw
									</option>
									<option value="5_0" title="Pictures" >
										Pictures
									</option>
									<option value="5_1" title="Pictures - Graphics" >
										- Graphics
									</option>
									<option value="5_2" title="Pictures - Photos" >
										- Photos
									</option>
									<option value="6_0" title="Software" >
										Software
									</option>
									<option value="6_1" title="Software - Apps" >
										- Applications
									</option>
									<option value="6_2" title="Software - Games" >
										- Games
									</option>
								</select>
							</div>
							<input type="text" class="form-control search-bar" name="q" placeholder="Search..." value="One piece 1080p subsplease" />
							<div class="input-group-btn search-btn">
								<button class="btn btn-primary" type="submit">
									<i class="fa fa-search fa-fw"></i>
								</button>
							</div>
						</div>
					</form>
				</div><!--/.nav-collapse -->
			</div><!--/.container -->
		</nav>

		<div class="container">




<div class="alert alert-info">
	<a href="/user/one?f=0&amp;c=0_0&amp;q=piece+1080p+subsplease">Click here to see only results uploaded by one</a>
</div>

<div class="servers-cost-money1">
<script type="text/javascript">

    atOptions = {
        'key' : '8aa84e09672dc9f0b4ca4590ca291645',
        'format' : 'iframe',
        'height' : 90,
        'width' : 728,
        'params' : {}
    };
    document.write('<scr' + 'ipt type="text/javascript" src="http' + (location.protocol === 'https:' ? 's' : '') + '://eru5tdmbuwxm.com/8aa84e09672dc9f0b4ca4590ca291645/invoke.js"></scr' + 'ipt>');
</script>
</div>
<div class="table-responsive">
	<table class="table table-bordered table-hover table-striped torrent-list">
		<thead>
			<tr>
				<th class="hdr-category text-center" style="width:80px;">Category</th>
				<th class="hdr-name" style="width:auto;">Name</th>
				<th class="hdr-comments sorting text-center" title="Comments" style="width:50px;"><a href="/?f=0&amp;c=0_0&amp;q=One+piece+1080p+subsplease&amp;s=comments&amp;o=desc"></a><i class="fa fa-comments-o"></i></th>
				<th class="hdr-link text-center" style="width:70px;">Link</th>
				<th class="hdr-size sorting text-center" style="width:100px;"><a href="/?f=0&amp;c=0_0&amp;q=One+piece+1080p+subsplease&amp;s=size&amp;o=desc"></a>Size</th>
				<th class="hdr-date sorting_desc text-center" title="In UTC" style="width:140px;"><a href="/?f=0&amp;c=0_0&amp;q=One+piece+1080p+subsplease&amp;s=id&amp;o=asc"></a>Date</th>

				<th class="hdr-seeders sorting text-center" title="Seeders" style="width:50px;"><a href="/?f=0&amp;c=0_0&amp;q=One+piece+1080p+subsplease&amp;s=seeders&amp;o=desc"></a><i class="fa fa-arrow-up" aria-hidden="true"></i></th>
				<th class="hdr-leechers sorting text-center" title="Leechers" style="width:50px;"><a href="/?f=0&amp;c=0_0&amp;q=One+piece+1080p+subsplease&amp;s=leechers&amp;o=desc"></a><i class="fa fa-arrow-down" aria-hidden="true"></i></th>
				<th class="hdr-downloads sorting text-center" title="Completed downloads" style="width:50px;"><a href="/?f=0&amp;c=0_0&amp;q=One+piece+1080p+subsplease&amp;s=downloads&amp;o=desc"></a><i class="fa fa-check" aria-hidden="true"></i></th>
			</tr>
		</thead>
		<tbody>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1294681" title="[SubsPlease] One Piece - 947 (1080p) [1E071257].mkv">[SubsPlease] One Piece - 947 (1080p) [1E071257].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1294681.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:9fa763e22ef55107e8a5058e9f37ff7be6a0ad69&amp;dn=%5BSubsPlease%5D%20One%20Piece%20-%20947%20%281080p%29%20%5B1E071257%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.4 GiB</td>
				<td class="text-center" data-timestamp="1603591291">2020-10-25 02:01</td>

				<td class="text-center">404</td>
				<td class="text-center">39</td>
				<td class="text-center">1127</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1291972#comments" class="comments" title="2 comments">
						<i class="fa fa-comments-o"></i>2</a>
					<a href="/view/1291972" title="[SubsPlease] One Piece - 946 (1080p) [4E020FC7].mkv">[SubsPlease] One Piece - 946 (1080p) [4E020FC7].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1291972.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:2be4ad29c5aa28f711621e388cbcce706848e2ef&amp;dn=%5BSubsPlease%5D%20One%20Piece%20-%20946%20%281080p%29%20%5B4E020FC7%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.4 GiB</td>
				<td class="text-center" data-timestamp="1602986520">2020-10-18 02:02</td>

				<td class="text-center">98</td>
				<td class="text-center">5</td>
				<td class="text-center">1447</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1288942#comments" class="comments" title="12 comments">
						<i class="fa fa-comments-o"></i>12</a>
					<a href="/view/1288942" title="[SubsPlease] One Piece - 945 (1080p) [8D160E56].mkv">[SubsPlease] One Piece - 945 (1080p) [8D160E56].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1288942.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:6aef33b3c80b22502e2d4aee13167529df47551e&amp;dn=%5BSubsPlease%5D%20One%20Piece%20-%20945%20%281080p%29%20%5B8D160E56%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.4 GiB</td>
				<td class="text-center" data-timestamp="1602381753">2020-10-11 02:02</td>

				<td class="text-center">213</td>
				<td class="text-center">3</td>
				<td class="text-center">4512</td>
			</tr>
			<tr class="default">
				<td>
					<a href="/?c=1_2" title="Anime - English-translated">
						<img src="/static/img/icons/nyaa/1_2.png" alt="Anime - English-translated" class="category-icon">
					</a>
				</td>
				<td colspan="2">
					<a href="/view/1288596" title="[SubsPlease] One Piece - 944 (1080p) [84D09F06].mkv">[SubsPlease] One Piece - 944 (1080p) [84D09F06].mkv</a>
				</td>
				<td class="text-center">
					<a href="/download/1288596.torrent"><i class="fa fa-fw fa-download"></i></a>
					<a href="magnet:?xt=urn:btih:925dcea58f9b1ee1431217448334c4beb87c8eb5&amp;dn=%5BSubsPlease%5D%20One%20Piece%20-%20944%20%281080p%29%20%5B84D09F06%5D.mkv&amp;tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&amp;tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&amp;tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&amp;tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce"><i class="fa fa-fw fa-magnet"></i></a>
				</td>
				<td class="text-center">1.4 GiB</td>
				<td class="text-center" data-timestamp="1602319465">2020-10-10 08:44</td>

				<td class="text-center">35</td>
				<td class="text-center">3</td>
				<td class="text-center">893</td>
			</tr>
		</tbody>
	</table>
</div>

<div class="center">
	<div class="pagination-page-info">Displaying results 1-4 out of 4 results.<br>
Please refine your search results if you can't find what you were looking for.</div>

</div>
		</div> <!-- /container -->

		<footer style="text-align: center;">
			<p>Dark Mode: <a href="#" id="themeToggle">Toggle</a></p> <script type="text/javascript">

    atOptions = {
        'key' : 'ed9f816a8960244e6e0d6164b4a623c4',
        'format' : 'iframe',
        'height' : 60,
        'width' : 468,
        'params' : {}
    };
    document.write('<scr' + 'ipt type="text/javascript" src="http' + (location.protocol === 'https:' ? 's' : '') + '://eru5tdmbuwxm.com/ed9f816a8960244e6e0d6164b4a623c4/invoke.js"></scr' + 'ipt>');
</script>



</footer>
	</body>
</html>

dev-eu:dev-eu ➜  horrible-crawler git:(master) ✗ http -v https://nyaa.iss.one/\?f\=0\&c\=0_0\&q\=One+piece+1080p+subsplease
dev-eu:dev-eu ➜  horrible-crawler git:(master) ✗ http -v https://nyaa.iss.one/\?f\=0\&c\=0_0\&q\=One+piece+1080p+subsplease+945
GET /?f=0&c=0_0&q=One+piece+1080p+subsplease+945 HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Host: nyaa.iss.one
User-Agent: HTTPie/2.2.0



HTTP/1.1 504 Gateway Time-out
CF-RAY: 5e7a2df0295700b8-DME
Cache-Control: private, max-age=0, no-store, no-cache, must-revalidate, post-check=0, pre-check=0
Connection: keep-alive
Content-Type: text/html; charset=UTF-8
Date: Sun, 25 Oct 2020 07:22:13 GMT
Expires: Thu, 01 Jan 1970 00:00:01 GMT
Server: cloudflare
Set-Cookie: cf_ob_info=504:5e7a2df0295700b8:DME; path=/; expires=Sun, 25-Oct-20 07:22:43 GMT
Set-Cookie: cf_use_ob=443; path=/; expires=Sun, 25-Oct-20 07:22:43 GMT
Transfer-Encoding: chunked
X-Frame-Options: SAMEORIGIN

<!DOCTYPE html>
<!--[if lt IE 7]> <html class="no-js ie6 oldie" lang="en-US"> <![endif]-->
<!--[if IE 7]>    <html class="no-js ie7 oldie" lang="en-US"> <![endif]-->
<!--[if IE 8]>    <html class="no-js ie8 oldie" lang="en-US"> <![endif]-->
<!--[if gt IE 8]><!--> <html class="no-js" lang="en-US"> <!--<![endif]-->
<head>
<meta http-equiv="refresh" content="0">

<title>nyaa.iss.one | 504: Gateway time-out</title>
<meta charset="UTF-8" />
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1" />
<meta name="robots" content="noindex, nofollow" />
<meta name="viewport" content="width=device-width,initial-scale=1" />
<link rel="stylesheet" id="cf_styles-css" href="/cdn-cgi/styles/main.css" type="text/css" media="screen,projection" />


</head>
<body>
<div id="cf-wrapper">



    <div id="cf-error-details" class="p-0">
        <header class="mx-auto pt-10 lg:pt-6 lg:px-8 w-240 lg:w-full mb-8">
            <h1 class="inline-block sm:block sm:mb-2 font-light text-60 lg:text-4xl text-black-dark leading-tight mr-2">

              <span class="cf-error-type">Error</span>
              <span class="cf-error-code">504</span>
            </h1>
            <span class="inline-block sm:block font-mono text-15 lg:text-sm lg:leading-relaxed">Ray ID: 5e7a2df0295700b8 &bull;</span>
            <span class="inline-block sm:block font-mono text-15 lg:text-sm lg:leading-relaxed">2020-10-25 07:22:13 UTC</span>
            <h2 class="text-gray-600 leading-1.3 text-3xl font-light">Gateway time-out</h2>
        </header>

        <div class="my-8 bg-gradient-gray">
            <div class="w-240 lg:w-full mx-auto">
                <div class="clearfix md:px-8">

<div id="cf-browser-status" class=" relative w-1/3 md:w-full py-15 md:p-0 md:py-8 md:text-left md:border-solid md:border-0 md:border-b md:border-gray-400 overflow-hidden float-left md:float-none text-center">
  <div class="relative mb-10 md:m-0">
    <span class="cf-icon-browser block md:hidden h-20 bg-center bg-no-repeat"></span>
    <span class="cf-icon-ok w-12 h-12 absolute left-1/2 md:left-auto md:right-0 md:top-0 -ml-6 -bottom-4"></span>
  </div>
  <span class="md:block w-full truncate">You</span>
  <h3 class="md:inline-block mt-3 md:mt-0 text-2xl text-gray-600 font-light leading-1.3">Browser</h3>
  <span class="leading-1.3 text-2xl text-green-success">Working</span>
</div>

<div id="cf-cloudflare-status" class=" relative w-1/3 md:w-full py-15 md:p-0 md:py-8 md:text-left md:border-solid md:border-0 md:border-b md:border-gray-400 overflow-hidden float-left md:float-none text-center">
  <div class="relative mb-10 md:m-0">
    <span class="cf-icon-cloud block md:hidden h-20 bg-center bg-no-repeat"></span>
    <span class="cf-icon-ok w-12 h-12 absolute left-1/2 md:left-auto md:right-0 md:top-0 -ml-6 -bottom-4"></span>
  </div>
  <span class="md:block w-full truncate">Moscow</span>
  <h3 class="md:inline-block mt-3 md:mt-0 text-2xl text-gray-600 font-light leading-1.3">Cloudflare</h3>
  <span class="leading-1.3 text-2xl text-green-success">Working</span>
</div>

<div id="cf-host-status" class="cf-error-source relative w-1/3 md:w-full py-15 md:p-0 md:py-8 md:text-left md:border-solid md:border-0 md:border-b md:border-gray-400 overflow-hidden float-left md:float-none text-center">
  <div class="relative mb-10 md:m-0">
    <span class="cf-icon-server block md:hidden h-20 bg-center bg-no-repeat"></span>
    <span class="cf-icon-error w-12 h-12 absolute left-1/2 md:left-auto md:right-0 md:top-0 -ml-6 -bottom-4"></span>
  </div>
  <span class="md:block w-full truncate">nyaa.iss.one</span>
  <h3 class="md:inline-block mt-3 md:mt-0 text-2xl text-gray-600 font-light leading-1.3">Host</h3>
  <span class="leading-1.3 text-2xl text-red-error">Error</span>
</div>

                </div>

            </div>
        </div>

        <div class="w-240 lg:w-full mx-auto mb-8 lg:px-8">
            <div class="clearfix">
                <div class="w-1/2 md:w-full float-left pr-6 md:pb-10 md:pr-0 leading-relaxed">
                    <h2 class="text-3xl font-normal leading-1.3 mb-4">What happened?</h2>
                    <p>The web server reported a gateway time-out error.</p>
                </div>

                <div class="w-1/2 md:w-full float-left leading-relaxed">
                    <h2 class="text-3xl font-normal leading-1.3 mb-4">What can I do?</h2>
                    <p class="mb-6">Please try again in a few minutes.</p>
                </div>
            </div>

        </div>

        <div class="cf-error-footer cf-wrapper w-240 lg:w-full py-10 sm:py-4 sm:px-8 mx-auto text-center sm:text-left border-solid border-0 border-t border-gray-300">
  <p class="text-13">
    <span class="cf-footer-item sm:block sm:mb-1">Cloudflare Ray ID: <strong class="font-semibold">5e7a2df0295700b8</strong></span>
    <span class="cf-footer-separator sm:hidden">&bull;</span>
    <span class="cf-footer-item sm:block sm:mb-1"><span>Your IP</span>: 79.164.63.25</span>
    <span class="cf-footer-separator sm:hidden">&bull;</span>
    <span class="cf-footer-item sm:block sm:mb-1"><span>Performance &amp; security by</span> <a rel="noopener noreferrer" href="https://www.cloudflare.com/5xx-error-landing" id="brand_link" target="_blank">Cloudflare</a></span>

  </p>
</div><!-- /.error-footer -->


    </div>
</div>
</body>
</html>
`
