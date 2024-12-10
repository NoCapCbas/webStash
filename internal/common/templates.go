package common

import (
	"encoding/json"
	"html/template"
)

// In your template setup
var FuncMap = template.FuncMap{
	"jsEscape": func(str interface{}) template.JS {
		b, err := json.Marshal(str)
		if err != nil {
			return template.JS("''")
		}
		return template.JS(b)
	},
}
