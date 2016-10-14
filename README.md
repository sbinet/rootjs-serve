rootjs-serve
============

A simple web server for JSROOT.

## Installation

```sh
$> go get github.com/sbinet/rootjs-serve
```

## Example

```sh
$> cat example.html 
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
		<meta charset="UTF-8" http-equiv="X-UA-Compatible" content="IE=Edge; text/html">
		<title>Reading object from the ROOT file</title>
		<script type="text/javascript" src="https://root.cern.ch/js/latest/scripts/JSRootCore.js"></script>
		<script type='text/javascript'>
var filename = "hsimple.root";
JSROOT.OpenFile(filename, function(file) {
	file.ReadObject("hpx", function(obj) {
		JSROOT.draw("drawing", obj, "colz");
	});
});
		</script>
	</head>

	<body>
		<div id="drawing" style="width:800px; height:600px"></div>
	</body>
</html>

$> root.exe -l /usr/share/doc/root/tutorials/hsimple.C
$> ll hsimple.root
-rw-r--r-- 1 binet binet 405K Oct 14 16:19 hsimple.root

$> rootjs-serve -dir=.
```

and in another terminal:

```sh
$> chromium localhost:5555/example.html
```
