// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package short

const templateHTML = `
<!doctype HTML>
<html lang="en">
<title>golang.org URL shortener</title>
<style>
* {
	box-sizing: border-box;
}
body {
	font-family: system-ui, sans-serif;
}
input {
	border: 1px solid #ccc;
}
input[type=text],
input[type=url] {
	width: 400px;
}
input, td, th {
	color: #333;
}
input, td {
	font-size: 14pt;
}
th {
	font-size: 16pt;
	text-align: left;
	padding-top: 10px;
}
.autoselect {
	border: none;
}
.error {
	color: #900;
}
table {
	margin-left: auto;
	margin-right: auto;
}
</style>
<table>

{{with .Error}}
	<tr>
		<th colspan="3">Error</th>
	</tr>
	<tr>
		<td class="error" colspan="3">{{.}}</td>
	</tr>
{{end}}

<tr>
	<th>Key</th>
	<th>Target</th>
	<th></th>
</tr>

<form method="POST">
<tr>
	<td><input type="text" name="key"{{with .New}} value="{{.Key}}"{{end}} required></td>
	<td><input type="url" name="target"{{with .New}} value="{{.Target}}"{{end}} required></td>
	<td><input type="submit" name="do" value="Add">
</tr>
</form>

{{with .Links}}
	<tr>
		<th>Short Link</th>
		<th>&nbsp;</th>
		<th>&nbsp;</th>
	</tr>
	{{range .}}
		<tr>
			<td><input class="autoselect" type="text" value="{{$.BaseURL}}/{{.Key}}" readonly></td>
			<td><input class="autoselect" type="text" value="{{.Target}}" readonly></td>
			<td>
				<form method="POST">
					<input type="hidden" name="key" value="{{.Key}}">
					<input type="submit" name="do" value="Delete" class="delete">
				</form>
			</td>
		</tr>
	{{end}}
{{end}}
</table>
<script>
document.querySelectorAll('.autoselect').forEach(el => {
	el.addEventListener('click', e => {
		e.target.select();
	});
});

const baseURL = '{{$.BaseURL}}';
document.querySelectorAll('.delete').forEach(el => {
	el.addEventListener('click', e => {
		const key = e.target.form.elements['key'].value;
		if (!confirm('Delete this link?\n' + baseURL + '/' + key)) {
			e.preventDefault();
			return;
		}
	})
});
</script>
`
