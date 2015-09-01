package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"crypto/md5"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Title string    `form:"title"`
	Date  time.Time `form:"time"`
	Id    string    `form:"id"`
	Text  string    `form:"text"`
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
		c.String(404, "Not Found\n")
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
		c.String(404, "Not Found\n")
	} else {
		c.JSON(200, entry)
	}
}

func newEntry(c *gin.Context) {
	var entry Entry
	c.Bind(&entry)
	// Use title as UID, so entry title should be unique
	// Base64 or MD5?
	md5ByteList := md5.Sum([]byte(entry.Title))
	computedCheckSum := fmt.Sprintf("%x", md5ByteList)
	entry.Id = string(computedCheckSum)
	entry.Date = time.Now()

	log.Println(entry)

	// year, month, day := entry.Date.Date()
	c.JSON(201, entry)
}

func deleteEntry(c *gin.Context) {
	var entry Entry
	c.Bind(&entry)
	log.Println(entry.Id)
	c.JSON(201, "deleted")
}
