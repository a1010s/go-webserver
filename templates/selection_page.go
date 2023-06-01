package templates

const TemplateSelectionPage = `
<!DOCTYPE html>
<html>
<head>
	<title>Self-Service Infrastructure Portal</title>
	<style>
		/* CSS styles omitted for brevity */
	</style>
</head>
<body>
	<h1>Choose an infrastructure template:</h1>
	<form method="post" action="/provision">
		<select name="template">
			<option value="debian">Debian Server</option>
			<option value="ubuntu">Ubuntu Server</option>
			<option value="k8s">K8s Cluster</option>
			<option value="windows">Windows Server</option>
			<option value="linux VMs Ntwrk">2 Linux VMs Network</option>
		</select>
		<br>
		<input type="text" name="project_name" placeholder="Project Name" required>
		<br>
		<input type="text" name="token" placeholder="GitLab Token" required>
		<br>
		<input type="submit" value="Provision">
	</form>
</body>
</html>`
