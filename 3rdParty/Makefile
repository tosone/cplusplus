include ../Makefile.constant

prefix_flags                  = --prefix=$(prefix)
static_flags                  = --enable-shared=no --enable-static=yes
cmake_static_flags            = BUILD_STATIC_LIBS=OFF BUILD_SHARED_LIBS=ON

OSName                        = linux

openssl_flags                 = darwin64-x86_64-cc

ifeq ($(OS),Windows_NT)
	OSName                    = Windows
else
	OSName = $(shell echo $(shell uname -s))
	ifeq ($(OSName), Linux)
		openssl_flags         = linux-x86_64
	else ifeq ($(OSName), Darwin)
		openssl_flags         = darwin64-x86_64-cc
	endif
endif

$(OSName): download depend

.PHONY: download
download:
	env GO111MODULE=off go run download.go

depend: icu boost mbedtls zlib openssl

.PHONY: boost
boost:
	cd $@ && ./bootstrap.sh --with-icu=$(prefix) --with-libraries=atomic,wave,date_time,thread,math,log $(prefix_flags) && \
	./b2 link=static threading=multi runtime-link=static $(prefix_flags) -a install

.PHONY: icu
icu:
	cd $@/icu4c/source && ./configure $(prefix_flags) $(static_flags) && \
	make -j8 && make install

.PHONY: mbedtls
mbedtls:
	cd $@ && make clean && make DESTDIR=$(prefix) install

.PHONY: zlib
zlib:
	cd $@ && if [ -f Makefile ]; then make clean; fi || \
	./configure $(prefix_flags) --static --64 && \
	make -j8 && make install

.PHONY: openssl
openssl:
	cd $@ && if [ -f Makefile ]; then make clean; fi && \
	./Configure $(openssl_flags) $(prefix_flags) no-shared no-asm no-async -I$(prefix)/include -L$(prefix)/lib zlib && \
	make -j8 && make install
