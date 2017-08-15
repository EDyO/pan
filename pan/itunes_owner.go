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

// ITunesOwner represents an ITunesOwner tag.
type ITunesOwner struct {
	XMLName xml.Name `xml:"itunes:owner"`
	Name    string   `yaml:"itunes_name" xml:"itunes:name,omitempty"`
	Email   string   `yaml:"itunes_email" xml:"itunes:email,omitempty"`
}

// ITunesOwnerFromMap is an ITunesOwner factory from map[interface{}]interface{}.
func ITunesOwnerFromMap(iTunesOwnerMap map[interface{}]interface{}) ITunesOwner {
	return ITunesOwner{
		Name:  iTunesOwnerMap["itunes_name"].(string),
		Email: iTunesOwnerMap["itunes_email"].(string),
	}
}
