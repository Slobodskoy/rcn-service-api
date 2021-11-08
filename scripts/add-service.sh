#!/bin/sh

GRPC_HOST="localhost:8082"
GRPC_METHOD="ozonmp.rcn_service_api.v1.RcnServiceApiService/CreateServiceV1"

payload=$(
  cat <<EOF
{
  "value" : {
  "id": 1,
  "title": "Created Service from gRPC",
  "description": "This is the entity created via grpcurl utility",
  "rating": "5"
  }
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}