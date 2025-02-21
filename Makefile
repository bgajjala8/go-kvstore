PROTO_DIRS := $(wildcard proto*)
PROTO_FILES := $(foreach dir,$(PROTO_DIRS),$(wildcard $(dir)/*.proto))
OUT_DIR := .

proto:
	buf dep update
	buf generate