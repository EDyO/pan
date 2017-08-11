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
)

// Item represents each episode.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
}

// Equal returns true if item2 is equal to i, false otherwise.
func (i *Item) Equal(item Item) bool {
	if i.Title != item.Title {
		return false
	}
	if i.Link != item.Link {
		return false
	}
	if i.Description != item.Description {
		return false
	}
	if i.PubDate != item.PubDate {
		return false
	}
	return true
}

// UnmarshalYAML is the unmarshaler for Item.
func (i *Item) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var item map[string]interface{}
	if err = unmarshal(&item); err != nil {
		return
	}
	i.Title = item["title"].(string)
	i.Link = item["link"].(string)
	i.Description = item["description"].(string)
	i.PubDate = item["pubDate"].(string)
	return
}
