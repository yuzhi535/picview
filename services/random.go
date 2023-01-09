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
	return id
}

func Random(c *gin.Context) {
	baseurl := os.Getenv("baseurl")

	len := 0
	fileList := []string{}
	filepath.Walk("./images", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".jpg" {
			len++
			fileList = append(fileList, info.Name())
		}
		return nil
	})

	id := GetId(len)
	imgname := fileList[id]
	// byteFile, err := os.ReadFile(imgname)
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"error": "read file error",
	// 	})
	// }
	// c.Header("Content-Disposition", "attachment; filename="+imgname)
	// c.Data(http.StatusOK, "application/octet-stream", byteFile)
	c.Redirect(http.StatusMovedPermanently, "http://"+baseurl+"api/images/"+imgname)
}
