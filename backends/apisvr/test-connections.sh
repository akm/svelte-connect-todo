set -x

curl -v -L --header "Content-Type: application/json" --data '{}' https://${GCP_SERVICE_DOMAIN}/task.v1.TaskService/List

buf curl --schema . --protocol connect https://${GCP_SERVICE_DOMAIN}/task.v1.TaskService/List
buf curl --schema . --protocol grpcweb https://${GCP_SERVICE_DOMAIN}/task.v1.TaskService/List
buf curl --schema . --protocol grpc https://${GCP_SERVICE_DOMAIN}/task.v1.TaskService/List

mkdir -p tmp
buf build -o tmp/protoset.bin
grpcurl -protoset tmp/protoset.bin ${GCP_SERVICE_DOMAIN}:443 task.v1.TaskService/List
