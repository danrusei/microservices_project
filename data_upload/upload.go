package upload

import (
	"context"
	"encoding/csv"
//	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"io"

	"cloud.google.com/go/functions/metadata"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/firestore"
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

//Team holds the team member information
type Team struct {
	Team          string `json:"team" firestore:"team"`
	Player        string `json:"player" firestore:"player"`
	Nationality   string `json:"nationality" firestore:"nationality"`
	Position      string `json:"postion" firestore:"position"`
	Appearences   int `json:"appearences" firestore:"appearences"`
	Goals         int `json:"goals" firestore:"goals"`
	Assists       int `json:"assists" firestore:"assists"`
	Passes        int `json:"passes" firestore:"passes"`
	Interceptions int `json:"interceptions" firestores:"interceptions"`
	Tackles       int `json:"tackles" firestores:"tackles"`
	Fouls         int `json:"fouls" firestores:"fouls"`
	Price         int `json:"price" firestores:"price"`
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

	//copy file from gs bucket
	//read csv file and populate Team struct
	//insert data in Firestore

	teams, err := getFileFromGCS(e.Bucket, e.Name)
	if err != nil {
		log.Printf("could not construct the struct : %v", err)
	}

	err = insertInFirestore(teams)
	if err != nil {
		log.Printf("could not create the Team Document in Firestore: %v: ", err)
	}
	//teamsJSON, _ := json.Marshal(teams)
    //fmt.Println(string(teamsJSON))
	
	return nil
}

func getFileFromGCS(bucket string, filename string) ([]Team, error) {
	ctx := context.Background()
 	client, err := storage.NewClient(ctx)
	if err != nil {
		panic("Unable to create the storage client")
	}
	bkt := client.Bucket(bucket)
	obj := bkt.Object(filename)
	r, err := obj.NewReader(ctx)
	if err != nil {
		panic("cannot read object")
	}
	defer r.Close()

	reader := csv.NewReader(r)

	var teams []Team
	// Loop through lines & turn into object
	for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
		}
		teamName := line[0]
		player := line[1]
		nationality := line[2]
		position := line[3]
		appearences, err := strconv.Atoi(line[4])
		if err != nil {
			panic(err)
		}
		goals, err := strconv.Atoi(line[5])
		if err != nil {
			panic(err)
		}
		assists, err := strconv.Atoi(line[6])
		if err != nil {
			panic(err)
		}
		passes, err := strconv.Atoi(line[7])
		if err != nil {
			panic(err)
		}
		interceptions, err := strconv.Atoi(line[8])
		if err != nil {
			panic(err)
		}
		tackles, err := strconv.Atoi(line[9])
		if err != nil {
			panic(err)
		}
		fouls, err := strconv.Atoi(line[10])
		if err != nil {
			panic(err)
		}
		price, err := strconv.Atoi(line[11])
		if err != nil {
			panic(err)
		}

		squand := Team{
			Team: teamName,
			Player: player,
			Nationality: nationality,
			Position: position,
			Appearences: appearences,
			Goals: goals,
			Assists: assists,
			Passes: passes,
			Interceptions: interceptions,
			Tackles: tackles,
			Fouls: fouls,
			Price: price,
		}

		teams = append(teams, squand)

	}	

	return teams, nil
}

func insertInFirestore(teams []Team) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "apps-microservices")
	if err != nil {
		log.Printf("cannot create new firestore client: %v", err)
	}
	teamsCol := client.Collection("Teams")
	for _, indTeam := range teams {
		ca := teamsCol.Doc(indTeam.Team + "" + indTeam.Player)
		_, err = ca.Set(ctx, Team{
			Team: indTeam.Team,
			Player: indTeam.Player,
			Nationality: indTeam.Nationality,
			Position: indTeam.Position,
			Appearences: indTeam.Appearences,
			Goals: indTeam.Goals,
			Assists: indTeam.Assists,
			Passes: indTeam.Passes,
			Interceptions: indTeam.Interceptions,
			Tackles: indTeam.Tackles,
			Fouls: indTeam.Fouls,
			Price: indTeam.Price,
		})

	}
	
	return nil

}

