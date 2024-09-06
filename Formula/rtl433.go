package main

// Rtl433 - Program to decode radio transmissions from devices
// Homepage: https://github.com/merbanan/rtl_433

import (
	"fmt"
	
	"os/exec"
)

func installRtl433() {
	// Método 1: Descargar y extraer .tar.gz
	rtl433_tar_url := "https://github.com/merbanan/rtl_433/archive/refs/tags/23.11.tar.gz"
	rtl433_cmd_tar := exec.Command("curl", "-L", rtl433_tar_url, "-o", "package.tar.gz")
	err := rtl433_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rtl433_zip_url := "https://github.com/merbanan/rtl_433/archive/refs/tags/23.11.zip"
	rtl433_cmd_zip := exec.Command("curl", "-L", rtl433_zip_url, "-o", "package.zip")
	err = rtl433_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rtl433_bin_url := "https://github.com/merbanan/rtl_433/archive/refs/tags/23.11.bin"
	rtl433_cmd_bin := exec.Command("curl", "-L", rtl433_bin_url, "-o", "binary.bin")
	err = rtl433_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rtl433_src_url := "https://github.com/merbanan/rtl_433/archive/refs/tags/23.11.src.tar.gz"
	rtl433_cmd_src := exec.Command("curl", "-L", rtl433_src_url, "-o", "source.tar.gz")
	err = rtl433_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rtl433_cmd_direct := exec.Command("./binary")
	err = rtl433_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librtlsdr")
	exec.Command("latte", "install", "librtlsdr").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
