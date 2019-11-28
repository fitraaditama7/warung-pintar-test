#Build
producer:
	@cd cmd/producer && go run main.go

consumer:
	@cd cmd/consumer && go run main.go
	
socket:
	@cd cmd/socket && go run main.go

.PHONY: producer consumer socket