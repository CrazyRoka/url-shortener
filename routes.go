package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

type ShortenRequest struct {
	Url string `json:"url" binding:"required"`
}

func CreateRoutes() *gin.Engine {
	log.Info("Creating routes")

	r := gin.New()
	r.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())
	r.POST("/shorten", ShortenUrl)
	r.GET("/:short", GetUrl)

	return r
}

func ShortenUrl(c *gin.Context) {
	var shortenRequest ShortenRequest
	if err := c.ShouldBindJSON(&shortenRequest); err != nil {
		log.WithError(err).Warn("Error parsing request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uri, err := url.ParseRequestURI(shortenRequest.Url)
	if err != nil {
		log.WithError(err).WithField("uri", shortenRequest.Url).Warn("Error parsing request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortLink := BuildShortenLink(uri)
	log.WithFields(log.Fields{
		"link": shortLink,
		"url":  uri,
	}).Info("Generating short link")

	if err := InsertShortLink(&Session, &shortLink); err != nil {
		log.WithError(err).WithField("link", shortLink).Warn("Error inserting into the database")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, shortLink)
}

func GetUrl(c *gin.Context) {
	short := c.Param("short")
	log.WithField("short", short).Info("Received get request")

	link, err := GetShortLink(&Session, short)
	if err != nil {
		log.WithError(err).WithField("short", short).Warn("Error fetching short link")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Url not found",
		})
		return
	}
	log.WithField("link", link).Printf("Found link")

	url, err := url.ParseRequestURI(link.Url)
	if err != nil {
		log.WithError(err).WithField("link", link).Fatal("Error parsing request uri")
	}
	log.WithFields(log.Fields{
		"url":   url,
		"link":  link,
		"short": short,
	}).Info("Found url for short")

	c.Redirect(http.StatusPermanentRedirect, link.Url)
}
