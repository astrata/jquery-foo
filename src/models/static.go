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
	"os"
	"strings"
)

// Root directory for serving static files.
var Root = "static"

// Model name
type Static struct {
}

// Initialization function
func init() {
	// Your initialization code goes here.
	app.Register("Static", &Static{})
	app.Fallback("/", app.App("Static"))
}

// Model's StartUp() function
func (self *Static) StartUp() {
	// Code to be executed when all the models are loaded and fully initialized.
	info, err := os.Stat(Root)
	if err == nil {
		if info.IsDir() == false {
			panic(fmt.Sprintf("%s is not a directory.\n", Root))
		}
	} else {
		panic(err.Error())
	}
}

// Catches all requests and serves files.
func (self *Static) CatchAll(path ...string) body.Body {

	content := body.File()

	filename := Root + tango.PS + strings.Trim(strings.Join(path, tango.PS), tango.PS)

	info, err := os.Stat(filename)

	if err == nil {

		if info.IsDir() == true {

			filename = strings.Trim(filename, tango.PS) + tango.PS + "index.html"

			info, err = os.Stat(filename)

			if err != nil {
				return nil
			}

			if info.IsDir() == true {
				return nil
			}

		}

		content.Set(filename)

		return content
	}

	return nil
}
