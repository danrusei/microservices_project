package service

//Table struct holds League table data
type Table struct {
	TeamName    string `json:"teamName"`
	TeamPlayed  int32  `json:"teamPlayed"`
	TeamWon     int32  `json:"teamWon"`
	TeamDrawn   int32  `json:"teamDrawn"`
	TeamLost    int32  `json:"teamLost"`
	TeamGF      int32  `json:"teamGF"`
	TeamGA      int32  `json:"teamGA"`
	TeamGD      int32  `json:"teamGD"`
	TeamPoints  int32  `json:"teamPoints"`
	TeamCapital int32  `json:"teamCapital"`
}

//Player struct holds player data
type Player struct {
	Name          string `json:"name"`
	Team          string `json:"team"`
	Nationality   string `json:"nationality"`
	Position      string `json:"position"`
	Appearences   int32  `json:"appearences"`
	Goals         int32  `json:"goals"`
	Assists       int32  `json:"assists"`
	Passes        int32  `json:"passes"`
	Interceptions int32  `json:"interceptions"`
	Tackles       int32  `json:"tackles"`
	Fouls         int32  `json:"fouls"`
	//Price         int32  `json:"price"`
}
