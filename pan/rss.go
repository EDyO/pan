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

// RSS represents a RSS Feed.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel
}

// Equal returns true if rss is equal to r, false otherwise.
func (r *RSS) Equal(rss RSS) bool {
	if !r.Channel.Equal(rss.Channel) {
		return false
	}
	return true
}
