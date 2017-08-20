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

package pan

import (
	"encoding/xml"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Process runs required operations for all arguments.
func Process(args []string) (result string, err error) {
	fileContent, err := ioutil.ReadFile(
		args[0],
	)
	if err != nil {
		return
	}
	content := string(fileContent)
	rss := RSS{}
	err = yaml.Unmarshal([]byte(content), &rss)
	if err != nil {
		return
	}
	b, err := xml.MarshalIndent(
		&rss,
		"",
		"  ",
	)
	if err != nil {
		return
	}
	result = xml.Header + string(b) + "\n"
	return
}
