package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Title string    `form:"name"`
	Date  time.Time `form:"time"`
	Id    int       `form:"id"`
	text  string    `form:text`
}

//TODO use database
var AllEntries map[string]Entry = make(map[string]Entry)

func aboutPage(c *gin.Context) {
	// TODO
	// read .md file to string
	// render a html page, fill up a proper var with string
	// render that string with https://github.com/evilstreak/markdown-js
	md, err := ioutil.ReadFile("about/about.md")
	if err != nil {
		log.Println(err)
		c.String(404, "Not Found")
	}
	c.String(200, string(md))
}

func allEntries(c *gin.Context) {
	c.JSON(200, AllEntries)
}

func entry(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	entry, ok := AllEntries[id]
	if !ok {
		c.String(404, "Not Found")
	}
	c.JSON(200, entry)
}

func newEntry(c *gin.Context) {
	return
}
