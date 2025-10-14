test:
	TRACE=1 go test -cover ./internal

clean:
	cd internal/testgraphs && rm -f *.dot *.png *.src
	cd internal && rm -f *.dot *.png *.src