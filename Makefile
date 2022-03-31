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

AccountType.go: AccountType.t
$(ETN2GO) . AccountType main $^

AccountServer: AccountServer.go AccountType.go
ethosGo $^

AccountClient: AccountClient.go AccountType.go
ethosGo $^

install: clean AccountClient AccountServer
(ethosParams server && cd server && ethosMinimaltdBuilder)
echo 7 > server/param/sleepTime
ethosTypeInstall AccountType
ethosDirCreate $(ETHOSROOT)/services/AccountType $(ETHOSROOT)/types/spec/AccountType/Account all
ethosDirCreate $(ETHOSROOT)/services/AccountType $(ETHOSROOT)/types/spec/AccountType/AccountStruct all
cp AccountServer $(ETHOSROOT)/programs
cp AccountClient $(ETHOSROOT)/programs
# install -D AccountServer AccountClient $(ETHOSROOT)/programs
ethosStringEncode /programs/AccountServer > $(ETHOSROOT)/etc/init/services/AccountServer

clean:
sudo rm -rf server
rm -rf AccountType/ AccountTypeIndex/
rm -f AccountType.go
rm -f AccountServer
rm -f AccountServer.goo.ethos
rm -f AccountClient
rm -f AccountClient.goo.ethos