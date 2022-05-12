build:
	go build cmd/auth/auth.go
	go build cmd/control_users/control_users.go
	go build cmd/news/news.go


#show_pid:
#	fuser -n tcp -k 8001
#	fuser -n tcp -k 8002
#	fuser -n tcp -k 8003