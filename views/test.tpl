<!DOCTYPE html>
<html>
<head>
    <title>Confirm Dialog Box</title>
    <style>
        #white-background{
            display: none;
            width: 100%;
            height: 100%;
            position: fixed;
            top: 0;
            left: 0;
            background-color: #fefefe;
            opacity: 0.7;
            z-index: 9999;
        }
        #dlgbox{
            /*initially dialog box is hidden*/
            display: none;
            position: fixed;
            width: 480px;
            z-index: 9999;
            border-radius: 10px;
            background-color: #7c7d7e;
        }
        #dlg-header{
            background-color: #6d84b4;
            color: white;
            font-size: 20px;
            padding: 10px;
            margin: 10px 10px 0 10px;
        }
        #dlg-body{
            background-color: white;
            color: black;
            font-size: 14px;
            padding: 10px;
            margin: 0 10px 0 10px;
        }
        #dlg-footer{
            background-color: #f2f2f2;
            text-align: right;
            padding: 10px;
            margin: 0 10px 10px 10px;
        }
        #dlg-footer button{
            background-color: #6d84b4;
            color: white;
            padding: 5px;
            border: 0;
        }
    </style>
</head>
<body>
<!-- dialog box -->
<div id="white-background">
</div>
<div id="dlgbox">
    <div id="dlg-header">Please Confirm</div>
    <div id="dlg-body">Do you want to continue?</div>
    <div id="dlg-footer">
        <button onclick="dlgOK()">OK</button>
        <button onclick="dlgCancel()">Cancel</button>
    </div>
</div>

<!-- rest of the page -->
<h1>Dialog Box Demo</h1>
<p>This is a dialog box example.</p>
<p>Feel free to experiment with the code.</p>
<p>Click the button below to see the dialog box.</p>
<button onclick="showDialog()">Click Me!</button>

<!-- script of dialog -->
<script>
    function dlgCancel(){
        dlgHide();
        document.getElementsByTagName("H1")[0].innerHTML = "You clicked Cancel.";
    }
    function dlgOK(){
        dlgHide();
        document.getElementsByTagName("H3")[0].innerHTML = "You clicked OK.";
	}
    function dlgHide(){
        var whitebg = document.getElementById("white-background");
        var dlg = document.getElementById("dlgbox");
        whitebg.style.display = "none";
        dlg.style.display = "none";
    }
    function showDialog(){
        var whitebg = document.getElementById("white-background");
        var dlg = document.getElementById("dlgbox");
        whitebg.style.display = "block";
        dlg.style.display = "block";
        var winWidth = window.innerWidth;
        dlg.style.left = (winWidth/2) - 480/2 + "px";
        dlg.style.top = "150px";
	}
</script>
</body>
</html>