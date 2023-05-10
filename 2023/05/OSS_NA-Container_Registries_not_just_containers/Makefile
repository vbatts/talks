
img/%.png: img/%.dot
	dot -Tpng $^ > $@

registry:
	PORT=1234 crane registry serve
