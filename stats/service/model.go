package service

//Table struct holds League table data
type Table struct {
	TeamName    string `firestore:"Name" json:"teamName"`
	TeamPlayed  int32  `firestore:"Played" json:"teamPlayed"`
	TeamWon     int32  `firestore:"Won" json:"teamWon"`
	TeamDrawn   int32  `firestore:"Drawn" json:"teamDrawn"`
	TeamLost    int32  `firestore:"Lost" json:"teamLost"`
	TeamGF      int32  `firestore:"GF" json:"teamGF"`
	TeamGA      int32  `firestore:"GA" json:"teamGA"`
	TeamGD      int32  `firestore:"GD" json:"teamGD"`
	TeamPoints  int32  `firestore:"Points" json:"teamPoints"`
	TeamCapital int32  `firestore:"Capital" json:"teamCapital"`
}

//Player struct holds player data
type Player struct {
	Name          string `firestore:"player" json:"name"`
	Team          string `firestore:"team" json:"team"`
	Nationality   string `firestore:"nationality" json:"nationality"`
	Position      string `firestore:"position" json:"position"`
	Appearences   int32  `firestore:"appearences" json:"appearences"`
	Goals         int32  `firestore:"goals" json:"goals"`
	Assists       int32  `firestore:"assists" json:"assists"`
	Passes        int32  `firestore:"passes" json:"passes"`
	Interceptions int32  `firestore:"Interceptions" json:"interceptions"`
	Tackles       int32  `firestore:"Tackles" json:"tackles"`
	Fouls         int32  `firestore:"Fouls" json:"fouls"`
	Price         int32  `firestore:"Price" json:"price"`
}
