SHELL = /bin/bash

all: clean
	@for dir in *; do \
		if [[ $${dir} =~ Learning ]]; then \
			echo $${dir} is building......; \
			$(MAKE) -C $${dir}; \
			echo; \
		fi \
	done

depend:
	$(MAKE) -C 3rdParty

clean:
	@${RM} */*.txt */*.test */*.o */*.html
