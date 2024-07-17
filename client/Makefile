rm_dir:
	rm -rf dist/client && rm -rf ../release/static/*

build:
	make rm_dir && npm run build && cp -rf dist/client/* ../release/static/