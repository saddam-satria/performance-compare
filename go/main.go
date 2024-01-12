package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saddam-satria/performance-compare/go/pkg"
)

type Post struct{
	PostId string `json:"postId" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title string `json:"title" gorm:"type:varchar(255)"`
	Body string `json:"body" gorm:"text"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}


func (Post) TableName() string {
	return "Post"
}

func main() {
	r := gin.New()


	r.GET("/api/v1/posts", func (ctx *gin.Context) {
		var post []Post
		if err:= pkg.Connection.Find(&post).Joins("Author").Debug().Error; err != nil {
			panic("Query Error")
		}

		response := pkg.Response[[]Post]("welcome to golang", http.StatusOK, post)
		ctx.JSON(http.StatusOK, response)
	})

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] %s %s - %v\n",
				params.TimeStamp.Format(time.RFC3339),
				params.Method,
				params.Path,
				params.Latency,
			)
		},
		Output: os.Stdout,
	}))

	if err := pkg.Connect(&pkg.DbConfig); err != nil {
		panic("Failed To Connect To Database")
	}

	if err := r.Run(":5003"); err != nil {
		panic("Failed To Start")
	}
}