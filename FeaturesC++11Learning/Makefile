CXX   = g++
FLAGS = -O3 -std=c++11 -static

all:
	@for file in *.cpp; do \
		${CXX} ${FLAGS} -o `echo $$file | cut -d. -f1`.test $$file; \
		strip `echo $$file | cut -d. -f1`.test; \
	done
