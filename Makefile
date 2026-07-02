OUTNAME := kaiten
ENTRYPOINT := ./cmd

$(OUTNAME): $(shell find cmd internal -type f)
	go build -o $(OUTNAME) $(ENTRYPOINT)

.PHONY: clean install

clean:
	rm -rv $(OUTNAME)

install: $(OUTNAME)
	cp -rv $(OUTNAME) ~/.local/bin/$(OUTNAME)
