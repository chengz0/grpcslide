from concurrent import futures

import grpc

import time

import proxy_pb2
import proxy_pb2_grpc

class WeedProxy(proxy_pb2_grpc.WeedProxyServicer):
	
	def GetFileContent(self, request, context):
		return proxy_pb2.WeedFileResponse(FileContent='hello, %s!' % request.Fid)

def serve():
	server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
	proxy_pb2_grpc.add_WeedProxyServicer_to_server(WeedProxy(), server)
	server.add_insecure_port('[::]:50051')
	server.start()
	try:
		while True:
		 time.sleep(100)
	except KeyboardInterrupt:
	 server.stop(0)

if __name__ == '__main__':
	serve()