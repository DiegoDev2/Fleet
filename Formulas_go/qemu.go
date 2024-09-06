package main

// Qemu - Generic machine emulator and virtualizer
// Homepage: https://www.qemu.org/

import (
	"fmt"
	
	"os/exec"
)

func installQemu() {
	// Método 1: Descargar y extraer .tar.gz
	qemu_tar_url := "https://download.qemu.org/qemu-9.1.0.tar.xz"
	qemu_cmd_tar := exec.Command("curl", "-L", qemu_tar_url, "-o", "package.tar.gz")
	err := qemu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qemu_zip_url := "https://download.qemu.org/qemu-9.1.0.tar.xz"
	qemu_cmd_zip := exec.Command("curl", "-L", qemu_zip_url, "-o", "package.zip")
	err = qemu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qemu_bin_url := "https://download.qemu.org/qemu-9.1.0.tar.xz"
	qemu_cmd_bin := exec.Command("curl", "-L", qemu_bin_url, "-o", "binary.bin")
	err = qemu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qemu_src_url := "https://download.qemu.org/qemu-9.1.0.tar.xz"
	qemu_cmd_src := exec.Command("curl", "-L", qemu_src_url, "-o", "source.tar.gz")
	err = qemu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qemu_cmd_direct := exec.Command("./binary")
	err = qemu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: spice-protocol")
exec.Command("latte", "install", "spice-protocol")
	fmt.Println("Instalando dependencia: capstone")
exec.Command("latte", "install", "capstone")
	fmt.Println("Instalando dependencia: dtc")
exec.Command("latte", "install", "dtc")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libslirp")
exec.Command("latte", "install", "libslirp")
	fmt.Println("Instalando dependencia: libssh")
exec.Command("latte", "install", "libssh")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
	fmt.Println("Instalando dependencia: pixman")
exec.Command("latte", "install", "pixman")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: vde")
exec.Command("latte", "install", "vde")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: attr")
exec.Command("latte", "install", "attr")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libcap-ng")
exec.Command("latte", "install", "libcap-ng")
	fmt.Println("Instalando dependencia: libepoxy")
exec.Command("latte", "install", "libepoxy")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxkbcommon")
exec.Command("latte", "install", "libxkbcommon")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
