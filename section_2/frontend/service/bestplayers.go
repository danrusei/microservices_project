package service

import "sort"

//some business logic which find the best player for each position from a team
func topteamplayers(teamplayers []*Player) []*Player {
	var topplayers []*Player
	var interception int32
	var tackles int32
	var assists int32
	var goals int32
	var passes int32

	bestdefender := make(map[int32]string)
	bestmidfilder := make(map[int32]string)
	bestforward := make(map[int32]string)

	if len(teamplayers) > 0 {

		for i := range teamplayers {
			switch teamplayers[i].Position {
			case "Defender":
				interception = teamplayers[i].Interceptions
				tackles = teamplayers[i].Tackles
				defskills := (interception*50)/100 + (tackles*50)/100
				bestdefender[defskills] = teamplayers[i].Name

			case "Forward":
				goals = teamplayers[i].Goals
				assists = teamplayers[i].Assists
				forskills := (goals*70)/100 + (assists*30)/100
				bestforward[forskills] = teamplayers[i].Name

			case "Midfielder":
				passes = teamplayers[i].Passes
				assists = teamplayers[i].Assists
				midskills := (passes*70)/100 + (assists*30)/100
				bestmidfilder[midskills] = teamplayers[i].Name

			}

		}

		nameDef := findplayername(bestdefender)
		nameFor := findplayername(bestforward)
		nameMid := findplayername(bestmidfilder)

		for i := range teamplayers {
			if teamplayers[i].Name == nameDef || teamplayers[i].Name == nameFor || teamplayers[i].Name == nameMid {
				topplayers = append(topplayers, teamplayers[i])
			}
		}

		return topplayers
	}

	return nil
}

func topposplayers(players []*Player, position string) ([]*Player, error) {
	var err error
	var topplayers []*Player
	var interception int32
	var tackles int32
	var assists int32
	var goals int32
	var passes int32

	bestposplayer := make(map[int32]string)

	if len(players) > 0 {

		for i := range players {

			switch position {
			case "Defender":
				interception = players[i].Interceptions
				tackles = players[i].Tackles
				defskills := (interception*50)/100 + (tackles*50)/100
				bestposplayer[defskills] = players[i].Name
			case "Forward":
				goals = players[i].Goals
				assists = players[i].Assists
				forskills := (goals*70)/100 + (assists*30)/100
				bestposplayer[forskills] = players[i].Name
			case "Midfielder":
				passes = players[i].Passes
				assists = players[i].Assists
				midskills := (passes*70)/100 + (assists*30)/100
				bestposplayer[midskills] = players[i].Name
			default:
				return nil, err
			}
		}

		namePos := findplayersname(bestposplayer)

		for i := range players {
			for v := range namePos {
				if players[i].Name == namePos[v] {
					topplayers = append(topplayers, players[i])
				}
			}
		}

		return topplayers, nil

	}

	return nil, nil

}

func findplayersname(bestposition map[int32]string) []string {
	keysDef := make([]int, 0, len(bestposition))
	for k := range bestposition {
		keysDef = append(keysDef, int(k))
	}
	sort.Ints(keysDef)

	namePos := make([]string, 3)

	namePos[0] = bestposition[int32(keysDef[len(keysDef)-1])]
	namePos[1] = bestposition[int32(keysDef[len(keysDef)-2])]
	namePos[2] = bestposition[int32(keysDef[len(keysDef)-3])]

	return namePos

}

func findplayername(bestposition map[int32]string) string {
	keysDef := make([]int, 0, len(bestposition))
	for k := range bestposition {
		keysDef = append(keysDef, int(k))
	}
	sort.Ints(keysDef)
	namePos := bestposition[int32(keysDef[len(keysDef)-1])]

	return namePos
}
