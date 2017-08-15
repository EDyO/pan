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

// Channel represents each episode.
type Channel struct {
	XMLName        xml.Name  `xml:"channel"`
	AtomLink       *AtomLink `yaml:"atom_link" xml:"atom:link,omitempty"`
	ITunesSubtitle string    `yaml:"itunes_subtitle" xml:"itunes:subtitle,omitempty"`
	ITunesAuthor   string    `yaml:"itunes_author" xml:"itunes:author,omitempty"`
	ITunesExplicit string    `yaml:"itunes_explicit" xml:"itunes:explicit,omitempty"`
	Title          string    `xml:"title"`
	Link           string    `xml:"link"`
	Language       string    `xml:"language"`
	Copyright      string    `xml:"copyright"`
	Description    string    `xml:"description"`
	Items          []Item    `yaml:"items"`
}

// ChannelFromMap is a Channel factory form map[interface{}]interface{}.
func ChannelFromMap(channelMap map[interface{}]interface{}) Channel {
	atomLink := AtomLink{}
	if channelMap["atom_link"] != nil {
		atomLinkMap := channelMap["atom_link"].(map[interface{}]interface{})
		atomLink = AtomLinkFromMap(atomLinkMap)
	}
	items := []Item{}
	itemsList := channelMap["items"].([]interface{})
	for _, itemMap := range itemsList {
		item := ItemFromMap(itemMap.(map[interface{}]interface{}))
		items = append(items, item)
	}
	explicit := "No"
	if channelMap["itunes_explicit"].(bool) {
		explicit = "Yes"
	}
	return Channel{
		AtomLink:       &atomLink,
		ITunesSubtitle: channelMap["itunes_subtitle"].(string),
		ITunesAuthor:   channelMap["itunes_author"].(string),
		ITunesExplicit: explicit,
		Title:          channelMap["title"].(string),
		Link:           channelMap["link"].(string),
		Language:       channelMap["language"].(string),
		Copyright:      channelMap["copyright"].(string),
		Description:    channelMap["description"].(string),
		Items:          items,
	}
}
