package upload

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/functions/metadata"
)

//GCSEvent is the payload of a GCS event.
type GCSEvent struct {
	Bucket         string    `json:"bucket"`
	Name           string    `json:"name"`
	Metageneration string    `json:"metageneration"`
	ResourceState  string    `json:"resourceState"`
	TimeCreated    time.Time `json:"timeCreated"`
	Updated        time.Time `json:"updated"`
}

type Team struct {
	Club   string
	Won    int32
	Drawn  int32
	Lost   int32
	GF     int32
	GA     int32
	GD     int32
	Points int32
}

type Player struct {
	Team          string
	Player        string
	Nationality   string
	Position      string
	Appearences   int32
	Goals         int32
	Assists       int32
	Passes        int32
	Interceptions int32
	Tackles       int32
	Fouls         int32
	Price         int32
}

// ToFirestore reads GCS file and upload contet to Firestore
func ToFirestore(ctx context.Context, e GCSEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Event ID: %v\n", meta.EventID)
	log.Printf("Event type: %v\n", meta.EventType)
	log.Printf("Bucket: %v\n", e.Bucket)
	log.Printf("File: %v\n", e.Name)
	log.Printf("Metageneration: %v\n", e.Metageneration)
	log.Printf("Created: %v\n", e.TimeCreated)
	log.Printf("Updated: %v\n", e.Updated)

	if e.Name == "Teams.csv" {
		teams, err := teamsTOstruct(e.Bucket, e.Name)
		if err != nil {
			panic(err)
		}
		fmt.Printf(teams[0].Club)

	} else {
		return nil
	}

	return nil
}

func teamsTOstruct(bucket string, file string) ([]Team, error) {

	// Open CSV file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	teams := []Team{}

	// Loop through lines & turn into object
	for _, line := range lines {
		var lineint []int64
		for i := 0; i <= len(line); i++ {
			switch line[i] {
			case line[0]:
				continue
			default:
				lineint[i], err = strconv.ParseInt(line[i], 10, 32)
				if err != nil {
					panic(err)
				}
			}
		}

		team := Team{
			Club:   line[0],
			Won:    int32(lineint[1]),
			Drawn:  int32(lineint[2]),
			Lost:   int32(lineint[3]),
			GF:     int32(lineint[4]),
			GA:     int32(lineint[5]),
			GD:     int32(lineint[6]),
			Points: int32(lineint[7]),
		}
		fmt.Println(team.Club + " " + string(team.Won) + " " + string(team.Drawn) + " " + string(team.Lost) + " " + string(team.GF) + " " + string(team.GA) + " " + string(team.GD) + " " + string(team.Points))
		teams = append(teams, team)
	}

	return teams, nil
}
