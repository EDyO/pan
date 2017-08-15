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
	GUID        string   `xml:"guid"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Enclosure   Enclosure
}

// ItemFromMap is an Item factory from a map[interface{}]interface{}.
func ItemFromMap(itemMap map[interface{}]interface{}) Item {
	enclosureMap := itemMap["enclosure"].(map[interface{}]interface{})
	enclosure := EnclosureFromMap(enclosureMap)
	if enclosure.URL == "" {
		enclosure.URL = itemMap["link"].(string)
	}
	return Item{
		Title:       itemMap["title"].(string),
		Link:        itemMap["link"].(string),
		GUID:        itemMap["link"].(string),
		Description: itemMap["description"].(string),
		PubDate:     itemMap["pubDate"].(string),
		Enclosure:   enclosure,
	}
}

// UnmarshalYAML is the unmarshaler for Item.
func (i *Item) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var itemMap map[interface{}]interface{}
	if err = unmarshal(&itemMap); err != nil {
		return
	}
	item := ItemFromMap(itemMap)
	i.Title = item.Title
	i.Link = item.Link
	i.GUID = item.GUID
	i.Description = item.Description
	i.PubDate = item.PubDate
	i.Enclosure = item.Enclosure
	return
}
