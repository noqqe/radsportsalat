package main

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