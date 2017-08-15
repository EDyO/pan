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

package pan

import (
	"encoding/xml"
	"fmt"
)

// RSS represents a RSS Feed.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	AtomNS  string   `xml:"xmlns:atom,attr"`
	Version string   `xml:"version,attr"`
	Channel Channel  `yaml:"channel"`
}

// RSSFromMap is a RSS factory from map[interface{}]interface{}.
func RSSFromMap(rssMap map[interface{}]interface{}) RSS {
	attributes := rssMap["attributes"].(map[interface{}]interface{})
	namespaces := rssMap["namespaces"].(map[interface{}]interface{})
	atomNS := ""
	channelMap := rssMap["channel"].(map[interface{}]interface{})
	for key, content := range namespaces {
		if key == "atom" {
			atomNS = content.(string)
		}
	}
	return RSS{
		Version: fmt.Sprintf("%.1f", attributes["version"].(float64)),
		AtomNS:  atomNS,
		Channel: ChannelFromMap(channelMap),
	}
}

// UnmarshalYAML is the unmarshaler for RSS.
func (r *RSS) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var rssMap map[interface{}]interface{}
	if err = unmarshal(&rssMap); err != nil {
		return
	}
	rss := RSSFromMap(rssMap)
	r.Version = rss.Version
	r.AtomNS = rss.AtomNS
	r.Channel = rss.Channel
	return
}
