package templates

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Read the about.md file
	mdFile, err := ioutil.ReadFile("templates/about.md")
	if err != nil {
		http.Error(w, "Failed to read about.md", http.StatusInternalServerError)
		return
	}

	// Create a new Markdown parser with extensions
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList, extension.Footnote, extension.Linkify, extension.Strikethrough, extension.Table),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	// Convert Markdown to HTML
	var htmlContent strings.Builder
	if err := md.Convert(mdFile, &htmlContent); err != nil {
		http.Error(w, "Failed to convert Markdown to HTML", http.StatusInternalServerError)
		return
	}

	// Set the response content type to HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the HTML content to the response, including CSS style
	htmlTemplate := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>About</title>
			<style>
				body {
					background-color: #202630;
					color: white;
					font-family: "Helvetica", sans-serif;
				}
			</style>
		</head>
		<body>
			%s
		</body>
		</html>
	`
	fmt.Fprintf(w, htmlTemplate, htmlContent.String())
}
