package main

// Libdrm - Library for accessing the direct rendering manager
// Homepage: https://dri.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installLibdrm() {
	// Método 1: Descargar y extraer .tar.gz
	libdrm_tar_url := "https://dri.freedesktop.org/libdrm/libdrm-2.4.123.tar.xz"
	libdrm_cmd_tar := exec.Command("curl", "-L", libdrm_tar_url, "-o", "package.tar.gz")
	err := libdrm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdrm_zip_url := "https://dri.freedesktop.org/libdrm/libdrm-2.4.123.tar.xz"
	libdrm_cmd_zip := exec.Command("curl", "-L", libdrm_zip_url, "-o", "package.zip")
	err = libdrm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdrm_bin_url := "https://dri.freedesktop.org/libdrm/libdrm-2.4.123.tar.xz"
	libdrm_cmd_bin := exec.Command("curl", "-L", libdrm_bin_url, "-o", "binary.bin")
	err = libdrm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdrm_src_url := "https://dri.freedesktop.org/libdrm/libdrm-2.4.123.tar.xz"
	libdrm_cmd_src := exec.Command("curl", "-L", libdrm_src_url, "-o", "source.tar.gz")
	err = libdrm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdrm_cmd_direct := exec.Command("./binary")
	err = libdrm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docutils")
exec.Command("latte", "install", "docutils")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libpciaccess")
exec.Command("latte", "install", "libpciaccess")
}
