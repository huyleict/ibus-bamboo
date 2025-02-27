#
# Bamboo - A Vietnamese Input method editor
# Copyright (C) 2018 Luong Thanh Lam <ltlam93@gmail.com>
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
#
PREFIX=/usr

engine_name=bamboo
engine_gui_name=ibus-setup-Bamboo.desktop
ibus_e_name=ibus-engine-$(engine_name)
keyboard_shortcut_editor=keyboard-shortcut-editor
macro_editor=macro-editor
pkg_name=ibus-$(engine_name)
version=0.8.2

engine_dir=$(PREFIX)/share/$(pkg_name)
ibus_dir=$(PREFIX)/share/ibus

GOLDFLAGS=-ldflags "-w -s -X main.Version=$(version)"

rpm_src_dir=~/rpmbuild/SOURCES
tar_file=$(pkg_name)-$(version).tar.gz
rpm_src_tar=$(rpm_src_dir)/$(tar_file)
tar_options_src=--transform "s/^\./$(pkg_name)-$(version)/" --exclude={"*.tar.gz",".git",".idea"} .

all: xml build

xml:
	glib-compile-resources --generate-source setup-ui/keyboard.gresource.xml

build:
	GOPATH=$(CURDIR) GO111MODULE=off go build $(GOLDFLAGS) -o $(ibus_e_name) ibus-$(engine_name)
	gcc -o $(keyboard_shortcut_editor) setup-ui/$(keyboard_shortcut_editor).c `pkg-config --libs --cflags gtk+-3.0`
	gcc -rdynamic -o $(macro_editor) setup-ui/$(macro_editor).c `pkg-config --libs --cflags gtk+-3.0`

t:
	GOPATH=$(CURDIR) GO111MODULE=off go test ./src/ibus-bamboo/...
	GOPATH=$(CURDIR) GO111MODULE=off go test ./src/github.com/BambooEngine/bamboo-core/...

clean:
	rm -f ibus-engine-* *_linux *_cover.html go_test_* go_build_* test *.gz test
	rm -f debian/files
	rm -rf debian/debhelper*
	rm -rf debian/.debhelper
	rm -rf debian/ibus-bamboo*


install: build
	mkdir -p $(DESTDIR)$(engine_dir)
	mkdir -p $(DESTDIR)$(PREFIX)/lib/ibus-$(engine_name)
	mkdir -p $(DESTDIR)$(ibus_dir)/component/
	mkdir -p $(DESTDIR)$(PREFIX)/share/applications/

	cp -R -f icons data $(DESTDIR)$(engine_dir)
	cp -f $(ibus_e_name) $(DESTDIR)$(PREFIX)/lib/ibus-${engine_name}/
	cp -f $(keyboard_shortcut_editor) $(DESTDIR)$(PREFIX)/lib/ibus-$(engine_name)/
	cp -f $(macro_editor) $(DESTDIR)$(PREFIX)/lib/ibus-$(engine_name)/
	cp -f $(engine_name).xml $(DESTDIR)$(ibus_dir)/component/
	cp -f $(engine_gui_name) $(DESTDIR)$(PREFIX)/share/applications/


uninstall:
	sudo rm -rf $(DESTDIR)$(engine_dir)
	sudo rm -rf $(DESTDIR)$(PREFIX)/lib/ibus-$(engine_name)/
	sudo rm -f $(DESTDIR)$(ibus_dir)/component/$(engine_name).xml
	sudo rm -rf $(DESTDIR)$(PREFIX)/share/applications/$(engine_gui_name)


src: clean
	tar -zcf $(DESTDIR)/$(tar_file) $(tar_options_src)
	cp -f $(pkg_name).spec $(DESTDIR)/
	cp -f $(pkg_name).dsc $(DESTDIR)/
	cp -f debian/changelog $(DESTDIR)/debian.changelog
	cp -f debian/control $(DESTDIR)/debian.control
	cp -f debian/compat $(DESTDIR)/debian.compat
	cp -f debian/rules $(DESTDIR)/debian.rules
	cp -f archlinux/PKGBUILD-obs $(DESTDIR)/PKGBUILD


rpm: clean
	tar -zcf $(rpm_src_tar) $(tar_options_src)
	rpmbuild $(pkg_name).spec -ba

deb: clean
	dpkg-buildpackage


.PHONY: build clean build install uninstall src rpm deb
