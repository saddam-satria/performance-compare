package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saddam-satria/performance-compare/go/pkg"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type Post struct {
	PostId    string    `json:"postId" gorm:"type:text;primaryKey;column:postId"`
	Title     string    `json:"title" gorm:"type:varchar(255);columm:title"`
	Body      string    `json:"body" gorm:"type:text;column:body"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime:true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime:true"`
}

type PostOnAuthor struct {
	Post
	AuthorName string `json:"authorName"`
}

type PostResponse struct {
	Total int            `json:"total"`
	Posts []PostOnAuthor `json:"posts"`
}

func MonitorUsage() gin.HandlerFunc {

	pid := os.Getpid()

	return func(c *gin.Context) {

		p, err := process.NewProcess(int32(pid))
		if err == nil {

			memInfo, _ := p.MemoryInfo()
			usedMemoryMB := float64(memInfo.RSS) / (1024 * 1024)

			cpuPercent, _ := cpu.Percent(time.Second, false)

			fmt.Printf("[%s] CPU Usage %.2f%% - Memory Usage %.2f MB ",
				time.Now().Format(time.RFC3339), cpuPercent[0], usedMemoryMB)
		}

		c.Next()
	}
}

func main() {
	r := gin.New()

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

	// r.Use(MonitorUsage())

	r.GET("/api/v1/posts", func(ctx *gin.Context) {

		query := `SELECT public."Post"."postId", "public"."Post"."title", "public"."Post"."body", "public"."Post"."createdAt", "public"."Post"."updatedAt", "public"."Author"."name" AS author_name, public."Post"."author_id" FROM "public"."Post" INNER JOIN "public"."Author" ON "public"."Author"."authorId" = "public"."Post"."author_id"`

		var posts []PostOnAuthor

		if err := pkg.Connection.Raw(query).Scan(&posts).Error; err != nil {
			panic(err.Error())
		}

		response := pkg.Response[PostResponse]("welcome to golang", http.StatusOK, PostResponse{
			Total: len(posts),
			Posts: posts,
		})

		ctx.JSON(http.StatusOK, response)
	})

	if err := pkg.Connect(&pkg.DbConfig); err != nil {
		panic("Failed To Connect To Database")
	}

	if err := r.Run(":5003"); err != nil {
		panic("Failed To Start")
	}
}
