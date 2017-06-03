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
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// YMLItem represents each episode.
type YMLItem struct {
	Description string `json:"description"`
	Link        string `json:"link"`
	PubDate     string `json:"pubDate"`
	Title       string `json:"title"`
}

// YMLRss represents a RSS feed.
type YMLRss struct {
	Title string    `json:"title"`
	Items []YMLItem `json:"items"`
}

// XMLItem represents each episode.
type XMLItem struct {
	XMLName xml.Name `xml:"item"`

	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Title       string `xml:"title"`
}

// XMLRss represents a RSS feed.
type XMLRss struct {
	XMLName xml.Name `xml:"rss"`

	Title string    `xml:"channel>title"`
	Items []XMLItem `xml:"channel>item"`
}

// readYML tries to unmarshall YML content.
func readYML(content []byte) (feed YMLRss, err error) {
	err = yaml.Unmarshal(content, &feed)
	return
}

// YML2XML converts the struct from YML to XML.
func YML2XML(feed YMLRss) (output XMLRss, err error) {
	items := []XMLItem{}
	for _, item := range feed.Items {
		newItem := XMLItem{
			XMLName:     xml.Name{Local: "item"},
			Description: item.Description,
			Link:        item.Link,
			PubDate:     item.PubDate,
			Title:       item.Title,
		}
		items = append(items, newItem)
	}
	output = XMLRss{
		XMLName: xml.Name{Local: "rss"},
		Title:   feed.Title,
		Items:   items,
	}
	return
}

// readXML tries to unmarshall XML content.
func readXML(content []byte) (feed XMLRss, err error) {
	err = xml.Unmarshal(content, &feed)
	return
}

// XML2YML converts the struct from XML to YML.
func XML2YML(feed XMLRss) (output YMLRss, err error) {
	items := []YMLItem{}
	for _, item := range feed.Items {
		newItem := YMLItem{
			Description: item.Description,
			Link:        item.Link,
			PubDate:     item.PubDate,
			Title:       item.Title,
		}
		items = append(items, newItem)
	}
	output = YMLRss{
		Title: feed.Title,
		Items: items,
	}
	return
}

// Process runs required operations for all arguments.
func Process(args []string) (err error) {
	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		return err
	}

	XMLFeed, err := readXML(content)
	if err != nil {
		YMLFeed, err := readYML(content)
		if err != nil {
			return err
		}
		XMLFeed, err := YML2XML(YMLFeed)
		if err != nil {
			return err
		}
		output, err := xml.Marshal(XMLFeed)
		if err != nil {
			return err
		}
		fmt.Printf("%s%s", xml.Header, string(output))
		return nil
	}
	YMLFeed, err := XML2YML(XMLFeed)
	if err != nil {
		return err
	}
	output, err := yaml.Marshal(YMLFeed)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
