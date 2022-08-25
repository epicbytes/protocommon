build:
	protoc -I./ --go_out=./ --go_opt=paths=source_relative ./common.proto
	ls ./*.pb.go | xargs -n1 -IX bash -c "gsed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"