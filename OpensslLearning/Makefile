CXX      = g++
CFLAGS   = -O3 -std=c++11 -static
INCLUDE  = -I../temp-drivers/include
LIBS     = -L../temp-drivers/lib -lcrypto -lssl -lsctp -lz -lpthread -ldl

all:
	@for file in *.cpp; do \
		${CXX} ${CFLAGS} ${INCLUDE} -o `echo $$file | cut -d. -f1`.test $$file ${LIBS}; \
		strip `echo $$file | cut -d. -f1`.test; \
	done
