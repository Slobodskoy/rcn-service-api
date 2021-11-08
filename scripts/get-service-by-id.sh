#!/bin/sh

GRPC_HOST="localhost:8082"
GRPC_METHOD="ozonmp.rcn_service_api.v1.RcnServiceApiService/DescribeServiceV1"

payload=$(
  cat <<EOF
{
  "service_id": 4
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}