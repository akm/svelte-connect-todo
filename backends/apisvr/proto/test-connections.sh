set -x

# 環境変数 APISVR_ORIGIN からドメイン名とポートを抽出し、
# ポートが指定されていない場合はスキームに基づいてデフォルトポートを設定する

# APISVR_ORIGIN が設定されていない場合はエラーメッセージを表示して終了
if [ -z "$APISVR_ORIGIN" ]; then
  echo "Error: APISVR_ORIGIN is not set."
  exit 1
fi

# スキームの抽出
SCHEME=$(echo $APISVR_ORIGIN | grep :// | sed -e's,^\(.*://\).*,\1,g')

# スキームに基づいてデフォルトポートを設定
if [ "$SCHEME" = "https://" ]; then
  DEFAULT_PORT=443
elif [ "$SCHEME" = "http://" ]; then
  DEFAULT_PORT=80
else
  echo "Error: Unknown scheme in APISVR_ORIGIN."
  exit 1
fi

# ドメイン名とポートの抽出
APISVR_DOMAIN_AND_PORT=${APISVR_ORIGIN#*://}
APISVR_DOMAIN=${APISVR_DOMAIN_AND_PORT%:*}
APISVR_PORT=${APISVR_DOMAIN_AND_PORT##*:}

# ポートがドメイン名と同じ、またはポートが指定されていない場合、デフォルトポートを使用
if [ "$APISVR_PORT" = "$APISVR_DOMAIN" ] || [ -z "$APISVR_PORT" ]; then
  APISVR_PORT=$DEFAULT_PORT
fi

curl -v -L --header "Content-Type: application/json" --data '{}' ${APISVR_ORIGIN}/task.v1.TaskService/List

curl -v -L --header "Content-Type: application/json" --data '{"name":"task1", "status":"TODO"}' ${APISVR_ORIGIN}/task.v1.TaskService/Create

buf curl --schema . --protocol connect ${APISVR_ORIGIN}/task.v1.TaskService/List
buf curl --schema . --protocol grpcweb ${APISVR_ORIGIN}/task.v1.TaskService/List

if [ "$SCHEME" != "https" ]; then
  set +x
  echo "SKIP accessing with grpc protocol because the SCHEME is not https"
  exit 0
fi

buf curl --schema . --protocol grpc ${APISVR_ORIGIN}/task.v1.TaskService/List

grpcurl -protoset proto/protoset.bin ${APISVR_DOMAIN}:${APISVR_PORT} task.v1.TaskService/List
