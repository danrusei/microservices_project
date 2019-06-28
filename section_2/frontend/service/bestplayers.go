package service

//
func topteamplayers(teamplayers []*Player) []*Player {
	var topplayers []*Player
	var defplayer *Player
	var forplayer *Player
	var midplayer *Player
	for i := range teamplayers {
		switch teamplayers[i].Position {
		case "Defender":
			if ((teamplayers[i].Tackles*50)/100 + (teamplayers[i].Interceptions*50)/100) >= ((defplayer.Tackles*50)/100 + (defplayer.Interceptions*50)/100) {
				defplayer = teamplayers[i]
			}
		case "Forward":
			if ((teamplayers[i].Goals*50)/100 + (teamplayers[i].Assists*50)/100) >= ((forplayer.Goals*50)/100 + (forplayer.Assists*50)/100) {
				forplayer = teamplayers[i]
			}
		case "Midfielder":
			if ((teamplayers[i].Passes*80)/100 + (teamplayers[i].Assists*20)/100) >= ((midplayer.Passes*80)/100 + (midplayer.Assists*20)/100) {
				midplayer = teamplayers[i]
			}
		}
		topplayers = append(topplayers, defplayer, forplayer, midplayer)

	}
	return topplayers
}
