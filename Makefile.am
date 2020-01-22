
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

EXTRA_PROGRAMS = builder
builder_SOURCES = builder.go

builder$(EXEEXT): $(builder_SOURCES)
	$(GO) build -o builder$(EXEEXT) $(builder_SOURCES)

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
	tools/file.go \
	bundle/public.go

bundle/public.go: builder$(EXEEXT) public/index.html
	./builder$(EXEEXT) --package=bundle -o bundle/public.go public/

syssrv$(EXEEXT): $(syssrv_SOURCES) $(EXTRA_syssrv_SOURCES) builder$(EXEEXT)
	$(GO) build -o syssrv$(EXEEXT) $(syssrv_SOURCES)

syspwd$(EXEEXT): $(syspwd_SOURCES)
	$(GO) build -o syspwd$(EXEEXT) $(syspwd_SOURCES)

EXTRA_DIST = README.md public/ \
	go.mod \
	go.sum \
	\
	front/.gitignore \
	front/babel.config.js \
	front/package-lock.json \
	front/package.json \
	front/postcss.config.js \
	front/public/favicon.ico \
	front/public/logo192.png \
	front/public/logo512.png \
	front/public/manifest.json \
	front/public/robots.txt \
	front/README.md \
	\
	front/src/_app-after.scss \
	front/src/_app-before.scss \
	front/src/app.scss \
	front/src/Clock.jsx \
	front/src/Empty.jsx \
	front/src/index.html \
	front/src/Layout.jsx \
	front/src/Login.jsx \
	front/src/main.jsx \
	front/src/Menu.jsx \
	front/src/NotFound.jsx \
	front/src/Pager.jsx \
	front/src/UserCreate.jsx \
	front/src/UserDelete.jsx \
	front/src/Users.jsx \
	front/src/UserUpdate.jsx \
	\
	front/src/Vms.jsx \
	front/src/VmStart.jsx \
	front/src/VmStop.jsx \
	\
	front/webpack.config.devel.js \
	front/webpack.config.prod.js

clean-local:
	$(RM) builder$(EXEEXT)
	$(RM) -r autom4te.cache

install-data-local:
	test -z $(DESTDIR)$(APP_LIBDIR) || $(MKDIR_P) $(DESTDIR)$(APP_LIBDIR)
	cp -r public $(DESTDIR)$(APP_LIBDIR)
	test -z $(DESTDIR)$(APP_CONFDIR) || $(MKDIR_P) $(DESTDIR)$(APP_CONFDIR)
	test -z $(DESTDIR)$(APP_LOGDIR) || $(MKDIR_P) $(DESTDIR)$(APP_LOGDIR)
	test -z $(DESTDIR)$(APP_RUNDIR) || $(MKDIR_P) $(DESTDIR)$(APP_RUNDIR)
#	test -z $(DESTDIR)$(APP_CACHEDIR) || $(MKDIR_P) $(DESTDIR)$(APP_CACHEDIR)
#	test -z $(DESTDIR)$(APP_DATABASEDIR) || $(MKDIR_P) $(DESTDIR)$(APP_DATABASEDIR)
#EOF
