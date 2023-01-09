package services

import (
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func GetId(len int) (id int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	id = r1.Intn(len)
	return
}

// consider moving fileList outside?
func Random(c *gin.Context) {
	baseurl := os.Getenv("baseurl")

	len := 0
	fileList := []string{}

	filepath.WalkDir("./images", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".jpg" {
			len++
			fileList = append(fileList, d.Name())
		}
		return nil
	})

	id := GetId(len)
	imgname := fileList[id]

	c.Redirect(http.StatusMovedPermanently, "http://"+baseurl+"api/images/"+imgname)
}
