.PHONY: build
build:
	@wails build

.PHONY: build-flatpak
build-flatpak:
	flatpak remote-add --if-not-exists --user flathub https://dl.flathub.org/repo/flathub.flatpakrepo && \
	flatpak-builder --force-clean --user --install-deps-from=flathub --repo=build/bin/repo build/bin/build-dir build/linux/app.runik.runik.yml && \
	flatpak build-bundle build/bin/repo build/bin/runik_linux.flatpak app.runik.runik

.PHONY: dev
dev:
	@wails dev

# Linux-specific tags and env
.PHONY: dev-linux
dev-linux:
	@WEBKIT_DISABLE_DMABUF_RENDERER=1 wails dev -tags webkit2_41
