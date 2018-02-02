run:
	go run picmaker.go

display:
	convert img.ppm img.png
	display img.png

clean:
	rm img.ppm img.png
