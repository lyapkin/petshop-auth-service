include .env
export $(shell sed 's/=.*//' .env)

.PHONY: certs

certs:
	mkdir -p $(SECRET_DIR)
	openssl genrsa -out $(JWT_ACCESS_SECRET_PATH) 2048
	openssl rsa -in $(JWT_ACCESS_SECRET_PATH) -pubout -out $(JWT_ACCESS_PUBLIC_PATH)
	openssl rand -base64 32 > $(JWT_REFRESH_SECRET_PATH)
	@chmod 600 $(JWT_REFRESH_SECRET_PATH)
	@chmod 600 $(JWT_ACCESS_SECRET_PATH)
	@chmod 644 $(JWT_ACCESS_PUBLIC_PATH)
	@echo "Ключи успешно созданы в $(SECRET_DIR)"

clean-certs:
	rm -rf $(SECRET_DIR)