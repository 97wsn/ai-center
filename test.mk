test_login:
	grpcurl -plaintext -d '{"user_name": "张三", "pwd": "123"}' 127.0.0.1:3000 ai.center.rpc.v1.User/Login
