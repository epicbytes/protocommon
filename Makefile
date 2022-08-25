build:
	protoc -I./common --go_out=./common --go_opt=paths=source_relative ./common/common.proto
	ls ./common/*.pb.go | xargs -n1 -IX bash -c "gsed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"