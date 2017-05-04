package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/router"
	"github.com/cwen0/tinMongo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	setLogger()
	r := gin.Default()
	setTemplate(r)

	r.StaticFS("/public", http.Dir("public"))
	router.SetRoutes(r)
	r.Run(":3000")
}

//setTemplate loads templates from rice box "views"
func setTemplate(r *gin.Engine) {
	box := rice.MustFindBox("views")
	tmpl := template.New("").Funcs(utils.TemplateFuncMap)

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && (strings.HasSuffix(f.Name(), ".html") || strings.HasSuffix(f.Name(), ".tpl")) {
			var err error
			tmpl, err = tmpl.Parse(box.MustString(path))
			if err != nil {
				return err
			}
		}
		return nil
	}

	err := box.Walk("", fn)
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(tmpl)
}

func setLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stderr)
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
