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

var rss1 = pan.RSS{
	AtomNS:   "http://www.w3.org/2005/Atom",
	ITunesNS: "http://www.itunes.com/dtds/podcast-1.0.dtd",
	Version:  "2.0",
	Channel:  channel1,
}

var rss2 = pan.RSS{
	Version: "2.0",
	Channel: channel2,
}

var rssFixtures = []fixture{
	{
		name:   "rss1",
		desc:   "Simple RSS",
		result: rss1,
	},
}

func TestRSSUnmarshalYAML(t *testing.T) {
	for _, fixture := range rssFixtures {
		content := fixture.load("yml")
		fixture.checkFail = func(result interface{}, t *testing.T) {
			rss := fixture.result.(pan.RSS)
			if diff := deep.Equal(rss, result.(pan.RSS)); diff != nil {
				t.Errorf(
					"Loaded RSSs should be equal:\n%s\n%s\n%s",
					rss,
					result,
					diff,
				)
			}
		}
		t.Run(
			fixture.desc,
			func(t *testing.T) {
				rss := pan.RSS{}
				err := yaml.Unmarshal([]byte(content), &rss)
				check(err)
				fixture.checkFail(rss, t)
			},
		)
	}
}

func TestRSSMarshalXML(t *testing.T) {
	for _, fixture := range rssFixtures {
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

var rssMap1 = map[interface{}]interface{}{
	"namespaces": map[interface{}]interface{}{
		"atom":   "http://www.w3.org/2005/Atom",
		"itunes": "http://www.itunes.com/dtds/podcast-1.0.dtd",
	},
	"attributes": map[interface{}]interface{}{
		"version": 2.0,
	},
	"channel": channelMap1,
}

func TestRSSFromMap(t *testing.T) {
	rss := pan.RSSFromMap(rssMap1)
	if diff := deep.Equal(rss1, rss); diff != nil {
		t.Errorf(
			"%s should be equal to %s\n%s",
			rss,
			rss1,
			diff,
		)
	}
}
