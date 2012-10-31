all:
	for i in $$(find ./plugins/ | grep '\.js$$' | grep -v '\.min\.js$$'); do \
		min=$$(echo $$i | sed s/\.js$$/.min.js/); \
		java -jar ./tools/yuicompressor --type js --charset utf8 $$i > $$min; \
	done;

	for i in $$(find ./plugins/ | grep '\.css$$' | grep -v '\.min\.css$$'); do \
		min=$$(echo $$i | sed s/\.css$$/.min.css/); \
		java -jar ./tools/yuicompressor --type css --charset utf8 $$i > $$min; \
	done;

	for i in $$(find ./jquery/ | grep '\.js$$' | grep -v '\.min\.js$$'); do \
		min=$$(echo $$i | sed s/\.js$$/.min.js/); \
		java -jar ./tools/yuicompressor --type js --charset utf8 $$i > $$min; \
	done;

	java -jar ./tools/yuicompressor --type js --charset utf8 jquery.foo.js > static/jquery.foo.min.js;

	cp -a ./jquery/*.min.js static/
	cp -a ./jquery/$$(cat jquery/LATEST | sed s/\.js$$/\.min\.js/g) static/jquery.js;
