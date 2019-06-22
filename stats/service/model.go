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
	Name          string `firestore:"player"`
	Team          string `firestore:"team"`
	Nationality   string `firestore:"nationality"`
	Position      string `firestore:"position"`
	Appearences   int32  `firestore:"appearences"`
	Goals         int32  `firestore:"goals"`
	Assists       int32  `firestore:"assists"`
	Passes        int32  `firestore:"passes"`
	Interceptions int32  `firestore:"Interceptions"`
	Tackles       int32  `firestore:"Tackles"`
	Fouls         int32  `firestore:"Fouls"`
	Price         int32  `firestore:"Price"`
}
