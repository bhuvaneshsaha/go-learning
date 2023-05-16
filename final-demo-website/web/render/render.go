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
	// useCache determines if templates should be cached
	useCache = true // set this to false in development mode
)

// RenderTemplate renders the template with the given data
func RenderTemplate(w http.ResponseWriter, data interface{}, templateFiles ...string) {
	// Check if the template is already in the cache
	key := getTemplateCacheKey(templateFiles...)
	tmpl, ok := getTemplateFromCache(key)
	if !ok || !useCache {
		// Parse the HTML template files
		var err error
		tmpl, err = template.ParseFiles(templateFiles...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if useCache {
			// Add the parsed template to the cache
			addTemplateToCache(key, tmpl)
		}
	}

	// Render the template with the data
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getTemplateFromCache gets a template from cache if exists
func getTemplateFromCache(key string) (*template.Template, bool) {
	templateCacheMutex.Lock()
	defer templateCacheMutex.Unlock()
	tmpl, ok := templateCache[key]
	return tmpl, ok
}

// addTemplateToCache adds a template to cache
func addTemplateToCache(key string, tmpl *template.Template) {
	templateCacheMutex.Lock()
	defer templateCacheMutex.Unlock()
	templateCache[key] = tmpl
}

// getTemplateCacheKey generates a unique key for the given template files
func getTemplateCacheKey(templateFiles ...string) string {
	key := ""
	for _, file := range templateFiles {
		key += file
	}
	return key
}

// SetUseCache sets the useCache variable
func SetUseCache(cache bool) {
	useCache = cache
}
