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

var channel1 = pan.Channel{
	AtomLink:       &atomLink1,
	ITunesSubtitle: "When something means something",
	ITunesAuthor:   "Somebody",
	ITunesExplicit: "No",
	ITunesSummary:  "This podcast is about something",
	ITunesImage:    &iTunesImage1,
	Title:          "Something Podcast",
	Link:           "http://link.to",
	Language:       "en-us",
	Copyright:      "creative commons - Attribution - Non commercial - Share Alike - http://creativecommons.org/licenses/by-nc-sa/3.0/deed.en_US",
	Description:    "This podcast is about something",
	Items:          []pan.Item{item1},
}

var channel2 = pan.Channel{
	Title:       "Some other Podcast",
	Link:        "http://link2.to",
	Language:    "en-uk",
	Copyright:   "creative commons - Attribution - Non commercial - Share Alike - http://creativecommons.org/licenses/by-nc-sa/3.0/deed.en_US",
	Description: "This podcast is about something else",
}

var channel3 = pan.Channel{
	Title:       "Something Podcast",
	Link:        "http://link.to",
	Language:    "en-us",
	Copyright:   "creative commons - Attribution - Non commercial - Share Alike - http://creativecommons.org/licenses/by-nc-sa/3.0/deed.en_US",
	Description: "This podcast is about something",
	Items: []pan.Item{
		item1,
		item2,
	},
}

var channel4 = pan.Channel{
	AtomLink:    &atomLink2,
	Title:       "Something Podcast",
	Link:        "http://link.to",
	Language:    "en-us",
	Copyright:   "creative commons - Attribution - Non commercial - Share Alike - http://creativecommons.org/licenses/by-nc-sa/3.0/deed.en_US",
	Description: "This podcast is about something",
	Items:       []pan.Item{item1},
}

var channelFixtures = []fixture{
	{
		name:   "channel1",
		desc:   "Simple channel",
		result: channel1,
	},
	{
		name:   "channel3",
		desc:   "Simple channel with two episodes",
		result: channel3,
	},
}

func TestChannelUnmarshalYAML(t *testing.T) {
	for _, fixture := range channelFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			channel := fixture.result.(pan.Channel)
			if diff := deep.Equal(channel, result.(pan.Channel)); diff != nil {
				t.Errorf(
					"Loaded channels should be equal:\n%s\n%s\n%s",
					channel,
					result,
					diff,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				channel := pan.Channel{}
				err := yaml.Unmarshal([]byte(content), &channel)
				check(err)
				fixture.checkFail(channel, t)
			},
		)
	}
}

func TestChannelMarshalXML(t *testing.T) {
	for _, fixture := range channelFixtures {
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

var channelMap1 = map[interface{}]interface{}{
	"atom_link":       atomLinkMap1,
	"itunes_subtitle": "When something means something",
	"itunes_author":   "Somebody",
	"itunes_explicit": false,
	"itunes_summary":  "This podcast is about something",
	"itunes_image":    iTunesImageMap1,
	"title":           "Something Podcast",
	"link":            "http://link.to",
	"language":        "en-us",
	"copyright":       "creative commons - Attribution - Non commercial - Share Alike - http://creativecommons.org/licenses/by-nc-sa/3.0/deed.en_US",
	"description":     "This podcast is about something",
	"items":           []interface{}{itemMap1},
}

func TestChannelFromMap(t *testing.T) {
	channel := pan.ChannelFromMap(channelMap1)
	if diff := deep.Equal(channel1, channel); diff != nil {
		t.Errorf(
			"%s should be equal to %s",
			channel,
			channel1,
		)
	}
}
