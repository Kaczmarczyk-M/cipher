package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var g string
var allChar = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", ".", "/", "?", "!", "@", "%", "(", ")", " "}

var jump int

func caesarcipher(g string) (string, int) {
	jump = rand.Intn(len(allChar) - 1)
	if jump == 0 {
		jump = 1
	}
	if !(0 < jump && jump < len(allChar)) {
		return "Error", 0
	}
	phrase := strings.Split(g, "")
	var workingHashedPhrase []string
	counter := 0
	for i := 0; i < len(phrase); i++ {
		counter = 0
		for j := 0; j < len(allChar); j++ {

			if phrase[i] == allChar[j] {
				switch len(allChar) > j+jump {
				case false:
					workingHashedPhrase = append(workingHashedPhrase, allChar[jump+j-len(allChar)])
				default:
					workingHashedPhrase = append(workingHashedPhrase, allChar[j+jump])
				}
				break
			}
			counter += 1
			if counter == len(allChar) {
				return "Error - invalid characters", 0
			}
		}
	}

	return strings.Join(workingHashedPhrase, ""), jump
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Encode your message",
		})
	})
	api.GET("/tobehashed", func(c *gin.Context) {
		if c.FullPath() == "/api/tobehashed" {
			g = c.Request.URL.Query()["inputValue"][0]
			list, jump := caesarcipher(g)
			c.JSON(200, gin.H{"msgCaesar": list, "jump": jump, "md5": md5.Sum([]byte(g))})
			fmt.Println(md5.Sum([]byte(g)))
		} else {
			fmt.Println("Error")
		}
	})

	r.Run(":8080")
}
