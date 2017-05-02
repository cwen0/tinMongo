package utils

import "html/template"

var TemplateFuncMap = template.FuncMap{
	"set":   Set,
	"equal": Equal,
}

func Set(args map[string]interface{}, key string, value interface{}) template.JS {
	args[key] = value
	return template.JS("")
}

func Equal(args ...interface{}) bool {
	return args[0] == args[1]
}
