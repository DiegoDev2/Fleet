package main

// Uuu - Universal Update Utility, mfgtools 3.0. NXP I.MX Chip image deploy tools
// Homepage: https://github.com/nxp-imx/mfgtools

import (
	"fmt"
	
	"os/exec"
)

func installUuu() {
	// Método 1: Descargar y extraer .tar.gz
	uuu_tar_url := "https://github.com/nxp-imx/mfgtools/releases/download/uuu_1.5.182/uuu_source-uuu_1.5.182.tar.gz"
	uuu_cmd_tar := exec.Command("curl", "-L", uuu_tar_url, "-o", "package.tar.gz")
	err := uuu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uuu_zip_url := "https://github.com/nxp-imx/mfgtools/releases/download/uuu_1.5.182/uuu_source-uuu_1.5.182.zip"
	uuu_cmd_zip := exec.Command("curl", "-L", uuu_zip_url, "-o", "package.zip")
	err = uuu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uuu_bin_url := "https://github.com/nxp-imx/mfgtools/releases/download/uuu_1.5.182/uuu_source-uuu_1.5.182.bin"
	uuu_cmd_bin := exec.Command("curl", "-L", uuu_bin_url, "-o", "binary.bin")
	err = uuu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uuu_src_url := "https://github.com/nxp-imx/mfgtools/releases/download/uuu_1.5.182/uuu_source-uuu_1.5.182.src.tar.gz"
	uuu_cmd_src := exec.Command("curl", "-L", uuu_src_url, "-o", "source.tar.gz")
	err = uuu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uuu_cmd_direct := exec.Command("./binary")
	err = uuu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: tinyxml2")
	exec.Command("latte", "install", "tinyxml2").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
