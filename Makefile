.PHONY: __init-db-args
__init-db-args:
ifndef HOST
	$(warning HOST was not set; localhost is used)
	$(eval HOST := localhost)
endif
ifndef PORT
	$(warning PORT was not set; 3306 is used)
	$(eval PORT := 3306)
endif
ifndef USERNAME
	$(warning USERNAME was not set; root is used)
	$(eval USERNAME := root)
endif
ifndef PASS
	$(warning PASS was not set; root is used)
	$(eval PASS := root)
endif
ifndef DBNAME
	$(warning DBNAME was not set; comiq_dev is used)
	$(eval DBNAME := comiq_dev)
endif

.PHONY: db-migrate
db-migrate: __init-db-args
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -source "file://ddl" -database "mysql://$(USERNAME):$(PASS)@tcp($(HOST):$(PORT))/$(DBNAME)" up
	$(MAKE) xo

.PHONY: xo
xo: __init-db-args
	go install github.com/xo/xo@42b11c7999bc6ac5be620949723f44bd0ec63e02
	xo schema --out "gen/comiq/dbschema" -t json 'mysql://$(USERNAME):$(PASS)@$(HOST):$(PORT)/$(DBNAME)'

.PHONY: dao
dao:
	go run ./script/dbgen ./gen/comiq/dbschema/xo.xo.json

.PHONY: seed
seed: __init-db-args
	for file in $$(find test/data -type f -name '*.sql'); do mysql -u$(USERNAME) -p$(PASS) -h$(HOST) -P$(PORT) --protocol='tcp' --database=$(DBNAME) < $$file; done
