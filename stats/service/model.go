package service

//Table model

type Table struct {
	TeamName    string
	TeamPlayed  int32
	TeamWon     int32
	TeamDrawn   int32
	TeamLost    int32
	TeamGF      int32
	TeamGA      int32
	TeamGD      int32
	TeamPoints  int32
	TeamCapital int32
}

//Player model

type Player struct {
	Name          string
	Team          string
	Nationality   string
	Position      string
	Appearences   int32
	Goals         int32
	Assists       int32
	Passes        int32
	Interceptions int32
	Tackles       int32
	Fouls         int32
}
