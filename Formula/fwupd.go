package main

// Fwupd - Firmware update daemon
// Homepage: https://github.com/fwupd/fwupd

import (
	"fmt"
	
	"os/exec"
)

func installFwupd() {
	// Método 1: Descargar y extraer .tar.gz
	fwupd_tar_url := "https://github.com/fwupd/fwupd/releases/download/1.9.24/fwupd-1.9.24.tar.xz"
	fwupd_cmd_tar := exec.Command("curl", "-L", fwupd_tar_url, "-o", "package.tar.gz")
	err := fwupd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fwupd_zip_url := "https://github.com/fwupd/fwupd/releases/download/1.9.24/fwupd-1.9.24.tar.xz"
	fwupd_cmd_zip := exec.Command("curl", "-L", fwupd_zip_url, "-o", "package.zip")
	err = fwupd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fwupd_bin_url := "https://github.com/fwupd/fwupd/releases/download/1.9.24/fwupd-1.9.24.tar.xz"
	fwupd_cmd_bin := exec.Command("curl", "-L", fwupd_bin_url, "-o", "binary.bin")
	err = fwupd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fwupd_src_url := "https://github.com/fwupd/fwupd/releases/download/1.9.24/fwupd-1.9.24.tar.xz"
	fwupd_cmd_src := exec.Command("curl", "-L", fwupd_src_url, "-o", "source.tar.gz")
	err = fwupd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fwupd_cmd_direct := exec.Command("./binary")
	err = fwupd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gi-docgen")
	exec.Command("latte", "install", "gi-docgen").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: gcab")
	exec.Command("latte", "install", "gcab").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libcbor")
	exec.Command("latte", "install", "libcbor").Run()
	fmt.Println("Instalando dependencia: libgusb")
	exec.Command("latte", "install", "libgusb").Run()
	fmt.Println("Instalando dependencia: libjcat")
	exec.Command("latte", "install", "libjcat").Run()
	fmt.Println("Instalando dependencia: libxmlb")
	exec.Command("latte", "install", "libxmlb").Run()
	fmt.Println("Instalando dependencia: protobuf-c")
	exec.Command("latte", "install", "protobuf-c").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
