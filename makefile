migration-script:
	@if [ "$(name)" = "" ]; then\
		printf '%s\n\n' '-- +migrate Up' '-- +migrate Down' > sql/$(shell date '+%Y%m%d%H%M%S')_new_script.sql;\
	else\
		printf '%s\n\n' '-- +migrate Up' '-- +migrate Down' > sql/$(shell date '+%Y%m%d%H%M%S')_$(name).sql;\
	fi

test:
	go test ./... -cover

mockery:
	mockery --all --inpackage

migrate-up:
	go run main.go migrate --direction=up --step=0

migrate-down-full:
	go run main.go migrate --direction=down --step=0

server:
	go run main.go server
