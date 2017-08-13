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

var channel1 = pan.Channel{
	Title:       "Something Podcast",
	Link:        "http://link.to",
	Language:    "en-us",
	Copyright:   "creative commons - Attribution - Non commercial - Share Alike - http://creativecommons.org/licenses/by-nc-sa/3.0/deed.en_US",
	Description: "This podcast is about something",
}

var channel2 = pan.Channel{
	Title:       "Some other Podcast",
	Link:        "http://link2.to",
	Language:    "en-uk",
	Copyright:   "creative commons - Attribution - Non commercial - Share Alike - http://creativecommons.org/licenses/by-nc-sa/3.0/deed.en_US",
	Description: "This podcast is about something else",
}

func TestChannelEqual(t *testing.T) {
	channelEqual := pan.Channel{
		Title:       channel1.Title,
		Link:        channel1.Link,
		Language:    channel1.Language,
		Copyright:   channel1.Copyright,
		Description: channel1.Description,
	}
	if !channel1.Equal(channelEqual) {
		t.Errorf("Channels should be equal:\n%s\n%s", channel1, channelEqual)
	}
	if channel1.Equal(channel2) {
		t.Errorf("Channels should not be equal:\n%s\n%s", channel1, channel2)
	}
}
