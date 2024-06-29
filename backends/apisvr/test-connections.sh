set -x

curl -v -L --header "Content-Type: application/json" --data '{}' https://${SERVER_DOMAIN}/task.v1.TaskService/List

buf curl --schema . --protocol connect https://${SERVER_DOMAIN}/task.v1.TaskService/List
buf curl --schema . --protocol grpcweb https://${SERVER_DOMAIN}/task.v1.TaskService/List
buf curl --schema . --protocol grpc https://${SERVER_DOMAIN}/task.v1.TaskService/List

grpcurl -protoset proto/protoset.bin ${SERVER_DOMAIN}:443 task.v1.TaskService/List
