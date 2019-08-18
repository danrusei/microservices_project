**Microservices -- Create Frontend and Stats Service :** 

* Frontend service is an http rest API which communicate with the rest of the world  
* Stats is a backend service which communicate with frontend microservices via GRPC  
* GO-KIT microservices toolkit is used

Run both Stats and Frontend services and use following commands to test functionality.

    curl http://localhost:8080/bestplayers/Tottenham%20Hotspur

    curl --header "Content-Type: application/json" --request GET --data '{"league":"League"}' http://localhost:8080/table

    curl --header "Content-Type: application/json" --request GET --data '{"position":"Defender"}' http://localhost:8080/bestposition
    position can be replaced with Forward and Midfielder  

**Blog Posts :**

* https://dev-state.com/posts/microservices_2_gokit1/  
* https://dev-state.com/posts/microservices_2_gokit2/