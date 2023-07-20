install:
	bundle install

preview:
	bundle exec jekyll serve
	
hot-reload:
	jekyll serve -l -o

build:
	bundle exec jekyll build