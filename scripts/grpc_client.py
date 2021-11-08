import asyncio

from grpclib.client import Channel

from ozonmp.rcn_service_api.v1.rcn_service_api_grpc import RcnServiceApiServiceStub
from ozonmp.rcn_service_api.v1.rcn_service_api_pb2 import DescribeServiceV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = RcnServiceApiServiceStub(channel)

        req = DescribeServiceV1Request(Sevice_id=1)
        reply = await client.DescribeSeviceV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
