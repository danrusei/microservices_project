**Create additional microservices** folowing same pattern (playerOps & transfer) and test their functionality:  

    curl --header "Content-Type: application/json" --request GET --data "@newplayer.json" http://localhost:8080/createplayer  
    curl --header "Content-Type: application/json" --request GET http://localhost:8080/deleteplayer/Manchester%20City/Gabriel%20Jesus  
    curl --header "Content-Type: application/json" --request GET --data '{"PlayerName":"Gabriel Jesus", "TeamFrom":"Manchester City", "TeamTO":"Chelsea"}' http://localhost:8080/transferplayer     
      
**Deploy on Minikube (Kubernetes local cluster)**

* https://dev-state.com/posts/microservices_4_kubernetes/  

**Deploy on GKE and enable Istio for observability and service control.**

* https://dev-state.com/posts/microservices_5_istio/