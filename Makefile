auto:
	go build
	mv goto_b ~/bin/gob
	mkdir -p ~/.config/goto
	cp goto* ~/.config/goto/
	cp gob* ~/.config/goto/
	#echo "source ~/.config/goto/*" >> ~/.bashrc
deps:
	go get -u github.com/mattn/go-sqlite3
