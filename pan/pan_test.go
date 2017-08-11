// Copyright © 2017 Ignasi Fosch
//
//  This file is part of pan.
//
//  pan is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Lesser General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  pan is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Lesser General Public License for more details.
//
//  You should have received a copy of the GNU Lesser General Public License
//  along with pan. If not, see <http://www.gnu.org/licenses/>.
//

package pan_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type fixture struct {
	name      string
	desc      string
	item      interface{}
	checkFail func(interface{}, *testing.T)
}

func (f *fixture) load(format string) (content string) {
	fixtureContent, err := ioutil.ReadFile(
		fmt.Sprintf("fixtures/%s.%s", f.name, strings.ToLower(format)),
	)
	check(err)
	content = string(fixtureContent)
	return
}
