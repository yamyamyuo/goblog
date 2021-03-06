package main

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Title string    `form:"title"`
	Date  time.Time `form:"time"`
	Id    string    `form:"id"`
	Text  string    `form:"text"`
}

var AllEntries map[string]Entry = make(map[string]Entry)

func aboutPage(c *gin.Context) {
	// TODO
	// read .md file to string
	// render a html page, fill up a proper var with string
	// render that string with https://github.com/evilstreak/markdown-js
	md, err := ioutil.ReadFile("about/about.md")
	if err != nil {
		c.String(404, "Not Found\n")
	}
	c.String(200, string(md))
}

func allEntries(c *gin.Context) {
	c.JSON(200, AllEntries)
}

func entry(c *gin.Context) {
	id := c.Param("id")
	entry, ok := AllEntries[id]
	if !ok {
		c.String(404, "Not Found\n")
	} else {
		c.JSON(200, entry)
	}
}

func newEntry(c *gin.Context) {
	var entry Entry
	c.Bind(&entry)
	// Use title as UID, so entry title should be unique
	lst := strings.Fields(removePunc(entry.Title))
	entry.Id = strings.Join(lst, "-")
	entry.Date = time.Now()
	AllEntries[entry.Id] = entry

	// year, month, day := entry.Date.Date()
	c.JSON(201, entry)
}

func deleteEntry(c *gin.Context) {
	id := c.Param("id")
	_, ok := AllEntries[id]
	if !ok {
		c.String(404, "Not Found\n")
	} else {
		delete(AllEntries, id)
		c.JSON(200, "deleted")
	}
}
