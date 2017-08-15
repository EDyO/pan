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
	"encoding/xml"
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/yaml.v2"

	"github.com/EDyO/pan/pan"
)

var enclosure1 = pan.Enclosure{
	Length: "234875",
	Type:   "audio/mpeg",
	URL:    "http://link.to/episode1.mp3",
}

var enclosure2 = pan.Enclosure{
	Length: "4805409",
	Type:   "audio/mpeg",
	URL:    "http://link.to/episode2.mp3",
}

var enclosureFixtures = []fixture{
	{
		name:   "enclosure1",
		desc:   "Simple enclosure",
		result: enclosure1,
	},
}

func TestEnclosureUnmarshalYAML(t *testing.T) {
	for _, fixture := range enclosureFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			enclosure := fixture.result.(pan.Enclosure)
			if diff := deep.Equal(enclosure, result); diff != nil {
				t.Errorf(
					"Loaded enclosures should be equal:\n%s\n%s",
					enclosure,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				enclosure := pan.Enclosure{}
				err := yaml.Unmarshal([]byte(content), &enclosure)
				check(err)
				fixture.checkFail(enclosure, t)
			},
		)
	}
}

func TestEnclosureMarshalXML(t *testing.T) {
	for _, fixture := range enclosureFixtures {
		content := fixture.load("xml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			if content != result.(string) {
				t.Errorf(
					"XML strings should be equal:\n%s\n%s",
					content,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				b, err := xml.MarshalIndent(
					&fixture.result,
					"",
					"  ",
				)
				check(err)
				result := xml.Header + string(b) + "\n"
				fixture.checkFail(result, t)
			},
		)
	}
}

var enclosureMap1 = map[interface{}]interface{}{
	"attributes": map[interface{}]interface{}{
		"length": 234875,
		"type":   "audio/mpeg",
		"url":    "http://link.to/episode1.mp3",
	},
}

func TestEnclosureFromMap(t *testing.T) {
	enclosure := pan.EnclosureFromMap(enclosureMap1)
	if diff := deep.Equal(enclosure1, enclosure); diff != nil {
		t.Errorf(
			"%s should be equal to %s",
			enclosure,
			enclosure1,
		)
	}
}
