CXX      = gcc
CFLAGS   = -O3 -std=c99 -static
INCLUDE  = -I.

all:
	@for file in *.c; do \
		if [ $$file != "aes.c" ]; then \
			${CXX} ${CFLAGS} ${INCLUDE} -o `echo $$file | cut -d. -f1`.test aes.c $$file; \
			strip `echo $$file | cut -d. -f1`.test; \
		fi \
	done
