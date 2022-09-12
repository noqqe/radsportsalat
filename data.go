package main

var teams = []string{
	"AG2R",
	"Citroën",
	"Astana",
	"Qazaqstan",
	"Bahrain",
	"Victorious",
	"BORA",
	"hansgrohe",
	"Cofidis",
	"EF",
	"Education",
	"EasyPost",
	"Groupama FDJ",
	"INEOS",
	"Grenadiers",
	"Intermarché",
	"Wanty",
	"Gobert",
	"Israel",
	"Premier Tech",
	"Jumbo",
	"Visma",
	"Lotto",
	"Soudal",
	"Movistar",
	"QuickStep",
	"Alpha",
	"Vinyl",
	"BikeExchange",
	"Jayco",
	"DSM",
	"Trek",
	"Segafredo",
	"UAE",
	"Emirates",
	"Alpecin",
	"Deceuninck",
	"B&B",
	"Hotels",
	"KTM",
	"CSF",
	"Burgos",
	"BH",
	"Caja Rural",
	"Drone",
	"Hopper",
	"EOLO",
	"Kometa",
	"Equipo",
	"Kern",
	"Pharma",
	"Euskaltel",
	"Euskadi",
	"Human",
	"Powered",
	"Health",
	"Sport",
	"Vlaanderen",
	"Baloise",
	"Arkéa",
	"Samsic",
	"TotalEnergies",
	"Uno-X",
	"Tirol",
	"KTM",
	"Union",
	"Raiffeisen",
	"KTM",
	"Bahrain",
	"Alpecin",
	"Deceuninck",
	"SwiftCarbon",
	"Premier",
	"China",
	"Glory",
	"Giant",
	"Lottery",
	"International",
	"Electro",
	"Sparta",
	"Coloquick",
	"Movistar",
	"FDJ",
	"Sauerland",
	"Dauner",
	"Lotto",
	"Kern Haus",
	"Dynatek",
	"Shimano",
	"Bridgestone",
	"Astana",
	"Qazaqstan",
	"Kaunas",
	"DSM",
	"Coop",
	"Uno-X",
	"DARE",
	"Tudor",
	"EF Education",
	"NIPPO",
}

var racePrefixes = []string{
	"Tour",
	"Race",
	"Classic",
	"Tour de",
	"Il",
	"La",
	"Criterium du",
	"Grand Prix",
	"Giro",
	"Dwars door",
	"Cicliste de",
}

var raceSuffixes = []string{
	"Tour",
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
	"Maryland",
	"Flanders",
	"Espana",
	"Toscana",
	"World Championship",
	"Deutschland",
	"Alsace",
	"Slovenia",
	"Brabant",
	"Brussels",
	"Camembert",
	"Rund um Köln",
	"Norway",
	"Hellas",
}

