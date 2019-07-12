WORK IN PROGRESS

Deploy Microservices:  
    -- create additional microservices folowing same pattern (player, transfer)
    curl --header "Content-Type: application/json" --request GET --data "@newplayer.json" http://localhost:8080/createplayer
    curl --header "Content-Type: application/json" --request GET http://localhost:8080/deleteplayer/Manchester%20City/Gabriel%20Jesus
    curl --header "Content-Type: application/json" --request GET --data '{"PlayerName":"Gabriel Jesus", "TeamFrom":"Manchester City", "TeamTO":"Chelsea"}' http://localhost:8080/transferplayer
      
    -- deploy on Kubernetes using minikube or GKE  

Blog Post : TBD 