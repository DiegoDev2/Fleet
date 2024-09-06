package main

// Qt - Cross-platform application and UI framework
// Homepage: https://www.qt.io/

import (
	"fmt"
	
	"os/exec"
)

func installQt() {
	// Método 1: Descargar y extraer .tar.gz
	qt_tar_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/single/qt-everywhere-src-6.7.0.tar.xz"
	qt_cmd_tar := exec.Command("curl", "-L", qt_tar_url, "-o", "package.tar.gz")
	err := qt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qt_zip_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/single/qt-everywhere-src-6.7.0.tar.xz"
	qt_cmd_zip := exec.Command("curl", "-L", qt_zip_url, "-o", "package.zip")
	err = qt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qt_bin_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/single/qt-everywhere-src-6.7.0.tar.xz"
	qt_cmd_bin := exec.Command("curl", "-L", qt_bin_url, "-o", "binary.bin")
	err = qt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qt_src_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/single/qt-everywhere-src-6.7.0.tar.xz"
	qt_cmd_src := exec.Command("curl", "-L", qt_src_url, "-o", "source.tar.gz")
	err = qt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qt_cmd_direct := exec.Command("./binary")
	err = qt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: vulkan-headers")
exec.Command("latte", "install", "vulkan-headers")
	fmt.Println("Instalando dependencia: vulkan-loader")
exec.Command("latte", "install", "vulkan-loader")
	fmt.Println("Instalando dependencia: assimp")
exec.Command("latte", "install", "assimp")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: double-conversion")
exec.Command("latte", "install", "double-conversion")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: hunspell")
exec.Command("latte", "install", "hunspell")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: jasper")
exec.Command("latte", "install", "jasper")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libb2")
exec.Command("latte", "install", "libb2")
	fmt.Println("Instalando dependencia: libmng")
exec.Command("latte", "install", "libmng")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: md4c")
exec.Command("latte", "install", "md4c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: molten-vk")
exec.Command("latte", "install", "molten-vk")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: bluez")
exec.Command("latte", "install", "bluez")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: gstreamer")
exec.Command("latte", "install", "gstreamer")
	fmt.Println("Instalando dependencia: gypsy")
exec.Command("latte", "install", "gypsy")
	fmt.Println("Instalando dependencia: libdrm")
exec.Command("latte", "install", "libdrm")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libice")
exec.Command("latte", "install", "libice")
	fmt.Println("Instalando dependencia: libsm")
exec.Command("latte", "install", "libsm")
	fmt.Println("Instalando dependencia: libvpx")
exec.Command("latte", "install", "libvpx")
	fmt.Println("Instalando dependencia: libxcomposite")
exec.Command("latte", "install", "libxcomposite")
	fmt.Println("Instalando dependencia: libxkbcommon")
exec.Command("latte", "install", "libxkbcommon")
	fmt.Println("Instalando dependencia: libxkbfile")
exec.Command("latte", "install", "libxkbfile")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxtst")
exec.Command("latte", "install", "libxtst")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: minizip")
exec.Command("latte", "install", "minizip")
	fmt.Println("Instalando dependencia: nss")
exec.Command("latte", "install", "nss")
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
	fmt.Println("Instalando dependencia: wayland")
exec.Command("latte", "install", "wayland")
	fmt.Println("Instalando dependencia: xcb-util")
exec.Command("latte", "install", "xcb-util")
	fmt.Println("Instalando dependencia: xcb-util-cursor")
exec.Command("latte", "install", "xcb-util-cursor")
	fmt.Println("Instalando dependencia: xcb-util-image")
exec.Command("latte", "install", "xcb-util-image")
	fmt.Println("Instalando dependencia: xcb-util-keysyms")
exec.Command("latte", "install", "xcb-util-keysyms")
	fmt.Println("Instalando dependencia: xcb-util-renderutil")
exec.Command("latte", "install", "xcb-util-renderutil")
	fmt.Println("Instalando dependencia: xcb-util-wm")
exec.Command("latte", "install", "xcb-util-wm")
}
