package main

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// This is a helper to make decisions
// based on percentage. Stupid but works.
func rollDice(prob int) bool {
	d := rand.Intn(100-1) + 1

	if d < prob {
		return true
	}

	return false
}

func genTeamName() string {
	// TODO RANDOMIZE Number of tokens
	pick1 := teams[rand.Intn(len(teams))]
	pick2 := teams[rand.Intn(len(teams))]

	// TODO RANDOMIZE Trenner
	return "Team " + pick1 + " " + pick2
}

func genRaceName() string {
	var pre, post string

	if rollDice(50) {
		pre = racePrefixes[rand.Intn(len(racePrefixes))] + " "
	}
	name1 := raceNames[rand.Intn(len(raceNames))]
	name2 := raceNames[rand.Intn(len(raceNames))]

	if rollDice(50) {
		post = " " + raceSuffixes[rand.Intn(len(raceSuffixes))]
	}

	// TODO RNADOMIZE Trenner
	return pre + name1 + " " + name2 + post
}

func genRiderName() string {
	name1 := riderNames[rand.Intn(len(riderNames))]
	name2 := riderNames[rand.Intn(len(riderNames))]

	firstname := strings.Split(name1, " ")[0]
	lastname := strings.Join(strings.Split(name2, " ")[1:], " ")
	return firstname + " " + lastname
}

func genRiderType() string {
	ridertype := riderTypes[rand.Intn(len(riderTypes))]
	return ridertype
}

func genTeamFunction() string {
	teamFunction := teamFunctions[rand.Intn(len(teamFunctions))]
	return teamFunction
}

func genAfterFeeling() string {
	feeling := afterFeelings[rand.Intn(len(afterFeelings))]
	return feeling
}

func genBeforeFeeling() string {
	feeling := beforeFeelings[rand.Intn(len(beforeFeelings))]
	return feeling
}

func genPosition() string {

	if rollDice(5) {
		return "DNS"
	}

	if rollDice(5) {
		return "DNF"
	}

	position := rand.Intn(150) + 1
	return strconv.Itoa(position)

}

func genZeit() string {
	zeit := zeit[rand.Intn(len(zeit))]
	return zeit
}

var Version string

type Params struct {
	Name string `uri:"name" binding:"required"`
}

type Rider struct {
	Position string
	Name     string
	Type     string
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/assets", "./assets")

	// Landing Page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Radsportsalat",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.tmpl", gin.H{
			"title":   "Radsportsalat",
			"version": Version,
		})
	})

	router.GET("/team/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"title": "Radsportsalat",
		})
	})

	// Team Page
	router.GET("/team/:name", func(c *gin.Context) {
		var params Params
		if err := c.ShouldBindUri(&params); err != nil {
			c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
				"title": "Radsportsalat",
			})
			return
		}

		// Seed based on Name
		h := md5.New()
		io.WriteString(h, params.Name)
		var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
		rand.Seed(int64(seed))

		c.HTML(http.StatusOK, "team.tmpl", gin.H{
			"title":         "Radsportsalat",
			"handle":        params.Name,
			"teamName":      genTeamName(),
			"teamFunction":  genTeamFunction(),
			"raceName":      genRaceName(),
			"riderName":     genRiderName(),
			"zeit":          genZeit(),
			"afterFeeling":  genAfterFeeling(),
			"beforeFeeling": genBeforeFeeling(),
			"team": []Rider{
				Rider{Name: genRiderName(), Position: genPosition(), Type: genRiderType()},
				Rider{Name: genRiderName(), Position: genPosition(), Type: genRiderType()},
				Rider{Name: genRiderName(), Position: genPosition(), Type: genRiderType()},
				Rider{Name: genRiderName(), Position: genPosition(), Type: genRiderType()},
				Rider{Name: genRiderName(), Position: genPosition(), Type: genRiderType()},
				Rider{Name: genRiderName(), Position: genPosition(), Type: genRiderType()},
				Rider{Name: genRiderName(), Position: genPosition(), Type: genRiderType()},
			},
		})
	})
	router.Run(":8080")

}
