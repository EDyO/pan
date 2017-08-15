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

var iTunesOwner1 = pan.ITunesOwner{
	Name:  "Somebody",
	Email: "somebody@link.to",
}

var iTunesOwner2 = pan.ITunesOwner{
	Name:  "Someone",
	Email: "someone@link.to",
}

var iTunesOwnerFixtures = []fixture{
	{
		name:   "itunes_owner1",
		desc:   "Simple itunes:owner",
		result: iTunesOwner1,
	},
}

func TestITunesOwnerUnmarshalYAML(t *testing.T) {
	for _, fixture := range iTunesOwnerFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			iTunesOwner := fixture.result.(pan.ITunesOwner)
			if diff := deep.Equal(iTunesOwner, result.(pan.ITunesOwner)); diff != nil {
				t.Errorf(
					"Loaded itunes:owner should be equal:\n%s\n%s",
					iTunesOwner,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				iTunesOwner := pan.ITunesOwner{}
				err := yaml.Unmarshal([]byte(content), &iTunesOwner)
				check(err)
				fixture.checkFail(iTunesOwner, t)
			},
		)
	}
}

func TestITunesOwnerMarshalXML(t *testing.T) {
	for _, fixture := range iTunesOwnerFixtures {
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

var iTunesOwnerMap1 = map[interface{}]interface{}{
	"itunes_name":  "Somebody",
	"itunes_email": "somebody@link.to",
}

func TestITunesOwnerFromMap(t *testing.T) {
	iTunesOwner := pan.ITunesOwnerFromMap(iTunesOwnerMap1)
	if diff := deep.Equal(iTunesOwner, iTunesOwner1); diff != nil {
		t.Errorf(
			"%s should be equal to %s",
			iTunesOwner,
			iTunesOwner1,
		)
	}
}
