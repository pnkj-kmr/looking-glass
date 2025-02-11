rm_dir:
	rm -rf dist/client 

build:
	make rm_dir && npm run build && mv dist/client/* ../static/

run:
	npm start
	