.PHONY: serve stop

serve:
	docker-compose up sp-mongo

stop:
	docker-compose down
