package main

// Libirecovery - Library and utility to talk to iBoot/iBSS via USB
// Homepage: https://www.libimobiledevice.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibirecovery() {
	// Método 1: Descargar y extraer .tar.gz
	libirecovery_tar_url := "https://github.com/libimobiledevice/libirecovery/releases/download/1.2.0/libirecovery-1.2.0.tar.bz2"
	libirecovery_cmd_tar := exec.Command("curl", "-L", libirecovery_tar_url, "-o", "package.tar.gz")
	err := libirecovery_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libirecovery_zip_url := "https://github.com/libimobiledevice/libirecovery/releases/download/1.2.0/libirecovery-1.2.0.tar.bz2"
	libirecovery_cmd_zip := exec.Command("curl", "-L", libirecovery_zip_url, "-o", "package.zip")
	err = libirecovery_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libirecovery_bin_url := "https://github.com/libimobiledevice/libirecovery/releases/download/1.2.0/libirecovery-1.2.0.tar.bz2"
	libirecovery_cmd_bin := exec.Command("curl", "-L", libirecovery_bin_url, "-o", "binary.bin")
	err = libirecovery_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libirecovery_src_url := "https://github.com/libimobiledevice/libirecovery/releases/download/1.2.0/libirecovery-1.2.0.tar.bz2"
	libirecovery_cmd_src := exec.Command("curl", "-L", libirecovery_src_url, "-o", "source.tar.gz")
	err = libirecovery_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libirecovery_cmd_direct := exec.Command("./binary")
	err = libirecovery_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libimobiledevice-glue")
	exec.Command("latte", "install", "libimobiledevice-glue").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
