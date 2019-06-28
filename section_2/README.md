Design Microservices:  
    -- frontend service is an http rest API which communicate with the rest of the world  
    -- stats is a backend service which communicate with frontend microservices via GRPC  
    -- I'm using GO-KIT toolkit which provides separation of concerns pattern in microservices creation

    curl http://localhost:8080/bestplayers/Tottenham%20Hotspur

    curl --header "Content-Type: application/json" --request GET --data '{"league":"League"}' http://localhost:8080/table

    curl --header "Content-Type: application/json" --request GET --data '{"position":"Defender"}' http://localhost:8080/bestposition
    position can be replaced with Forward and Midfielder  

Blog Post : TBD 