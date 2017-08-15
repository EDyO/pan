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

var iTunesCategory1 = pan.ITunesCategory{
	Text: "Society & Culture",
}

var iTunesCategory2 = pan.ITunesCategory{
	Text: "Technology",
}

var iTunesCategoryFixtures = []fixture{
	{
		name:   "itunes_category1",
		desc:   "Simple itunes:category",
		result: iTunesCategory1,
	},
}

func TestITunesCategoryUnmarshalYAML(t *testing.T) {
	for _, fixture := range iTunesCategoryFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			iTunesCategory := fixture.result.(pan.ITunesCategory)
			if diff := deep.Equal(iTunesCategory, result.(pan.ITunesCategory)); diff != nil {
				t.Errorf(
					"Loaded itunes:category should be equal:\n%s\n%s",
					iTunesCategory,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				iTunesCategory := pan.ITunesCategory{}
				err := yaml.Unmarshal([]byte(content), &iTunesCategory)
				check(err)
				fixture.checkFail(iTunesCategory, t)
			},
		)
	}
}

func TestITunesCategoryMarshalXML(t *testing.T) {
	for _, fixture := range iTunesCategoryFixtures {
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

var iTunesCategoryMap1 = map[interface{}]interface{}{
	"attributes": map[interface{}]interface{}{
		"text": "Society & Culture",
	},
}

func TestITunesCategoryFromMap(t *testing.T) {
	iTunesCategory := pan.ITunesCategoryFromMap(iTunesCategoryMap1)
	if diff := deep.Equal(iTunesCategory, iTunesCategory1); diff != nil {
		t.Errorf(
			"%s should be equal to %s",
			iTunesCategory,
			iTunesCategory1,
		)
	}
}
