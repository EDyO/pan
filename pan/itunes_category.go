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
	XMLName    xml.Name `xml:"itunes:category"`
	Text       string   `xml:"text,attr,omitempty"`
	Categories []ITunesCategory
}

// ITunesCategoryFromMap is an ITunesCategory factory from map[interface{}]interface{}.
func ITunesCategoryFromMap(iTunesCategoryMap map[interface{}]interface{}) ITunesCategory {
	attributesMap := map[interface{}]interface{}{}
	if iTunesCategoryMap["itunes_category"] != nil {
		categoryMap := iTunesCategoryMap["itunes_category"].(map[interface{}]interface{})
		attributesMap = categoryMap["attributes"].(map[interface{}]interface{})
	} else {
		attributesMap = iTunesCategoryMap["attributes"].(map[interface{}]interface{})
	}
	categoriesMap := []interface{}{}
	if iTunesCategoryMap["itunes_categories"] != nil {
		categoriesMap = iTunesCategoryMap["itunes_categories"].([]interface{})
	}
	categories := []ITunesCategory{}
	for _, categoryMap := range categoriesMap {
		category := ITunesCategoryFromMap(categoryMap.(map[interface{}]interface{}))
		categories = append(categories, category)
	}
	return ITunesCategory{
		Text:       attributesMap["text"].(string),
		Categories: categories,
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
	i.Categories = iTunesCategory.Categories
	return
}
