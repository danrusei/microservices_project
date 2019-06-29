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

/*
func topposplayers(players []*Player, position string) ([]*Player, error) {
	var err error
	var skill1, skill2 string

	switch position {
	case "Defender":
		skill1 = "Interceptions"
		skill2 = "Tackles"
		top3defenders := toppositional(players, skill1, skill2)
		return top3defenders, nil
	case "Forward":
		skill1 = "Goals"
		skill2 = "Assists"
		top3forwarders := toppositional(players, skill1, skill2)
		return top3forwarders, nil
	case "Midfielder":
		skill1 = "Passes"
		skill2 = "Assists"
		top3midfielders := toppositional(players, skill1, skill2)
		return top3midfielders, nil
	default:
		return nil, err
	}

	return nil, nil

}

func toppositional(posplayers []*Player, skill1 string, skill2 string) []*Player {
	var topplayers []*Player

	allposplayers := make(map[int32]string)

	if len(posplayers) > 0 {
		for i := range posplayers {
			gotskill1 := posplayers[i].+"skill1"
			gotskill2 := posplayers[i].+"skill2"
		}

	}

	return nil

}
*/

func findplayername(bestposition map[int32]string) string {
	keysDef := make([]int, 0, len(bestposition))
	for k := range bestposition {
		keysDef = append(keysDef, int(k))
	}
	sort.Ints(keysDef)
	namePos := bestposition[int32(keysDef[len(keysDef)-1])]

	return namePos
}
