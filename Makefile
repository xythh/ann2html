# ann2html
# See LICENSE file for copyright and license details.

.PHONY: dist linux-x86_64  windows-x86_64 macos-x86_64 macos-arm64

GOOPTIONS = CGO_ENABLED=0
LDFLAGS = -ldflags="-w -s -buildid=" -trimpath -o

all: dist

dist: linux-x86_64  windows-x86_64 macos-x86_64 macos-arm64

linux-x86_64:
	${GOOPTIONS} GOOS=linux GOARCH=amd64 go build ${LDFLAGS} release/$@/ann2html/ann2html
	cp release-includes/README-linux release/$@/ann2html/README
	cp release-includes/ann release/$@/ann2html/ann
	cp release-includes/config release/$@/ann2html/config
	cp release-includes/template.html release/$@/ann2html/template.html
	tar --owner=0 --group=0 --mode='og-w' -czvf release/ann2html-$@.tar.gz -C release/$@ . 
	rm -r release/$@
	

windows-x86_64:
	${GOOPTIONS} GOOS=windows GOARCH=amd64 go build ${LDFLAGS} release/$@/ann2html/ann2html.exe
	cp release-includes/README-windows release/$@/ann2html/README.txt
	cp release-includes/config release/$@/ann2html/config
	cp release-includes/template.html release/$@/ann2html/template.html
	(cd release/$@ && zip -9 -y -r -X - ann2html/ > ../ann2html-$@.zip)
	rm -r release/$@
	

macos-x86_64:
	${GOOPTIONS} GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} release/$@/ann2html/ann2html
	cp release-includes/README-macos release/$@/ann2html/README
	cp release-includes/config release/$@/ann2html/config
	cp release-includes/template.html release/$@/ann2html/template.html
	(cd release/$@ && zip -9 -y -r -X - ann2html/ > ../ann2html-$@.zip)
	rm -r release/$@

macos-arm64:
	${GOOPTIONS} GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} release/$@/ann2html/ann2html
	cp release-includes/README-macos release/$@/ann2html/README
	cp release-includes/config release/$@/ann2html/config
	cp release-includes/template.html release/$@/ann2html/template.html
	(cd release/$@ && zip -9 -y -r -X - ann2html/ > ../ann2html-$@.zip)
	rm -r release/$@

