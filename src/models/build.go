/*
  Tango!

  Copyright (c) 2012 Astrata Software, <http://astrata.mx>
  Written by José Carlos Nieto <xiam@menteslibres.org>

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package models

import (
	"fmt"
	"github.com/astrata/tango"
	"github.com/astrata/tango/app"
	"github.com/astrata/tango/body"
	"github.com/gosexy/yaml"
	"github.com/gosexy/to"
	"regexp"
	"bytes"
	"strings"
	"os"
)

// Root directory for serving static files.
var PluginsRoot = "plugins"

// Model name
type Build struct {
	Params tango.Value
	pattern *regexp.Regexp
}

// Initialization function
func init() {
	// Your initialization code goes here.
	app.Register("Build", &Build{})
	app.Route("/build", app.App("Build"))
}

// Model's StartUp() function
func (self *Build) StartUp() {
	// Code to be executed when all the models are loaded and fully initialized.
	info, err := os.Stat(PluginsRoot)
	if err == nil {
		if info.IsDir() == false {
			panic(fmt.Sprintf("%s is not a directory.\n", Root))
		}
	} else {
		panic(err.Error())
	}

	self.pattern, _ = regexp.Compile(`[^a-z0-9\-:._,]`)
}

// Catches all requests and serves files.
func (self *Build) Index() body.Body {

	content := body.Text()

	params := self.Params.Filter("load")

	load := self.pattern.ReplaceAllString(params.Get("load"), "")

	files := strings.Split(load, ",")

	/* Output */
	output := bytes.NewBuffer(nil)

	/* Packages list */
	pkgs := []string{}
	var pkg string

	/* Testing all files. */
	for _, file := range files {
		pkg = PluginsRoot + tango.PS + file + tango.PS + "package.yaml"

		stat, err := os.Stat(pkg)

		if err != nil || stat.IsDir() == true {
			content.Set(fmt.Sprintf("/* Package \"%s\" was not found. */", file))
			return content
		}

		pkgs = append(pkgs, pkg)
	}

	/* Reading package files */
	for _, pkg := range pkgs {
		y, err := yaml.Open(pkg)

		if err != nil {
			content.Set(fmt.Sprintf("/* Package \"%s\" is malformed. */", pkg))
			return content
		}

		output.Write([]byte(fmt.Sprintf(
				"/* %s */\n\n",
				to.String(y.Get("name")),
		)))

	}

	content.Set(output)

	return content
}
