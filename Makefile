build:
	protoc -I./options --go_out=./options --go_opt=paths=source_relative ./options/common.proto
	ls ./*.pb.go | xargs -n1 -IX bash -c "gsed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"