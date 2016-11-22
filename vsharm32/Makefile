export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)


install.rootfs = /var/lib/ethos/ethos-default-$(TARGET_ARCH)/rootfs
install.minimaltd.rootfs = /var/lib/ethos/minimaltd/rootfs


.PHONY: all install
all: Box

install: Box
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) Box1
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) Box2 
	install Box  $(install.rootfs)/programs
	install Box  $(install.minimaltd.rootfs)/programs
	echo -n /programs/Box | ethosStringEncode > $(install.rootfs)/etc/init/console
	mkdir -p $(install.rootfs)/user/nobody/myDir
	setfattr -n user.ethos.typeHash -v $(shell egPrint Box1 hash Box1) $(install.rootfs)/user/nobody/myDir
	setfattr -n user.ethos.typeHash -v $(shell egPrint Box2 hash Box2) $(install.rootfs)/user/nobody/myDir	

Box1.go: Box1.t Box2.t
	$(ETN2GO) . Box1 main $^
	$(ETN2GO) . Box2 main $^

Box: Box.go Box1.go
	ethosGo $^ 

clean:
	rm -rf Box1/ Box1Index/
	rm -f Box1.go
	rm -f Box
	rm -f Box.goo.ethos
