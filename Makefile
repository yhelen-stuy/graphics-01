all:
	go build picmaker.go

run:
	./picmaker
	convert img.ppm img.png
	display img.png

clean:
	rm picmaker
	rm img.ppm img.png
