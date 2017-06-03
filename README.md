# pan

`pan` is a RSS feed management tool.

## Installation

You must setup Go in your computer, and, then, go get pan:

    go get github.com/EDyO/pan
    
## Usage

`pan` allows you to convert a RSS feed in XML format an get a YAML representation, or the other way around:

    pan my_feed.xml > my_feed.yml
    pan my_feed.yml > my_feed.xml
