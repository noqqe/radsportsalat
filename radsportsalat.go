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

var racePrefixes = []string{
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
}

var raceSuffixes = []string{
	"Tour",
	"Classic",
	"Cyclassics",
	"Criterium",
	"Grand Prix",
	"Giro",
}

// TODO Proper tokenizing
var raceNames = []string{
	"Santos Down Under",
	"Cadel Evans",
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

var riderNames = []string{
	"Tadej Pogacar",
	"Wout van Aert",
	"Aleksandr Vlasov",
	"Remco Evenepoel",
	"Alejandro Valverde",
	"Jonas Vingegaard",
	"Mathieu van der Poel",
	"Pello Bilbao",
	"Primoz Roglic",
	"Jai Hindley",
	"Arnaud De Lie",
	"Richard Carapaz",
	"Daniel Martinez",
	"Nairo Quintana",
	"Alexander Kristoff",
	"Mads Pedersen",
	"Hugo Hofstetter",
	"Stefan Küng",
	"Sergio Higuita",
	"Biniam Girmay",
	"Guillaume Martin",
	"Brandon McNulty",
	"Michael Matthews",
	"Warren Barguil",
	"Damiano Caruso",
	"Ben O'Connor",
	"Benoît Cosnefroy",
	"Mikel Landa",
	"Tim Wellens",
	"Fabio Jakobsen",
	"Jasper Philipsen",
	"Arnaud Démare",
	"Matej Mohoric",
	"Michael Woods",
	"Danny van Poppel",
	"Dylan Teuns",
	"Romain Bardet",
	"Ethan Hayter",
	"Giacomo Nizzolo",
	"Jan Hirt",
	"Olav Kooij",
	"Benjamin Thomas",
	"Tim Merlier",
	"Joao Almeida",
	"Jesus Herrada",
	"Geraint Thomas",
	"Dylan van Baarle",
	"Vincenzo Albanese",
	"Tobias Halland Johannessen",
	"Simon Yates",
	"Lennard Kämna",
	"Diego Ulissi",
	"Juan Ayuso",
	"Simone Consonni",
	"Christophe Laporte",
	"Vincenzo Nibali",
	"Thymen Arensman",
	"Julien Simon",
	"Marc Hirschi",
	"Mauro Schmid",
	"Adam Yates",
	"Domenico Pozzovivo",
	"Jakob Fuglsang",
	"Ion Izagirre",
	"Amaury Capiot",
	"Santiago Buitrago",
	"David Gaudu",
	"Carlos Rodriguez",
	"Caleb Ewan",
	"Mark Cavendish",
	"Dylan Groenewegen",
	"Neilson Powless",
	"Tom Pidcock",
	"Enric Mas",
	"Axel Zingle",
	"Thibaut Pinot",
	"Quinten Hermans",
	"Luca Mozzato",
	"Simon Clarke",
	"Dries De Bondt",
	"Louis Meintjes",
	"Natnael Tesfazion",
	"Alexis Vuillermoz",
	"Steff Cras",
	"Fernando Gaviria",
	"Emanuel Buchmann",
	"Andrea Vendrame",
	"Andreas Kron",
	"Jasper Stuyven",
	"Juan Pedro Lopez",
	"Pierre Latour",
	"Giulio Ciccone",
	"Julian Alaphilippe",
	"Piet Allegaert",
	"Valentin Madouas",
	"Max Walscheid",
	"Alberto Dainese",
	"Bauke Mollema",
	"Nacer Bouhanni",
	"Jack Haig",
	// TODO complete names from https://firstcycling.com/ranking.php?k=fc&rank=&y=2022&page=1
}

var riderTypes = []string{
	"Sprinter",
	"Bergfahrer",
	"Puncheur",
	"Domestique",
	"Zeitfahrer",
}

var feelings = []string{
	"Glücklich",
	"Erschöpft",
	"Traurig",
	"Enttäuscht",
	"Niedergeschlagen",
	"Stolz",
	"Gelangweilt",
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

func genFeeling() string {

	feeling := feelings[rand.Intn(len(feelings))]
	return feeling

}

func genPosition() string {

	if rollDice(5) {
		return "DNS"
	}

	if rollDice(5) {
		return "DNF"
	}

	position := rand.Intn(180)
	return strconv.Itoa(position)

}

type Params struct {
	Name string `form:"name" binding:"required"`
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

	// Team Page
	router.GET("/team", func(c *gin.Context) {
		var params Params
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		// Seed based on Name
		h := md5.New()
		io.WriteString(h, params.Name)
		var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
		rand.Seed(int64(seed))

		c.HTML(http.StatusOK, "team.tmpl", gin.H{
			"title":     "Radsportsalat",
			"handle":    params.Name,
			"teamName":  genTeamName(),
			"raceName":  genRaceName(),
			"riderName": genRiderName(),
			"feeling":   genFeeling(),
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
