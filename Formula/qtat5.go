package main

// QtAT5 - Cross-platform application and UI framework
// Homepage: https://www.qt.io/

import (
	"fmt"
	
	"os/exec"
)

func installQtAT5() {
	// Método 1: Descargar y extraer .tar.gz
	qtat5_tar_url := "https://download.qt.io/official_releases/qt/5.15/5.15.13/single/qt-everywhere-opensource-src-5.15.13.tar.xz"
	qtat5_cmd_tar := exec.Command("curl", "-L", qtat5_tar_url, "-o", "package.tar.gz")
	err := qtat5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qtat5_zip_url := "https://download.qt.io/official_releases/qt/5.15/5.15.13/single/qt-everywhere-opensource-src-5.15.13.tar.xz"
	qtat5_cmd_zip := exec.Command("curl", "-L", qtat5_zip_url, "-o", "package.zip")
	err = qtat5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qtat5_bin_url := "https://download.qt.io/official_releases/qt/5.15/5.15.13/single/qt-everywhere-opensource-src-5.15.13.tar.xz"
	qtat5_cmd_bin := exec.Command("curl", "-L", qtat5_bin_url, "-o", "binary.bin")
	err = qtat5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qtat5_src_url := "https://download.qt.io/official_releases/qt/5.15/5.15.13/single/qt-everywhere-opensource-src-5.15.13.tar.xz"
	qtat5_cmd_src := exec.Command("curl", "-L", qtat5_src_url, "-o", "source.tar.gz")
	err = qtat5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qtat5_cmd_direct := exec.Command("./binary")
	err = qtat5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libproxy")
	exec.Command("latte", "install", "libproxy").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libvpx")
	exec.Command("latte", "install", "libvpx").Run()
	fmt.Println("Instalando dependencia: libxcomposite")
	exec.Command("latte", "install", "libxcomposite").Run()
	fmt.Println("Instalando dependencia: libxdamage")
	exec.Command("latte", "install", "libxdamage").Run()
	fmt.Println("Instalando dependencia: libxkbcommon")
	exec.Command("latte", "install", "libxkbcommon").Run()
	fmt.Println("Instalando dependencia: libxkbfile")
	exec.Command("latte", "install", "libxkbfile").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: libxtst")
	exec.Command("latte", "install", "libxtst").Run()
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
	fmt.Println("Instalando dependencia: xcb-util-image")
	exec.Command("latte", "install", "xcb-util-image").Run()
	fmt.Println("Instalando dependencia: xcb-util-keysyms")
	exec.Command("latte", "install", "xcb-util-keysyms").Run()
	fmt.Println("Instalando dependencia: xcb-util-renderutil")
	exec.Command("latte", "install", "xcb-util-renderutil").Run()
	fmt.Println("Instalando dependencia: xcb-util-wm")
	exec.Command("latte", "install", "xcb-util-wm").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
