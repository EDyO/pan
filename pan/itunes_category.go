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

// ITunesCategory represents an AtomLink tag.
type ITunesCategory struct {
	XMLName xml.Name `xml:"itunes:category"`
	Text    string   `xml:"text,attr,omitempty"`
}

// ITunesCategoryFromMap is an ITunesCategory factory from map[interface{}]interface{}.
func ITunesCategoryFromMap(iTunesCategoryMap map[interface{}]interface{}) ITunesCategory {
	attributesMap := iTunesCategoryMap["attributes"].(map[interface{}]interface{})
	return ITunesCategory{
		Text: attributesMap["text"].(string),
	}
}

// UnmarshalYAML is the YAML unmarshaler for ITunesCategory.
func (i *ITunesCategory) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var iTunesCategoryMap map[interface{}]interface{}
	if err = unmarshal(&iTunesCategoryMap); err != nil {
		return
	}
	iTunesCategory := ITunesCategoryFromMap(iTunesCategoryMap)
	i.Text = iTunesCategory.Text
	return
}
