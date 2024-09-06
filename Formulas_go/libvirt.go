package main

// Libvirt - C virtualization API
// Homepage: https://libvirt.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibvirt() {
	// Método 1: Descargar y extraer .tar.gz
	libvirt_tar_url := "https://download.libvirt.org/libvirt-10.7.0.tar.xz"
	libvirt_cmd_tar := exec.Command("curl", "-L", libvirt_tar_url, "-o", "package.tar.gz")
	err := libvirt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvirt_zip_url := "https://download.libvirt.org/libvirt-10.7.0.tar.xz"
	libvirt_cmd_zip := exec.Command("curl", "-L", libvirt_zip_url, "-o", "package.zip")
	err = libvirt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvirt_bin_url := "https://download.libvirt.org/libvirt-10.7.0.tar.xz"
	libvirt_cmd_bin := exec.Command("curl", "-L", libvirt_bin_url, "-o", "binary.bin")
	err = libvirt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvirt_src_url := "https://download.libvirt.org/libvirt-10.7.0.tar.xz"
	libvirt_cmd_src := exec.Command("curl", "-L", libvirt_src_url, "-o", "source.tar.gz")
	err = libvirt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvirt_cmd_direct := exec.Command("./binary")
	err = libvirt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docutils")
exec.Command("latte", "install", "docutils")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libiscsi")
exec.Command("latte", "install", "libiscsi")
	fmt.Println("Instalando dependencia: libssh2")
exec.Command("latte", "install", "libssh2")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: yajl")
exec.Command("latte", "install", "yajl")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: acl")
exec.Command("latte", "install", "acl")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
