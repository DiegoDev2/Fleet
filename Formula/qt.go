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
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: vulkan-headers")
	exec.Command("latte", "install", "vulkan-headers").Run()
	fmt.Println("Instalando dependencia: vulkan-loader")
	exec.Command("latte", "install", "vulkan-loader").Run()
	fmt.Println("Instalando dependencia: assimp")
	exec.Command("latte", "install", "assimp").Run()
	fmt.Println("Instalando dependencia: brotli")
	exec.Command("latte", "install", "brotli").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: double-conversion")
	exec.Command("latte", "install", "double-conversion").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: hunspell")
	exec.Command("latte", "install", "hunspell").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: jasper")
	exec.Command("latte", "install", "jasper").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libb2")
	exec.Command("latte", "install", "libb2").Run()
	fmt.Println("Instalando dependencia: libmng")
	exec.Command("latte", "install", "libmng").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: md4c")
	exec.Command("latte", "install", "md4c").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: molten-vk")
	exec.Command("latte", "install", "molten-vk").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: bluez")
	exec.Command("latte", "install", "bluez").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: gstreamer")
	exec.Command("latte", "install", "gstreamer").Run()
	fmt.Println("Instalando dependencia: gypsy")
	exec.Command("latte", "install", "gypsy").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libvpx")
	exec.Command("latte", "install", "libvpx").Run()
	fmt.Println("Instalando dependencia: libxcomposite")
	exec.Command("latte", "install", "libxcomposite").Run()
	fmt.Println("Instalando dependencia: libxkbcommon")
	exec.Command("latte", "install", "libxkbcommon").Run()
	fmt.Println("Instalando dependencia: libxkbfile")
	exec.Command("latte", "install", "libxkbfile").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: libxtst")
	exec.Command("latte", "install", "libxtst").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: nss")
	exec.Command("latte", "install", "nss").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
	fmt.Println("Instalando dependencia: xcb-util")
	exec.Command("latte", "install", "xcb-util").Run()
	fmt.Println("Instalando dependencia: xcb-util-cursor")
	exec.Command("latte", "install", "xcb-util-cursor").Run()
	fmt.Println("Instalando dependencia: xcb-util-image")
	exec.Command("latte", "install", "xcb-util-image").Run()
	fmt.Println("Instalando dependencia: xcb-util-keysyms")
	exec.Command("latte", "install", "xcb-util-keysyms").Run()
	fmt.Println("Instalando dependencia: xcb-util-renderutil")
	exec.Command("latte", "install", "xcb-util-renderutil").Run()
	fmt.Println("Instalando dependencia: xcb-util-wm")
	exec.Command("latte", "install", "xcb-util-wm").Run()
}
