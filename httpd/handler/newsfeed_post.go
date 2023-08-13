package handler

import (
	"net/http"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

// news feed post request struct to pass it into the request
type NewsFeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// main
func NewsFeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		requestBody := NewsFeedPostRequest{}

		// use bind method to bind the request body to the incoming request
		c.Bind(&requestBody)

		// add the news feed item
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		// we use c.Status(http.StatusNoContent) isntead of c.JSON(http.StatusOK, "insett content") when
		// we need to send the ok status with no body :)
		c.Status(http.StatusNoContent)
	}
}
