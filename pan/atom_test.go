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

	"gopkg.in/yaml.v2"

	"github.com/EDyO/pan/pan"
)

var atom1 = pan.AtomLink{
	Href: "http://link.to/feed.xml",
	Rel:  "self",
	Type: "application/rss+xml",
}

var atom2 = pan.AtomLink{
	Href: "http://link2.to/another_feed.xml",
	Rel:  "self",
	Type: "application/rss+xml",
}

func TestAtomLinkEqual(t *testing.T) {
	atomEqual := pan.AtomLink{
		Href: atom1.Href,
		Rel:  atom1.Rel,
		Type: atom1.Type,
	}
	if !atom1.Equal(atomEqual) {
		t.Errorf("Atoms should be equal:\n%s\n%s", atom1, atomEqual)
	}
	if atom1.Equal(atom2) {
		t.Errorf("Atoms should not be equal:\n%s\n%s", atom1, atom2)
	}
}

var atomFixtures = []fixture{
	{
		name:   "atom1",
		desc:   "Simple atom",
		result: atom1,
	},
}

func TestAtomLinkUnmarshalYAML(t *testing.T) {
	for _, fixture := range atomFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			atom := fixture.result.(pan.AtomLink)
			if !atom.Equal(result.(pan.AtomLink)) {
				t.Errorf(
					"Loaded atoms should be equal:\n%s\n%s",
					atom,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				atom := pan.AtomLink{}
				err := yaml.Unmarshal([]byte(content), &atom)
				check(err)
				fixture.checkFail(atom, t)
			},
		)
	}
}

func TestAtomLinkMarshalXML(t *testing.T) {
	for _, fixture := range atomFixtures {
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
