auto:
	go build
	mv goto_b ~/bin/goto_b
	mkdir -p ~/.config/goto
	cp goto* ~/.config/goto/
	echo "source ~/.config/goto/*" >> ~/.bashrc
deps:
	go get -u github.com/mattn/go-sqlite3
	
