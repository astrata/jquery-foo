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
	"encoding/json"
	"github.com/gosexy/to"
	"regexp"
	"bytes"
	"strings"
	"os"
)

// Root directory for serving static files.
var PluginsRoot = "plugins"


// Plugin context
type Context struct {
	Loaded map[string] bool
	Buf *bytes.Buffer
}

// Model name
type Build struct {
	Params tango.Value
	pattern *regexp.Regexp
}

// Initialization function
func init() {
	// Your initialization code goes here.
	app.Register("Build", &Build{})
	app.Route("/", app.App("Build"))
}

func newContext() *Context {
	self := &Context{}
	self.Buf = bytes.NewBuffer(nil)
	self.Loaded = make(map[string] bool)
	return self
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

func (self *Build) read(file string) []byte {
	var err error

	st, err := os.Stat(file)

	if err != nil {
		return nil
	}

	fh, err := os.Open(file)

	if err != nil {
		return nil
	}

	defer fh.Close()

	buf := make([]byte, st.Size())

	fh.Read(buf)

	return buf
}

func (self *Build) load(pkg string, ctx *Context) {
	var err error
	var version string

	minjs, _ := regexp.Compile(`\.js$`)

	if ctx.Loaded[pkg] == true {
		return
	} else {
		ctx.Loaded[pkg] = true
	}

	if strings.Contains(pkg, ":") {
		i := strings.LastIndex(pkg, ":")
		version = pkg[i+1:]
		pkg = pkg[0:i]
	}

	filename := PluginsRoot + tango.PS + pkg + tango.PS + "package.yaml"

	_, err = os.Stat(filename)

	if err == nil {

		info, err := yaml.Open(filename)

		if err == nil {

			if version == "" {
				version = to.String(info.Get("latest"))
			}

			files := to.Map(info.Get(fmt.Sprintf("packages/%s", version)))

			if files == nil {
				ctx.Buf.Write([]byte(fmt.Sprintf("/* Package \"%s\" with version \"%s\" was not found. */\n", pkg, version)))
			} else {

				if files["requires"] != nil {
					for _, req := range to.List(files["requires"]) {
						self.load(req.(string), ctx)
					}
				}

				ctx.Buf.Write([]byte(fmt.Sprintf("/* %s: %s. %s */\n", pkg, to.String(info.Get("name")), to.String(info.Get("copyright")))))

				if files["source"] != nil {
					for _, jsfile := range to.List(files["source"]) {
						minfile := minjs.ReplaceAllString(jsfile.(string), ".min.js")
						ctx.Buf.Write(self.read(PluginsRoot + tango.PS + pkg + tango.PS + minfile))
						ctx.Buf.Write([]byte(fmt.Sprintf("\n")))
					}
				}

				ctx.Buf.Write([]byte(fmt.Sprintf("\n\n")))

				if files["style"] != nil {

					styles := to.List(files["style"])

					cssfiles := make([]string, len(styles))

					for i, _ := range styles {
						cssfiles[i] = "media" + tango.PS + pkg + tango.PS + styles[i].(string)
					}

					css, _ := json.Marshal(cssfiles)

					ctx.Buf.Write([]byte(fmt.Sprintf("$.foo.styles.apply($.foo, %s);\n", string(css))))
				}

			}

			ctx.Buf.Write([]byte(fmt.Sprintf("\n\n")))

		} else {
			ctx.Buf.Write([]byte(fmt.Sprintf("/* Package \"%s\": metadata error. */\n", pkg)))
		}
	} else {
		ctx.Buf.Write([]byte(fmt.Sprintf("/* Package \"%s\": missing. */\n", pkg)))
	}


}

// Catches all requests and serves files.
func (self *Build) Index() body.Body {

	content := body.Text()

	params := self.Params.Filter("load")

	load := self.pattern.ReplaceAllString(params.Get("load"), "")

	pkgs := strings.Split(load, ",")

	ctx := newContext()

	for _, pkg := range pkgs {
		self.load(pkg, ctx)
	}

	content.Set(ctx.Buf)

	return content
}
