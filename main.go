package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/julianshen/text2img"
)

func randomString() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		content := c.PostForm("content")
		c.JSON(http.StatusOK, gin.H{"content": content})
	})

	path := "src/fonts/Koruri.ttf"
  d, err := text2img.NewDrawer(text2img.Params{
    FontPath: path,
  })
	fmt.Print(err)

  img, err := d.Draw("text2img generates the image from a text")
  fmt.Print(err)

  file, err := os.Create("dist/"+ randomString() + ".jpg")
  fmt.Print(err)
  defer file.Close()

  err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
  fmt.Print(err)

	r.Run(":8080")
}
