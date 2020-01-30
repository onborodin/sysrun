
AUTOMAKE_OPTIONS = foreign no-dependencies no-installinfo

SUFFIXES = .go .c
OBJEXT= none

if LINUX_OS
if SYSTEMD
systemddir = /lib/systemd/system
nobase_systemd_DATA = syssrv.service
else
etcdir = @sysconfdir@
nobase_etc_SCRIPTS = init.d/syssrv
endif
endif



sbin_PROGRAMS = syssrv syspwd

syssrv_SOURCES = syssrv.go
syspwd_SOURCES = syspwd.go

EXTRA_syspwd_SOURCES = \
	config/config.go \
	server/user-model/user_model.go \
	server/vm-controller/vm_controller.go \
	server/vm-model/vm_model.go

EXTRA_syssrv_SOURCES = \
	config/config.go \
	daemon/daemon.go \
	server/server.go \
	server/user-controller/user_controller.go \
	server/user-model/user_model.go \
	server/vm-controller/vm_controller.go \
	server/vm-model/vm_model.go \
	tools/file.go

syssrv$(EXEEXT): $(syssrv_SOURCES) $(EXTRA_syssrv_SOURCES) builder$(EXEEXT)
	$(GO) build -o syssrv$(EXEEXT) $(syssrv_SOURCES)

syspwd$(EXEEXT): $(syspwd_SOURCES) $(EXTRA_syspwd_SOURCES)
	$(GO) build -o syspwd$(EXEEXT) $(syspwd_SOURCES)

EXTRA_PROGRAMS = builder
builder_SOURCES = builder.go

builder$(EXEEXT): $(builder_SOURCES)
	$(GO) build -o builder$(EXEEXT) $(builder_SOURCES)

bundle/public.go: builder$(EXEEXT) public/index.html
	test -z bundle || $(MKDIR_P) bundle
	./builder$(EXEEXT) --package=bundle -o bundle/public.go public/

EXTRA_syssrv_SOURCES += \
	bundle/public.go

EXTRA_DIST = README.md public/ \
	go.mod \
	go.sum \
	\
	front/src/ \
	front/public/ \
	\
	front/.gitignore \
	front/babel.config.js \
	front/package-lock.json \
	front/package.json \
	front/postcss.config.js \
	front/README.md \
	front/webpack.config.devel.js \
	front/webpack.config.prod.js

clean-local:
	rm -f builder$(EXEEXT)
	rm -rf autom4te.cache

install-data-local:
	test -z $(DESTDIR)$(APP_LIBDIR) || $(MKDIR_P) $(DESTDIR)$(APP_LIBDIR)
#	cp -r public $(DESTDIR)$(APP_LIBDIR)
	test -z $(DESTDIR)$(APP_CONFDIR) || $(MKDIR_P) $(DESTDIR)$(APP_CONFDIR)
	test -z $(DESTDIR)$(APP_LOGDIR) || $(MKDIR_P) $(DESTDIR)$(APP_LOGDIR)
	test -z $(DESTDIR)$(APP_RUNDIR) || $(MKDIR_P) $(DESTDIR)$(APP_RUNDIR)
#EOF
