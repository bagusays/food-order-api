migration-script:
	@if [ "$(name)" = "" ]; then\
		printf '%s\n\n' '-- +migrate Up' '-- +migrate Down' > sql/migration/$(shell date '+%Y%m%d%H%M%S')_new_script.sql;\
	else\
		printf '%s\n\n' '-- +migrate Up' '-- +migrate Down' > sql/migration/$(shell date '+%Y%m%d%H%M%S')_$(name).sql;\
	fi

mockery:
	mockery --all --inpackage

migrate-up:
	go run main.go migrate --direction=up --step=0

migrate-down:
	go run main.go migrate --direction=down --step=1

server:
	go run main.go server
