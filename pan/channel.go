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
	XMLName          xml.Name     `xml:"channel"`
	AtomLink         *AtomLink    `yaml:"atom_link" xml:"atom:link,omitempty"`
	ITunesSubtitle   string       `yaml:"itunes_subtitle" xml:"itunes:subtitle,omitempty"`
	ITunesAuthor     string       `yaml:"itunes_author" xml:"itunes:author,omitempty"`
	ITunesExplicit   string       `yaml:"itunes_explicit" xml:"itunes:explicit,omitempty"`
	ITunesSummary    string       `yaml:"itunes_summary" xml:"itunes:summary,omitempty"`
	ITunesImage      *ITunesImage `yaml:"itunes_image" xml:"itunes:image,omitempty"`
	ITunesOwner      *ITunesOwner `yaml:"itunes_owner" xml:"itunes:owner,omitempty"`
	ITunesCategories []ITunesCategory
	Title            string `xml:"title"`
	Link             string `xml:"link"`
	Language         string `xml:"language"`
	Copyright        string `xml:"copyright"`
	Description      string `xml:"description"`
	Items            []Item `yaml:"items"`
}

// ChannelFromMap is a Channel factory form map[interface{}]interface{}.
func ChannelFromMap(channelMap map[interface{}]interface{}) Channel {
	var atomLink *AtomLink
	if channelMap["atom_link"] != nil {
		atomLinkMap := channelMap["atom_link"].(map[interface{}]interface{})
		atomLinkObject := AtomLinkFromMap(atomLinkMap)
		atomLink = &atomLinkObject
	}
	items := []Item{}
	itemsList := channelMap["items"].([]interface{})
	for _, itemMap := range itemsList {
		item := ItemFromMap(itemMap.(map[interface{}]interface{}))
		items = append(items, item)
	}
	explicit := ""
	if channelMap["itunes_explicit"] != nil {
		explicit = "No"
		if channelMap["itunes_explicit"].(bool) {
			explicit = "Yes"
		}
	}
	iTunesSubtitle := ""
	if channelMap["itunes_subtitle"] != nil {
		iTunesSubtitle = channelMap["itunes_subtitle"].(string)
	}
	iTunesAuthor := ""
	if channelMap["itunes_author"] != nil {
		iTunesAuthor = channelMap["itunes_author"].(string)
	}
	iTunesSummary := ""
	if channelMap["itunes_summary"] != nil {
		iTunesSummary = channelMap["itunes_summary"].(string)
	}
	var iTunesImage *ITunesImage
	if channelMap["itunes_image"] != nil {
		iTunesImageMap := channelMap["itunes_image"].(map[interface{}]interface{})
		iTunesImageObject := ITunesImageFromMap(iTunesImageMap)
		iTunesImage = &iTunesImageObject
	}
	var iTunesOwner *ITunesOwner
	if channelMap["itunes_owner"] != nil {
		iTunesOwnerMap := channelMap["itunes_owner"].(map[interface{}]interface{})
		iTunesOwnerObject := ITunesOwnerFromMap(iTunesOwnerMap)
		iTunesOwner = &iTunesOwnerObject
	}
	var iTunesCategories []ITunesCategory
	if channelMap["itunes_categories"] != nil {
		iTunesCategories = []ITunesCategory{}
		iTunesCategoriesMap := channelMap["itunes_categories"].([]interface{})
		for _, iTunesCategoryItem := range iTunesCategoriesMap {
			iTunesCategoryMap := iTunesCategoryItem.(map[interface{}]interface{})
			iTunesCategory := ITunesCategoryFromMap(iTunesCategoryMap)
			iTunesCategories = append(iTunesCategories, iTunesCategory)
		}
	}
	return Channel{
		AtomLink:         atomLink,
		ITunesSubtitle:   iTunesSubtitle,
		ITunesAuthor:     iTunesAuthor,
		ITunesExplicit:   explicit,
		ITunesSummary:    iTunesSummary,
		ITunesImage:      iTunesImage,
		ITunesOwner:      iTunesOwner,
		ITunesCategories: iTunesCategories,
		Title:            channelMap["title"].(string),
		Link:             channelMap["link"].(string),
		Language:         channelMap["language"].(string),
		Copyright:        channelMap["copyright"].(string),
		Description:      channelMap["description"].(string),
		Items:            items,
	}
}

// UnmarshalYAML is the YAML unmarshaler for Channel.
func (c *Channel) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var channelMap map[interface{}]interface{}
	if err = unmarshal(&channelMap); err != nil {
		return
	}
	channel := ChannelFromMap(channelMap)
	c.AtomLink = channel.AtomLink
	c.ITunesSubtitle = channel.ITunesSubtitle
	c.ITunesAuthor = channel.ITunesAuthor
	c.ITunesExplicit = channel.ITunesExplicit
	c.ITunesSummary = channel.ITunesSummary
	c.ITunesImage = channel.ITunesImage
	c.ITunesOwner = channel.ITunesOwner
	c.ITunesCategories = channel.ITunesCategories
	c.Title = channel.Title
	c.Link = channel.Link
	c.Language = channel.Language
	c.Copyright = channel.Copyright
	c.Description = channel.Description
	c.Items = channel.Items
	return
}
