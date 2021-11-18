package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type uRequest struct {
	Id       int    `json:"id"`
	Req      string `json:"req"`
	Dateunix int64  `json:"date"`
}

var AllRequests []uRequest

const placeToSaveData string = "../data.json"

var g string
var allChar = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", ".", "/", "?", "!", "@", "%", "(", ")", " ", "ś", "Ś", "Ą", "Ż", "ż", "ą", "ź", "Ź", "ł", "Ł", "ę", "Ę", "Ó", "ó", "ć", "Ć"}

var jump int

func caesarcipher(g string) (string, int) {
	jump = rand.Intn(len(allChar) - 1)
	if jump == 0 {
		jump = 1
	}
	if !(0 < jump && jump < len(allChar)) {
		return "Error", 0
	}
	// commetned works with ascii but doesent include polish char
	// var workingHashedPhrase []string
	// byteArray := []byte(g)
	// for _, v := range byteArray {
	// 	workingHashedPhrase = append(workingHashedPhrase, string(v+byte(jump)))
	// }

	phrase := strings.Split(g, "")
	var workingHashedPhrase []string
	counter := 0
	for i := 0; i < len(phrase); i++ {
		counter = 0
		for j := 0; j < len(allChar); j++ {
			//helper is used to calculate if table would run out of range
			var helper = jump + j
			if phrase[i] == allChar[j] {
				switch len(allChar) > helper {
				case false:
					workingHashedPhrase = append(workingHashedPhrase, allChar[helper-len(allChar)])
				default:
					workingHashedPhrase = append(workingHashedPhrase, allChar[helper])
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

func addRequestToJson(nameOfReq string, reqs []uRequest) {
	// check if list isnt out of range and calculate id
	var whichId int
	if len(reqs) > 0 {
		whichId = reqs[len(reqs)-1].Id + 1
	} else {
		whichId = 0
	}

	AllRequests = append(AllRequests, uRequest{
		Id:       whichId,
		Req:      nameOfReq,
		Dateunix: time.Now().Unix(),
	})

	file, _ := json.MarshalIndent(AllRequests, "", " ")
	_ = ioutil.WriteFile(placeToSaveData, file, 0644)
}

func main() {
	//read previous messages
	file, err := ioutil.ReadFile(placeToSaveData)
	if err != nil {
		fmt.Println(err)
	}
	var holdDataFromFile []uRequest
	json.Unmarshal(file, &holdDataFromFile)
	AllRequests = append(AllRequests, holdDataFromFile...)
	//start serving
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
			addRequestToJson(g, AllRequests)
			c.JSON(200, gin.H{"msgCaesar": list, "jump": jump, "md5": md5.Sum([]byte(g))})
		} else {
			fmt.Println("Error")
		}
	})

	r.Run(":8080")
}
