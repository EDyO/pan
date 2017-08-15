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

var atomLink1 = pan.AtomLink{
	Href: "http://link.to/feed.xml",
	Rel:  "self",
	Type: "application/rss+xml",
}

var atomLink2 = pan.AtomLink{
	Href: "http://link2.to/another_feed.xml",
	Rel:  "self",
	Type: "application/rss+xml",
}

var atomLinkFixtures = []fixture{
	{
		name:   "atom_link1",
		desc:   "Simple atom:link",
		result: atomLink1,
	},
}

func TestAtomLinkUnmarshalYAML(t *testing.T) {
	for _, fixture := range atomLinkFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			atomLink := fixture.result.(pan.AtomLink)
			if diff := deep.Equal(atomLink, result.(pan.AtomLink)); diff != nil {
				t.Errorf(
					"Loaded atom:links should be equal:\n%s\n%s",
					atomLink,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				atomLink := pan.AtomLink{}
				err := yaml.Unmarshal([]byte(content), &atomLink)
				check(err)
				fixture.checkFail(atomLink, t)
			},
		)
	}
}

func TestAtomLinkMarshalXML(t *testing.T) {
	for _, fixture := range atomLinkFixtures {
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

var atomLinkMap1 = map[interface{}]interface{}{
	"href": "http://link.to/feed.xml",
	"rel":  "self",
	"type": "application/rss+xml",
}

func TestAtomLinkFromMap(t *testing.T) {
	atomLink := pan.AtomLinkFromMap(atomLinkMap1)
	if diff := deep.Equal(atomLink, atomLink1); diff != nil {
		t.Errorf(
			"%s should be equal to %s",
			atomLink,
			atomLink1,
		)
	}
}
