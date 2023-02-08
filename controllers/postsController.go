package controllers

import (
	"demo-app/initializers"
	"demo-app/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var body struct {
	Title string
	Body  string
}

type ApiError struct {
	Field string
	Msg   string
}

// return custom error messages for validation errors
func MsgForTag(tag, field string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "alpha":
		return field + " must only contain alphabets"
	}
	return ""
}

func CreatePost(c *gin.Context) {

	fmt.Println("creating post")

	// get data off request body
	c.Bind(&body)

	// create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	//create new validator instance
	v := validator.New()
	//validate post
	err := v.Struct(post)
	// return validation error if exists
	if err != nil {
		var error_msgs []any
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				error_msgs = append(error_msgs, ApiError{fe.Field(), MsgForTag(fe.Tag(), fe.Field())})
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": error_msgs})
		}
		return
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	// get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//return posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostByID(c *gin.Context) {
	// get params from url
	id := c.Param(("id"))
	// get posts
	var posts models.Post
	initializers.DB.First(&posts, id)

	//return posts
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func UpdatePost(c *gin.Context) {
	//get id off the url
	id := c.Param(("id"))

	//get data off the req body
	c.Bind(&body)

	//find the post we want to update
	var post models.Post
	initializers.DB.First(&post, id)

	//update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//return updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePostByID(c *gin.Context) {
	//get id off url
	id := c.Param(("id"))

	//delete post
	initializers.DB.Delete(&models.Post{}, id)

	//return
	c.JSON(200, gin.H{
		"message": "post has been deleted",
	})
}
