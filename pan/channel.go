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
	XMLName     xml.Name  `xml:"channel"`
	AtomLink    *AtomLink `yaml:"atom_link" xml:"atom:link,omitempty"`
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Language    string    `xml:"language"`
	Copyright   string    `xml:"copyright"`
	Description string    `xml:"description"`
	Items       []Item    `yaml:"items"`
}

// Equal returns true if channel is equal to c, false otherwise.
func (c *Channel) Equal(channel Channel) bool {
	if c.Title != channel.Title ||
		c.Link != channel.Link ||
		c.Language != channel.Language ||
		c.Copyright != channel.Copyright ||
		c.Description != channel.Description {
		return false
	}
	if c.AtomLink == nil && channel.AtomLink != nil ||
		c.AtomLink != nil && channel.AtomLink == nil {
		return false
	}
	if c.AtomLink != channel.AtomLink &&
		!(*c.AtomLink).Equal(*channel.AtomLink) {
		return false
	}
	if c.Items == nil && channel.Items != nil ||
		c.Items != nil && channel.Items == nil {
		return false
	}
	for i, item := range c.Items {
		if !item.Equal(channel.Items[i]) {
			return false
		}
	}
	return true
}
