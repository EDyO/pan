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
	"strconv"
)

// RSS represents a RSS Feed.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `yaml:"channel"`
}

// Equal returns true if rss is equal to r, false otherwise.
func (r *RSS) Equal(rss RSS) bool {
	if r.Version != rss.Version ||
		!r.Channel.Equal(rss.Channel) {
		return false
	}
	return true
}

// UnmarshalYAML is the unmarshaler for RSS.
func (r *RSS) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var rss map[string]interface{}
	if err = unmarshal(&rss); err != nil {
		return
	}
	attributes := rss["attributes"].(map[interface{}]interface{})
	r.Version = fmt.Sprintf("%.1f", attributes["version"].(float64))
	channel := rss["channel"].(map[interface{}]interface{})
	r.Channel = Channel{
		Title:       channel["title"].(string),
		Link:        channel["link"].(string),
		Language:    channel["language"].(string),
		Copyright:   channel["copyright"].(string),
		Description: channel["description"].(string),
		Items:       []Item{},
	}
	for _, itemElement := range channel["items"].([]interface{}) {
		itemMap := itemElement.(map[interface{}]interface{})
		enclosure := itemMap["enclosure"].(map[interface{}]interface{})
		enclosureAttrs := enclosure["attributes"].(map[interface{}]interface{})
		item := Item{
			Title:       itemMap["title"].(string),
			Link:        itemMap["link"].(string),
			GUID:        itemMap["link"].(string),
			Description: itemMap["description"].(string),
			PubDate:     itemMap["pubDate"].(string),
			Enclosure: Enclosure{
				Length: strconv.Itoa(enclosureAttrs["length"].(int)),
				Type:   enclosureAttrs["type"].(string),
				URL:    enclosureAttrs["url"].(string),
			},
		}
		r.Channel.Items = append(r.Channel.Items, item)
	}
	return
}
