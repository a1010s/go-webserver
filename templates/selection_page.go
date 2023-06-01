package templates

const TemplateSelectionPage = `
<!DOCTYPE html>
        <html>
        <head>
                <title>Self-Service Infrastructure Portal</title>
                <style>
                        body {
                                background-color: #222222;
                                color: white;
                                font-family: "Helvetica", sans-serif;
                        }

                        h1 {
                                color: yellow;
                        }

                        form {
                                margin: 20px;
                        }

                        select, input[type="text"], input[type="submit"] {
                                margin: 10px;
                                padding: 5px;
                                border-radius: 5px;
                        }

                        select, input[type="text"] {
                                width: 200px;
                        }

                        input[type="submit"] {
                                background-color: #555555;
                                color: white;
                        }
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
