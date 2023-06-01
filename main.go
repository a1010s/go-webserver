package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"templates/templates"
)

const port = ":8080"

func main() {
	// Set up HTTP request handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/provision", provisionHandler)

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Display the template selection page
	fmt.Fprint(w, templates.TemplateSelectionPage)
}

func provisionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Only allow POST requests
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	template := r.FormValue("template")
	projectName := r.FormValue("project_name")
	token := r.FormValue("token")

	if template == "" || projectName == "" || token == "" {
		http.Error(w, "Invalid template, project name, or token", http.StatusBadRequest)
		return
	}

	switch template {
	case "debian":
		// Trigger the Debian pipeline
		err := triggerPipeline(token, "debian", projectName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to trigger pipeline: %v", err), http.StatusInternalServerError)
			return
		}
	case "ubuntu":
		// Trigger the Debian pipeline
		err := triggerPipeline(token, "ubuntu", projectName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to trigger pipeline: %v", err), http.StatusInternalServerError)
			return
		}
	case "k8s":
		// Trigger the K8s pipeline
		err := triggerPipeline(token, "k8s", projectName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to trigger pipeline: %v", err), http.StatusInternalServerError)
			return
		}
	case "windows":
		// Trigger the Windows pipeline
		err := triggerPipeline(token, "windows", projectName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to trigger pipeline: %v", err), http.StatusInternalServerError)
			return
		}
	default:
		// Invalid template selection
		http.Error(w, "Invalid template selection", http.StatusBadRequest)
		return
	}

	// Provisioning successful, update the response with a message
	message := fmt.Sprintf("Your project '%s' have been successfuly provisioned!", projectName)
	fmt.Fprintf(w, "<h1>%s</h1>", message)

}

func triggerPipeline(token, template, projectName string) error {
	// Define the project ID and the endpoint URL for triggering the pipeline
	projectID := "46...your Gitlab project ID"
	endpoint := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/trigger/pipeline?ref=main&token=%s", projectID, token)

	// Prepare the data payload for the HTTP request
	data := url.Values{}
	data.Set("variables[TEMPLATE_SELECTION]", template)
	data.Set("variables[PROJECT_NAME]", projectName)

	// Create an HTTP client
	client := &http.Client{}

	// Create a new POST request with the endpoint URL and encoded data payload
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	// Set the Content-Type header to indicate the form-urlencoded data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the HTTP request and get the response
	resp, err := client.Do(req)

	// Check the response status code
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("pipeline trigger failed with status code: %d", resp.StatusCode)
	}

	return nil
}