var riderNames = []string{
	"Primoz Roglic",
	"Jonas Vingegaard",
	"Remco Evenepoel",
	"Aleksandr Vlasov",
	"Stefan Küng",
	"Adam Yates",
	"Alexander Kristoff",
	"Pello Bilbao",
	"Sergio Higuita",
	"Biniam Girmay",
	"Alejandro Valverde",
	"Jasper Philipsen",
	"Arnaud Démare",
	"Daniel Martinez",
	"Christophe Laporte",
	"Fabio Jakobsen",
	"Julian Alaphilippe",
	"Nairo Quintana",
	"Jai Hindley",
	"Richard Carapaz",
	"Tim Merlier",
	"Geraint Thomas",
	"Hugo Hofstetter",
	"Joao Almeida",
	"Benoît Cosnefroy",
	"Matej Mohoric",
	"Simon Yates",
	"Ethan Hayter",
	"Jasper Stuyven",
	"Mads Pedersen",
	"David Gaudu",
	"Guillaume Martin",
	"Michael Matthews",
	"Neilson Powless",
	"Jack Haig",
	"Tiesj Benoot",
	"Dylan Teuns",
	"Valentin Madouas",
	"Warren Barguil",
	"Enric Mas",
	"Olav Kooij",
	"Matteo Trentin",
	"Romain Bardet",
	"Juan Ayuso",
	"Ben O'Connor",
	"Carlos Rodriguez",
	"Damiano Caruso",
	"Michael Woods",
	"Quinten Hermans",
	"Mikel Landa",
	"Marc Hirschi",
	"Diego Ulissi",
	"Giacomo Nizzolo",
	"Brandon McNulty",
	"Michael Valgren",
	"Florian Sénéchal",
	"Tom Pidcock",
	"Gino Mäder",
	"Thymen Arensman",
	"Benjamin Thomas",
	"Mark Cavendish",
	"Alessandro Covi",
	"Dylan Groenewegen",
	"Filippo Ganna",
	"Victor Campenaerts",
	"Jesus Herrada",
	"Jan Hirt",
	"Magnus Cort",
	"Alex Aranburu",
	"Louis Meintjes",
	"Fausto Masnada",
	"Ion Izagirre",
	"Tim Wellens",
	"Bauke Mollema",
	"Pavel Sivakov",
	"Julien Simon",
	"Vincenzo Nibali",
	"Max Walscheid",
	"Matis Louvel",
	"Nacer Bouhanni",
	"Anthony Turgis",
	"Alberto Dainese",
	"Jakob Fuglsang",
	"Alexey Lutsenko",
	"Magnus Sheffield",
	"Ruben Guerreiro",
	"Samuele Battistella",
	"Kasper Asgreen",
	"Sonny Colbrelli",
	"Simon Clarke",
	"Vincenzo Albanese",
	"Michael Storer",
	"Thibaut Pinot",
	"Gerben Thijssen",
	"Simone Consonni",
	"Rasmus Tiller",
	"Piet Allegaert",
	"Fernando Gaviria",
	"Elia Viviani",
	"Mauro Schmid",
	"Sam Bennett",
	"Simon Geschke",
	"Mathieu Burgaudeau",
	"Michal Kwiatkowski",
	"Pascal Ackermann",
	"Jhonnatan Narvaez",
	"Attila Valter",
	"Sepp Kuss",
	"Oliver Naesen",
	"Domenico Pozzovivo",
	"Clément Champoussin",
	"Alexis Vuillermoz",
	"Bob Jungels",
	"Marco Haller",
	"Felix Großschartner",
	"Pierre Latour",
	"Nils Politt",
	"Rafal Majka",
	"Hugo Houle",
	"Gianni Vermeersch",
	"Santiago Buitrago",
	"Mauri Vansevenant",
	"Emanuel Buchmann",
	"Luca Mozzato",
	"Patrick Konrad",
	"Koen Bouwman",
	"Peter Sagan",
	"Lucas Plapp",
	"Yves Lampaert",
	"Caleb Ewan",
	"Ben Turner",
	"Filippo Zana",
	"Esteban Chaves",
	"Bryan Coquard",
	"Lennard Kämna",
	"Giulio Ciccone",
	"Philippe Gilbert",
	"Jan Tratnik",
	"Sebastien Reichenbach",
	"Andreas Kron",
	"Georg Zimmermann",
	"Tom Devriendt",
	"Franck Bonnamour",
	"Stan Dewulf",
	"Zdenek Stybar",
	"Ben Tulett",
	"Steven Kruijswijk",
	"Rohan Dennis",
	"Connor Swift",
	"Jay Vine",
	"Edward Dunbar",
	"Dorian Godon",
	"Jan Polanc",
	"Victor Lafay",
	"Rudy Molard",
	"Wout Poels",
	"Edvald Boasson-Hagen",
	"Cian Uijtdebroeks",
	"Egan Bernal",
	"Rein Taaramäe",
	"Florian Vermeersch",
	"Geoffrey Bouchard",
	"Nikias Arndt",
	"Daniel McLay",
	"Niccolo Bonifazio",
	"Tom Sexton",
	"Sam Watson",
	"Timothy Dupont",
	"Emiliano Contreras",
	"Aurélien Paret-Peintre",
	"Davide Formolo",
	"Rune Herregodts",
	"Johannes Staune-Mittet",
	"Matteo Jorgenson",
	"Stefan Bissegger",
	"Max Kanter",
	"Lorenzo Fortunato",
	"Fred Wright",
	"Kaden Groves",
	"Sep Vanmarcke",
	"Stephen Williams",
	"Alberto Bettiol",
	"George Bennett",
	"Edward Theuns",
	"Hugh Carthy",
	"Sam Welsford",
	"Baptiste Planckaert",
	"Harm Vanhoucke",
	"Wilco Kelderman",
	"Tadej Pogacar",
	"Wout van Aert",
	"Arnaud De Lie",
	"Mathieu van der Poel",
	"Dylan van Baarle",
	"Danny van Poppel",
	"Van Avermaet Greg",
	"Dries De Bondt",
	"Kragh Andersen Søren",
	"Juan Pedro Lopez",
	"Taco van der Hoorn",
	"Ángel López Miguel",
	"Alessandro De Marchi",
	"Mick Van Dijke",
	"Richie Porte",
	"Christian Eiking Odd",
	"Tao Geoghegan Hart",
	"Jasper De Buyst",
	"Maximilian Schachmann",
	"John Degenkolb",
	"Marco Brenner",
	"Jonas Koch",
	"Ben Zwiehoff",
	"Rick Zabel",
	"Jannik Steimle",
	"Maurice Ballerstedt",
	"Henri Uhlig",
	"Pirmin Benz",
	"Nico Denz",
	"Felix Groß",
	"Jonas Rutsch",
	"Rüdiger Selig",
	"Georg Steinhauser",
	"Anton Palzer",
	"Alexander Krieger",
}

var riderTypes = []string{
	"Sprinter",
	"Bergfahrer",
	"Edelhelfer",
	"Etappenjäger",
	"GC-Fahrer",
	"Zeitfahrer",
}

var zeit = []string{
	"deinen ersten Tag",
	"deinen zweiten Tag",
	"deinen dritten Tag",
	"deinen vierten Tag",
	"deinen fünften Tag",
	"deinen letzter Tag vor Renteneintritt",
	"dein erstes Jahr",
	"dein drittes Jahr",
	"deinen zweiten Monat",
	"deinen ersten Monat",
	"deinen zwöflten Tag",
	"deinen vorletzten Tag",
	"deinen letzten Tag",
}

var beforeFeelings = []string{
	"nervös",
	"entspannt",
	"genervt",
	"gestresst",
	"zuversichtlich",
	"sicher",
	"stolz",
	"übermütig",
	"optimistisch",
	"pessimistisch",
}

var afterFeelings = []string{
	"Glücklich",
	"Erschöpft",
	"Traurig",
	"Enttäuscht",
	"Niedergeschlagen",
	"Stolz",
	"Gelangweilt",
	"Zufrieden",
	"Erleichtert",
	"Nachdenklich",
}

var teamFunctions = []string{
	"Sportlicher Leiter*in",
	"Sportlicher Leiter*in",
	"Physio",
	"Ernährungsberater*in",
	"Mechaniker*in",
	"Praktikant*in",
	"Sportdirektor*in",
	"Social-Media Berater*in",
	"Pressesprecher*in",
	"Manager*in",
}
