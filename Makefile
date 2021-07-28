PROTOC := protoc -I. -Ivendor --go_out=paths=source_relative:. --plugin protoc-gen-go="${GOPATH}/bin/protoc-gen-go" --go-vtproto_out=paths=source_relative:. --plugin protoc-gen-go-vtproto="${GOPATH}/bin/protoc-gen-go-vtproto" --go-vtproto_opt=pool=github.com/molotovtv/benchvitess.Obj1 --go-vtproto_opt=pool=github.com/molotovtv/benchvitess.Obj2
PROTO_DIRS = $(shell find . -type f -name '*.proto' -not -path "./vendor/*" -print0 | xargs -0 -n1 dirname | sort --unique)
PROTO_FILES = $(shell find . -type f -name '*.proto' -not -path "./vendor/*")
MODROOT := github.com/molotovtv/benchvitess

ALL: proto

proto: $(PROTO_FILES)
	rm -rf vendor
	#go get github.com/srikrsna/protoc-gen-gotag
	go mod vendor
	mkdir -p vendor/$(MODROOT)
	for PROTO_DIR in $(PROTO_DIRS); do rsync -R -a --prune-empty-dirs --include '*/' --include '*.proto' --exclude '*' $$PROTO_DIR/*.proto vendor/$(MODROOT)/; done;
	for PROTO_FILE in $(PROTO_FILES); do $(PROTOC) $$PROTO_FILE; done;
	rm -rf vendor