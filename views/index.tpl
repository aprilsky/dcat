<html>
<head>
	<meta charset="utf-8">
	<title>Cobub Razor - 开源移动应用分析平台</title>
	<style>* { margin:0;padding:0; }</style>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="../css/layout.css" type="text/css" media="screen">
	<link rel="stylesheet" href="../css/helplayout.css" type="text/css" media="screen">
	<link rel="stylesheet" type="text/css" href="../../css/default/easyui.css">

	<script src="../js/jquery.select.js" type="text/javascript"></script>
	<script src="../js/hideshow.js" type="text/javascript"></script>
	<script src="../js/jquery-ui-1.8.min.js" type="text/javascript"></script>
	<script src="../js/jquery-ui-1.8.16.custom.min.js" type="text/javascript"></script>
	<script src="../js/jquery.blockUI.js" type="text/javascript"></script>
	<script src="../js/jquery.equalHeight.js" type="text/javascript"></script>
	<script src="../js/jquery.pagination.js" type="text/javascript"></script>
	<script src="../js/jquery.tablesorter.min.js" type="text/javascript"></script>

	<script src="../js/easydialog.js" type="text/javascript"></script>
	<script src="./js/easydialog.min.js" type="text/javascript"></script>
	<script src="../js/estimate.js" type="text/javascript"></script>
	<script src="../js/exporting.js" type="text/javascript"></script>

	<script src="../js/jquery-1.7.1.min.js" type="text/javascript"></script>
	<script src="../js/highcharts-more.js" type="text/javascript"></script>
	<script src="../js/highcharts.js" type="text/javascript"></script>
	<script src="../js/json2.js" type="text/javascript"></script>
	<script type="text/javascript" src="../js/jquery.easyui.min.js"></script>

	<script  type="text/javascript">

	function loadNewApp(){
		$('#main').load("/add",function(){

		});
	}

	function appInfo(appname,packagename,versions,platforms,market){
　　　　this.appname=appname;
　　　　this.packagename=packagename;
        this.versions=versions;
        this.platforms=platforms;
        this.market=market;
　　}

     function submit(){
        var appname = $('#appname').val();
        var packagename = $('#packagename').val();
        var version = $('#version').val();
        var platform = $('#platform').val();
        var market = $('#market').val();

        var app = new appInfo(appname,packagename,version,platform,market)
        var jsondata= JSON.stringify(app)
        $.ajax({
                url: '/app/add',
                type: 'POST',
                dataType: 'json',
                data: {"data":jsondata},
                complete: function(xhr, textStatus) {
                   
                },
                success: function(data, textStatus, xhr) {
                          alert(data.appname);
                },
                error: function(xhr, textStatus, errorThrown) {
     
                }
        });
    }
		 
	</script>
</head>
<body>
	<header id="header">
		<hgroup>
			<h1 class="site_title">
				<a href="http://demo.cobub.com/razor/">
					<img class="logo" src="http://demo.cobub.com/razor/assets/images/razorlogo.png" style="border:0">
					<span style="">Cobub Razor</span>
				</a>
			</h1>
			<h3 class="section_title">
				<a class="colorMediumBlue bold spanHover" href="http://demo.cobub.com/razor/index.php?/profile/modify">个人资料</a>
				|
				<a class="colorMediumBlue bold spanHover" href="http://demo.cobub.com/razor/index.php?/auth/change_password">修改密码</a>
				|
				<a class="colorMediumBlue bold spanHover" href="http://demo.cobub.com/razor/index.php?/auth/logout">退出</a>
			</h3>
		</hgroup>
	</header>
	<section id="secondary_bar">
		<div class="user">
			<p>
				{{.AccountName}} 
				(
				<a class="colorMediumBlue bold spanHover" href="http://demo.cobub.com/razor/index.php?/profile/modify">个人资料</a>
				)
			</p>
			<!-- <a class="logout_user" href="#" title="Logout">Logout</a>
		-->
	</div>
	<div class="breadcrumbs_container">
		<article id="breadcrumbs" class="breadcrumbs">
			<a href="javascript:loadIndex();">控制台</a>
			<div class="breadcrumb_divider"></div>
			<a class="current">我的应用</a>

		</article>
	</div>

</section>
<aside id="sidebar" class="column" style="height: 100%;">
	<h3>
		管理
		<a href="#" class="toggleLink">hide</a>
	</h3>
	<hr>
	<br>
	<ul class="toggle" style="">
		<li class="icn_my_application">
			<a class="colorMediumBlue bold spanHover" href="">应用列表</a>
		</li>
		<li class="icn_add_apps">
			<a class="colorMediumBlue bold spanHover"  href="javascript:loadNewApp();">新建应用</a>
		</li>
		<li class="icn_app_channel">
			<a class="colorMediumBlue bold spanHover" href="http://demo.cobub.com/razor/index.php?/manage/channel">渠道管理</a>
		</li>
	</ul>
	<br>
	<footer>
		<hr>

		<p>
			&nbsp;&nbsp;
			<a id="greencss"
					href="javascript:setcssstyle('greenlayout')">
				<img
					src="http://demo.cobub.com/razor/assets/images/greenbtn.png" style="border:0"/>
			</a>
			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
			<a id="layoutcss" href="javascript:setcssstyle('layout')">
				<img
					src="http://demo.cobub.com/razor/assets/images/graybtn.png" style="border:0"/>
			</a>
			&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
			<a id="bluecss" href="javascript:setcssstyle('bluelayout')">
				<img
					src="http://demo.cobub.com/razor/assets/images/bluebtn.png" style="border:0"/>
			</a>

		</p>

	</footer>

</aside>

<section id="main" class="column" style="height: 100%;">
	<div style="margin-left:1px;">
		<table  id="dg" class="easyui-datagrid"  style="margin-left:1px;height:100%;width:auto"
            data-options="singleSelect:true,url:'../datagrid/datagrid_data1.json'">
			<thead>
				<tr>
					<th data-options="field:'itemid',width:80">Item ID</th>
					<th data-options="field:'productid',width:100">Product</th>
					<th data-options="field:'listprice',width:80,align:'right'">List Price</th>
					<th data-options="field:'unitcost',width:80,align:'right'">Unit Cost</th>
					<th data-options="field:'attr1',width:250">Attribute</th>
					<th data-options="field:'status',width:60,align:'center'">Status</th>
				</tr>
			</thead>
		</table>
	</div>
</section>
</body>
</html>