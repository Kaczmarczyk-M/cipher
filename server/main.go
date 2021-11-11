package main

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var g string

// var uppercaseLetters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// var lowercaseLetters = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var allChar = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", ".", "/", "?", "!", "@", "%", "(", ")", " "}

var jump int = 1

func caesarcipher(g string) string {
	phrase := strings.Split(g, "")
	var workingHashedPhrase []string
	//check if all of phrase chars are in array
	for i := 0; i < len(phrase); i++ {
		for j := 0; j < len(allChar); j++ {
			if phrase[i] == allChar[j] {
				switch len(allChar) > j+jump {
				case false:
					workingHashedPhrase = append(workingHashedPhrase, allChar[len(allChar)-j-jump])
				default:
					workingHashedPhrase = append(workingHashedPhrase, allChar[j+jump])
				}
			}
		}
	}
	return strings.Join(workingHashedPhrase, "")
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Use caesar cipher to encode your message",
		})
	})
	api.GET("/tobehashed", func(c *gin.Context) {
		if c.FullPath() == "/api/tobehashed" {
			g = c.Request.URL.Query()["inputValue"][0]
			c.JSON(200, gin.H{"msg": caesarcipher(g)})
		} else {
			fmt.Println("Error")
		}
	})

	r.Run(":8080")
}
