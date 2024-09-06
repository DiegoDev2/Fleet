package main

// Upx - Compress/expand executable files
// Homepage: https://upx.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installUpx() {
	// Método 1: Descargar y extraer .tar.gz
	upx_tar_url := "https://github.com/upx/upx/releases/download/v4.1.0/upx-4.1.0-src.tar.xz"
	upx_cmd_tar := exec.Command("curl", "-L", upx_tar_url, "-o", "package.tar.gz")
	err := upx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	upx_zip_url := "https://github.com/upx/upx/releases/download/v4.1.0/upx-4.1.0-src.tar.xz"
	upx_cmd_zip := exec.Command("curl", "-L", upx_zip_url, "-o", "package.zip")
	err = upx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	upx_bin_url := "https://github.com/upx/upx/releases/download/v4.1.0/upx-4.1.0-src.tar.xz"
	upx_cmd_bin := exec.Command("curl", "-L", upx_bin_url, "-o", "binary.bin")
	err = upx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	upx_src_url := "https://github.com/upx/upx/releases/download/v4.1.0/upx-4.1.0-src.tar.xz"
	upx_cmd_src := exec.Command("curl", "-L", upx_src_url, "-o", "source.tar.gz")
	err = upx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	upx_cmd_direct := exec.Command("./binary")
	err = upx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ucl")
	exec.Command("latte", "install", "ucl").Run()
}
