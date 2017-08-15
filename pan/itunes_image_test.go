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

var iTunesImage1 = pan.ITunesImage{
	Href: "http://link.to/podcast.png",
}

var iTunesImage2 = pan.ITunesImage{
	Href: "http://link2.to/another_podcast.png",
}

var iTunesImageFixtures = []fixture{
	{
		name:   "itunes_image1",
		desc:   "Simple itunes:image",
		result: iTunesImage1,
	},
}

func TestITunesImageUnmarshalYAML(t *testing.T) {
	for _, fixture := range iTunesImageFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			iTunesImage := fixture.result.(pan.ITunesImage)
			if diff := deep.Equal(iTunesImage, result.(pan.ITunesImage)); diff != nil {
				t.Errorf(
					"Loaded itunes:image should be equal:\n%s\n%s",
					iTunesImage,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				iTunesImage := pan.ITunesImage{}
				err := yaml.Unmarshal([]byte(content), &iTunesImage)
				check(err)
				fixture.checkFail(iTunesImage, t)
			},
		)
	}
}

func TestITunesImageMarshalXML(t *testing.T) {
	for _, fixture := range iTunesImageFixtures {
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

var iTunesImageMap1 = map[interface{}]interface{}{
	"attributes": map[interface{}]interface{}{
		"href": "http://link.to/podcast.png",
	},
}

func TestITunesImageFromMap(t *testing.T) {
	iTunesImage := pan.ITunesImageFromMap(iTunesImageMap1)
	if diff := deep.Equal(iTunesImage, iTunesImage1); diff != nil {
		t.Errorf(
			"%s should be equal to %s",
			iTunesImage,
			iTunesImage1,
		)
	}
}
