# ------------------------------------------------------------------ #
#                              protobuf                              #
# ------------------------------------------------------------------ #
.PHONY:pb/gen
pb/gen:
	@protoc --proto_path=. \
		--go_out=. --go_opt=module=github.com/ucok-man/h8-p3-ngc \
		--go-grpc_out=. --go-grpc_opt=module=github.com/ucok-man/h8-p3-ngc \
		./proto/*.proto

.PHONY:pb/rm
pb/rm:
	@rm -rf pb/*



# ------------------------------------------------------------------ #
#                              database                              #
# ------------------------------------------------------------------ #
.PHONY: db/create
db/create:
	@echo "creating and running database..."
	@docker container run -d --name mongo-server \
		--publish 27017:27017 \
		--volume mongo-server:/data/db \
		-e MONGO_INITDB_ROOT_USERNAME=root \
		-e MONGO_INITDB_ROOT_PASSWORD=root \
		mongo:6.0
	@echo "done!"

.PHONY: db/start
db/start:
	@echo "running database..."
	@docker container start mongo-server
	@echo "done"

.PHONY: db/stop
db/stop:
	@echo "stopping database..."
	@docker container stop mongo-server
	@echo "done!"

.PHONY: db/remove
db/remove:
	@echo "removing database..."
	@docker container remove -f -v mongo-server
	@echo "done"

.PHONY: db/cleanup
db/cleanup:
	@echo "removing database and volumes..."
	@docker container remove -f -v mongo-server
	@docker volume remove mongo-server
	@echo "done"

.PHONY: db/restart
db/restart: db/cleanup db/create