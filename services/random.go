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

	category := c.Query("category") // comic, scene
	device := c.Query("device")     // mobile, pc

	if category == "" {
		category = "scene"
	}
	if device == "" {
		device = "pc"
	}

	len := 0
	fileList := []string{}

	filepath.WalkDir("./images/"+category, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".jpg" && filepath.Base(d.Name())[:1] == device[:1] {
			len++
			fileList = append(fileList, d.Name())
		}
		return nil
	})

	id := GetId(len)
	imgname := fileList[id]

	c.Redirect(http.StatusMovedPermanently, baseurl+"api/images/"+category+"/"+imgname)
}
