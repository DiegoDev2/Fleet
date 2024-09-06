package main

// Ideviceinstaller - Tool for managing apps on iOS devices
// Homepage: https://www.libimobiledevice.org/

import (
	"fmt"
	
	"os/exec"
)

func installIdeviceinstaller() {
	// Método 1: Descargar y extraer .tar.gz
	ideviceinstaller_tar_url := "https://github.com/libimobiledevice/ideviceinstaller/releases/download/1.1.1/ideviceinstaller-1.1.1.tar.bz2"
	ideviceinstaller_cmd_tar := exec.Command("curl", "-L", ideviceinstaller_tar_url, "-o", "package.tar.gz")
	err := ideviceinstaller_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ideviceinstaller_zip_url := "https://github.com/libimobiledevice/ideviceinstaller/releases/download/1.1.1/ideviceinstaller-1.1.1.tar.bz2"
	ideviceinstaller_cmd_zip := exec.Command("curl", "-L", ideviceinstaller_zip_url, "-o", "package.zip")
	err = ideviceinstaller_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ideviceinstaller_bin_url := "https://github.com/libimobiledevice/ideviceinstaller/releases/download/1.1.1/ideviceinstaller-1.1.1.tar.bz2"
	ideviceinstaller_cmd_bin := exec.Command("curl", "-L", ideviceinstaller_bin_url, "-o", "binary.bin")
	err = ideviceinstaller_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ideviceinstaller_src_url := "https://github.com/libimobiledevice/ideviceinstaller/releases/download/1.1.1/ideviceinstaller-1.1.1.tar.bz2"
	ideviceinstaller_cmd_src := exec.Command("curl", "-L", ideviceinstaller_src_url, "-o", "source.tar.gz")
	err = ideviceinstaller_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ideviceinstaller_cmd_direct := exec.Command("./binary")
	err = ideviceinstaller_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libimobiledevice")
exec.Command("latte", "install", "libimobiledevice")
	fmt.Println("Instalando dependencia: libplist")
exec.Command("latte", "install", "libplist")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
}
