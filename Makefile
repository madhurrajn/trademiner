TOPTARGETS := all clean

#POSTSUBDIRS := 
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u
CP=cp
BINARY_NAME=trademiner

all:  \
	$(SUBDIRS) \
	go-build \
	$(POSTSUBDIRS)

clean: go-clean $(SUBDIRS) $(POSTSUBDIRS)

go-build:
	$(GOBUILD) -o $(BINARY_NAME) -v 
	$(CP) $(BINARY_NAME) ./bin

go-clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

$(POSTSUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

deps:
	$(GOGET) github.com/tiloso/googlefinance
	$(GOGET) golang.org/x/oauth2/google
	$(GOGET) github.com/nanobox-io/golang-scribble
	$(GOGET) github.com/360EntSecGroup-Skylar/excelize
	$(GOGET) google.golang.org/api/sheets/v4

.PHONY: $(TOPTARGETS) $(SUBDIRS) $(POSTSUBDIRS)

