package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var teams = []string{
	"INEOS",
	"Grenadiers",
	"Jumbo",
	"Visma",
	"Groupama",
	"FDJ",
	"UAE",
	"Emirates",
	"BORA",
	"hansgrohe",
	"Movistar",
	"Arkéa Samsic",
	"Bahrain",
	"Victorious",
	"DSM",
	"Astana",
	"Qazaqstan",
	"Cofidis",
	"Intermarché",
	"Wanty",
	"Gobert",
	"Matériaux",
	"Trek",
	"Segafredo",
	"EF Education",
	"EasyPost",
	"Israel Premier Tech",
	"AG2R",
	"Citroën",
	"B&B Hotels",
	"KTM",
	"Total",
	"Energies",
	"BikeExchange",
	"Jayco",
	"Alpecin",
	"Deceuninck",
	"Quick-Step",
	"Alpha Vinyl",
	"Lotto",
	"Soudal",
}

var raceFixes = []string{
	"Tour",
	"Race",
	"Classic",
	"Tour de",
	"Il",
	"La",
	"Cyclassics",
	"Criterium du",
	"Grand Prix",
	"Giro",
	"Dwars door",
	"Cicliste de",
	"-",
}

// TODO Proper tokenizing
var raceNames = []string{
	"Santos Down Under",
	"Cadel Evant",
	"Great Ocean Road",
	"UAE",
	"Omloop",
	"Het Nieuwsblad",
	"Strade Bianche",
	"Paris",
	"Nice",
	"Tirreno",
	"Adriatico",
	"Milano",
	"Sanremo",
	"Volta",
	"Ciclista ",
	"a Catalunya",
	"Minerva",
	"Brugge",
	"De Panne",
	"E3 Saxo Bank",
	"Gent",
	"Wevelgem",
	"Vlaanderen",
	"Ronde van",
	"Itzulia Basque Country",
	"Amstel Gold Race",
	"Paris",
	"Roubaix",
	"Flèche",
	"Wallonne",
	"Liège",
	"Bastogne",
	"Liège",
	"Romandie",
	"Eschborn",
	"Frankfurt",
	"d'Italia",
	"Dauphiné",
	"Suisse",
	"France",
	"Donostia San Sebastian Klasikoa",
	"Pologne",
	"Vuelta ciclista a España",
	"BEMER",
	"Bretagne",
	"Ouest-France",
	"Benelux",
	"Québec",
	"Montréal",
	"Lombardia",
	"Guangxi",
}

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
		pre = raceFixes[rand.Intn(len(raceFixes))] + " "
	}
	name1 := raceNames[rand.Intn(len(raceNames))]
	name2 := raceNames[rand.Intn(len(raceNames))]

	if rollDice(50) {
		post = " " + raceFixes[rand.Intn(len(raceFixes))]
	}

	// TODO RNADOMIZE Trenner
	return pre + name1 + " " + name2 + post
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "Radsportsalat",
			"teamName": genTeamName(),
			"raceName": genRaceName(),
		})
	})
	router.Run(":8080")

}
