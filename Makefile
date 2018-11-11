.PHONY: run,clean

run:
	gin -a 14000 -p 13000 -i run
clean:
	-rm *.db gin-bin