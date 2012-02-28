all:
	for i in $$(find ./plugins/ | grep '\.js$$' | grep -v '\.min\.js$$'); do \
		min=$$(echo $$i | sed s/\.js$$/.min.js/); \
		java -jar ./bin/yuicompressor --type js --charset utf8 $$i > $$min; \
	done;
	
	for i in $$(find ./jquery/ | grep '\.js$$' | grep -v '\.min\.js$$'); do \
		min=$$(echo $$i | sed s/\.js$$/.min.js/); \
		java -jar ./bin/yuicompressor --type js --charset utf8 $$i > $$min; \
	done;
	
	java -jar ./bin/yuicompressor --type js --charset utf8 jquery.foo.js > jquery.foo.min.js;
	
	ln -sf ../jquery.foo.min.js public/jquery.foo.js;

	ln -sf ../jquery/$$(cat jquery/LATEST | sed s/\.js$$/\.min\.js/g) public/jquery.js;
