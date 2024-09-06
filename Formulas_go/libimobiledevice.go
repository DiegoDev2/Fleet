package main

// Libimobiledevice - Library to communicate with iOS devices natively
// Homepage: https://www.libimobiledevice.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibimobiledevice() {
	// Método 1: Descargar y extraer .tar.gz
	libimobiledevice_tar_url := "https://github.com/libimobiledevice/libimobiledevice/releases/download/1.3.0/libimobiledevice-1.3.0.tar.bz2"
	libimobiledevice_cmd_tar := exec.Command("curl", "-L", libimobiledevice_tar_url, "-o", "package.tar.gz")
	err := libimobiledevice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libimobiledevice_zip_url := "https://github.com/libimobiledevice/libimobiledevice/releases/download/1.3.0/libimobiledevice-1.3.0.tar.bz2"
	libimobiledevice_cmd_zip := exec.Command("curl", "-L", libimobiledevice_zip_url, "-o", "package.zip")
	err = libimobiledevice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libimobiledevice_bin_url := "https://github.com/libimobiledevice/libimobiledevice/releases/download/1.3.0/libimobiledevice-1.3.0.tar.bz2"
	libimobiledevice_cmd_bin := exec.Command("curl", "-L", libimobiledevice_bin_url, "-o", "binary.bin")
	err = libimobiledevice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libimobiledevice_src_url := "https://github.com/libimobiledevice/libimobiledevice/releases/download/1.3.0/libimobiledevice-1.3.0.tar.bz2"
	libimobiledevice_cmd_src := exec.Command("curl", "-L", libimobiledevice_src_url, "-o", "source.tar.gz")
	err = libimobiledevice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libimobiledevice_cmd_direct := exec.Command("./binary")
	err = libimobiledevice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libimobiledevice-glue")
exec.Command("latte", "install", "libimobiledevice-glue")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libplist")
exec.Command("latte", "install", "libplist")
	fmt.Println("Instalando dependencia: libtasn1")
exec.Command("latte", "install", "libtasn1")
	fmt.Println("Instalando dependencia: libusbmuxd")
exec.Command("latte", "install", "libusbmuxd")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
