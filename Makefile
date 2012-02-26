all:
	for i in $$(find ./plugins/ | grep '\.js$$' | grep -v '\.min\.js$$'); do \
		min=$$(echo $$i | sed s/\.js$$/.min.js/); \
		echo $$min; \
		java -jar ./bin/yuicompressor -v --type js --charset utf8 $$i > $$min.tmp && mv $$min.tmp $$min; \
	done;
