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

package pan_test

import (
	"testing"

	"github.com/EDyO/pan/pan"
)

var enclosure1 = pan.Enclosure{
	Length: "234875",
	Type:   "audio/mpeg",
	URL:    "http://link.to/episode1.mp3",
}

var enclosure2 = pan.Enclosure{
	Length: "4805409",
	Type:   "audio/mpeg",
	URL:    "http://link.to/episode2.mp3",
}

func TestEnclosureEqual(t *testing.T) {
	enclosureEqual := pan.Enclosure{
		Length: enclosure1.Length,
		Type:   enclosure1.Type,
		URL:    enclosure1.URL,
	}
	if !enclosure1.Equal(enclosureEqual) {
		t.Errorf("Enclosures should be equal:\n%s\n%s", enclosure1, enclosureEqual)
	}
	if enclosure1.Equal(enclosure2) {
		t.Errorf("Enclosures should not be equal:\n%s\n%s", enclosure1, enclosure2)
	}
}
