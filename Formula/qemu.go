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
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: spice-protocol")
	exec.Command("latte", "install", "spice-protocol").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
	fmt.Println("Instalando dependencia: dtc")
	exec.Command("latte", "install", "dtc").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libslirp")
	exec.Command("latte", "install", "libslirp").Run()
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
	fmt.Println("Instalando dependencia: vde")
	exec.Command("latte", "install", "vde").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: attr")
	exec.Command("latte", "install", "attr").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: libcap-ng")
	exec.Command("latte", "install", "libcap-ng").Run()
	fmt.Println("Instalando dependencia: libepoxy")
	exec.Command("latte", "install", "libepoxy").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxkbcommon")
	exec.Command("latte", "install", "libxkbcommon").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
