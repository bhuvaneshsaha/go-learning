// package render

// import (
// 	"html/template"
// 	"net/http"
// )

// // renderTemplate renders the template with the given data
// func RenderTemplate(w http.ResponseWriter, data interface{}, templateFiles ...string) {
// 	// Parse the HTML template files
// 	tmpl, err := template.ParseFiles(templateFiles...)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Render the template with the data
// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

package render

import (
	"html/template"
	"net/http"
	"sync"
)

var (
	// templateCache stores parsed templates
	templateCache = make(map[string]*template.Template)
	// mutex for templateCache
	templateCacheMutex sync.Mutex
)

// RenderTemplate renders the template with the given data
func RenderTemplate(w http.ResponseWriter, data interface{}, templateFiles ...string) {
	// Check if the template is already in the cache
	key := getTemplateCacheKey(templateFiles...)
	templateCacheMutex.Lock()
	tmpl, ok := templateCache[key]
	templateCacheMutex.Unlock()
	if !ok {
		// Parse the HTML template files
		var err error
		tmpl, err = template.ParseFiles(templateFiles...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Add the parsed template to the cache
		templateCacheMutex.Lock()
		templateCache[key] = tmpl
		templateCacheMutex.Unlock()
	}

	// Render the template with the data
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getTemplateCacheKey generates a unique key for the given template files
func getTemplateCacheKey(templateFiles ...string) string {
	key := ""
	for _, file := range templateFiles {
		key += file
	}
	return key
}
