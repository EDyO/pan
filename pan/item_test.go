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

var item1 = pan.Item{
	Title:       "This is the title of the episode",
	Link:        "http://link.to/episode1.mp3",
	GUID:        "http://link.to/episode1.mp3",
	Description: "Small summary of the episode",
	PubDate:     "Sun, 26 Jan 2014 23:00:00 +0000",
	Enclosure:   enclosure1,
}

var item2 = pan.Item{
	Title:       "This is the title of another episode",
	Link:        "http://link.to/episode2.mp3",
	GUID:        "http://link.to/episode2.mp3",
	Description: "Small summary of the second episode",
	PubDate:     "Sun, 26 Feb 2014 23:00:00 +0000",
	Enclosure:   enclosure2,
}

var item3 = pan.Item{
	Title:       "This is the title of the episode",
	Link:        "http://link.to/episode1.mp3",
	GUID:        "http://link.to/episode1.mp3",
	Description: "Small summary of the episode",
	PubDate:     "Sun, 26 Jan 2014 23:00:00 +0000",
	Enclosure:   enclosure2,
}

var item4 = pan.Item{
	Title:       "May the fourth be with you",
	Link:        "http://link.to/episode4.mp3",
	GUID:        "http://link.to/episode4.mp3",
	Description: "A long time ago in a galaxy far, far away...",
	PubDate:     "Sun, 26 Jan 2015 23:00:00 +0000",
	Enclosure: pan.Enclosure{
		Length: "789345",
		Type:   "audio/mpeg",
		URL:    "http://link.to/episode4.mp3",
	},
}

var itemFixtures = []fixture{
	{
		name:   "item1",
		desc:   "Simple item",
		result: item1,
	},
	{
		name:   "item4",
		desc:   "Item without URL in Enclosure",
		result: item4,
	},
}

func TestItemUnmarshalYAML(t *testing.T) {
	for _, fixture := range itemFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			item := fixture.result.(pan.Item)
			if diff := deep.Equal(item, result.(pan.Item)); diff != nil {
				t.Errorf(
					"Loaded items should be equal:\n%s\n%s",
					item,
					result,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				item := pan.Item{}
				err := yaml.Unmarshal([]byte(content), &item)
				check(err)
				fixture.checkFail(item, t)
			},
		)
	}
}

func TestItemMarshalXML(t *testing.T) {
	for _, fixture := range itemFixtures {
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

var itemMap1 = map[interface{}]interface{}{
	"title":       "This is the title of the episode",
	"link":        "http://link.to/episode1.mp3",
	"guid":        "http://link.to/episode1.mp3",
	"description": "Small summary of the episode",
	"pubDate":     "Sun, 26 Jan 2014 23:00:00 +0000",
	"enclosure":   enclosureMap1,
}

func TestItemFromMap(t *testing.T) {
	item := pan.ItemFromMap(itemMap1)
	if diff := deep.Equal(item1, item); diff != nil {
		t.Errorf(
			"%s should be equal to %s",
			item,
			item1,
		)
	}
}
