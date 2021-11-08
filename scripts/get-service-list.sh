#!/bin/sh

GRPC_HOST="localhost:8082"
GRPC_METHOD="ozonmp.rcn_service_api.v1.RcnServiceApiService/ListServiceV1"

payload=$(
  cat <<EOF
{
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}