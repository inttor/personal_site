import grpc
from concurrent import futures
import time

# 等待 protobuf 编译后导入
# import pb.api_pb2 as api_pb2
# import pb.api_pb2_grpc as api_pb2_grpc

class AIServiceServicer:
    def SubmitTask(self, request, context):
        print(f"Received request: {request.prompt}, Type: {request.task_type}")
        # 这里可调用 agents 或 tasks 下的具体逻辑
        # return api_pb2.TaskResponse(status=200, message="Success", result_data="...")
        pass

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # api_pb2_grpc.add_AIServiceServicer_to_server(AIServiceServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Python gRPC Server started at :50051")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
