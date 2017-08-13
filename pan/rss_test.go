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

var rss1 = pan.RSS{
	Channel: channel1,
}

var rss2 = pan.RSS{
	Channel: channel2,
}

func TestRSSEqual(t *testing.T) {
	rssEqual := pan.RSS{
		Channel: rss1.Channel,
	}
	if !rss1.Equal(rssEqual) {
		t.Errorf("RSSs should be equal:\n%s\n%s", rss1, rssEqual)
	}
	if rss1.Equal(rss2) {
		t.Errorf("RSSs should not be equal:\n%s\n%s", rss1, rss2)
	}
}
