GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = random-user-agent

build:
	env GOPATH=$(GOPATH) go install $(PROGRAMS)

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/bin

strip: build
	strip --strip-all $(BINDIR)/random-user-agent

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/random-user-agent $(DESTDIR)/usr/bin

clean:
	/bin/rm -f bin/random-user-agent

distclean: clean

uninstall:
	/bin/rm -f $(DESTDIR)/usr/bin

all: depend build strip install

