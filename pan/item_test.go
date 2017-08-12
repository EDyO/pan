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

var item1 = pan.Item{
	Title:       "This is the title of the episode",
	Link:        "http://link.to/episode1.mp3",
	GUID:        "http://link.to/episode1.mp3",
	Description: "Small summary of the episode",
	PubDate:     "Sun, 26 Jan 2014 23:00:00 +0000",
}

var item2 = pan.Item{
	Title:       "This is the title of another episode",
	Link:        "http://link.to/episode2.mp3",
	GUID:        "http://link.to/episode2.mp3",
	Description: "Small summary of the second episode",
	PubDate:     "Sun, 26 Feb 2014 23:00:00 +0000",
}

func TestItemEqual(t *testing.T) {
	itemEqual := pan.Item{
		Title:       item1.Title,
		Link:        item1.Link,
		GUID:        item1.Link,
		Description: item1.Description,
		PubDate:     item1.PubDate,
	}
	if !item1.Equal(itemEqual) {
		t.Errorf("Items should be equal:\n%s\n%s", item1, itemEqual)
	}
	if item1.Equal(item2) {
		t.Errorf("Items should not be equal:\n%s\n%s", item1, item2)
	}
}

var fixtures = []fixture{
	{
		name: "item1",
		desc: "Simple item",
		item: item1,
	},
}

func TestItemUnmarshalYAML(t *testing.T) {
	for _, fixture := range fixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			item := fixture.item.(pan.Item)
			if !item.Equal(result.(pan.Item)) {
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
	for _, fixture := range fixtures {
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
					&fixture.item,
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
