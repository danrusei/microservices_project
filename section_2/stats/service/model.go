package service

//Table struct holds League table data
type Table struct {
	TeamName    string `firestore:"Name"`
	TeamPlayed  int32  `firestore:"Played"`
	TeamWon     int32  `firestore:"Won"`
	TeamDrawn   int32  `firestore:"Drawn"`
	TeamLost    int32  `firestore:"Lost"`
	TeamGF      int32  `firestore:"GF"`
	TeamGA      int32  `firestore:"GA"`
	TeamGD      int32  `firestore:"GD"`
	TeamPoints  int32  `firestore:"Points"`
	TeamCapital int32  `firestore:"Capital"`
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
	//	Price         int32  `firestore:"Price" json:"price"`
}
