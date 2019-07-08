
This is the main service, which expose a JSON over HTTP endpoint to customers.  
It talks by GRPC with  Stats Service.

Temp:
export STATS_SERVICE_ADDR="localhost:8081"
export PLAYER_SERVICE_ADDR="localhost:8082"
export TRANSFER_SERVICE_ADDR="localhost:8083"
