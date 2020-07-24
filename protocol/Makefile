.PHONY : all user

all: user

user:
	protoc -I . user.proto --go_out=plugins=grpc:../protogo/user/