export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G = et2g
export EG2GO = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)

export ETHOSROOT=server/rootfs
export MINIMALTDROOT=server/minimaltdfs

.PHONY: all install clean
all: AccountClient AccountServer

Account.go: Account.t
	$(ETN2GO) . Account main $^

AccountServer: AccountServer.go Account.go
	ethosGo $^

AccountClient: AccountClient.go Account.go
	ethosGo $^

install: clean AccountClient AccountServer
	(ethosParams server && cd server && ethosMinimaltdBuilder)
	echo 7 > server/param/sleepTime
	ethosTypeInstall Account
	ethosDirCreate $(ETHOSROOT)/services/Account $(ETHOSROOT)/types/spec/Account/Account all
	ethosDirCreate $(ETHOSROOT)/services/Account $(ETHOSROOT)/types/spec/Account/AccountStruct all
	cp AccountServer $(ETHOSROOT)/programs
	cp AccountClient $(ETHOSROOT)/programs
#   install -D AccountServer AccountClient $(ETHOSROOT)/programs
	ethosStringEncode /programs/AccountServer > $(ETHOSROOT)/etc/init/services/AccountServer

clean:
	sudo rm -rf server
	rm -rf Account/ AccountIndex/
	rm -f Account.go
	rm -f AccountServer
	rm -f AccountServer.goo.ethos
	rm -f AccountClient
	rm -f AccountClient.goo.ethos