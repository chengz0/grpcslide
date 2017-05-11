from __future__ import print_function

import grpc

import proxy_pb2
import proxy_pb2_grpc


def run():
  channel = grpc.insecure_channel('localhost:50051')
  stub = proxy_pb2_grpc.WeedProxyStub(channel)
  response = stub.GetFileContent(proxy_pb2.WeedFileRequest(Fid='1,a1b2c3'))
  print("Weed proxy client received: " + response.FileContent)


if __name__ == '__main__':
  run()