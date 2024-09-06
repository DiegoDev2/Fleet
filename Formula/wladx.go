package main

// WlaDx - Yet another crossassembler package
// Homepage: https://github.com/vhelin/wla-dx

import (
	"fmt"
	
	"os/exec"
)

func installWlaDx() {
	// Método 1: Descargar y extraer .tar.gz
	wladx_tar_url := "https://github.com/vhelin/wla-dx/archive/refs/tags/v10.6.tar.gz"
	wladx_cmd_tar := exec.Command("curl", "-L", wladx_tar_url, "-o", "package.tar.gz")
	err := wladx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wladx_zip_url := "https://github.com/vhelin/wla-dx/archive/refs/tags/v10.6.zip"
	wladx_cmd_zip := exec.Command("curl", "-L", wladx_zip_url, "-o", "package.zip")
	err = wladx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wladx_bin_url := "https://github.com/vhelin/wla-dx/archive/refs/tags/v10.6.bin"
	wladx_cmd_bin := exec.Command("curl", "-L", wladx_bin_url, "-o", "binary.bin")
	err = wladx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wladx_src_url := "https://github.com/vhelin/wla-dx/archive/refs/tags/v10.6.src.tar.gz"
	wladx_cmd_src := exec.Command("curl", "-L", wladx_src_url, "-o", "source.tar.gz")
	err = wladx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wladx_cmd_direct := exec.Command("./binary")
	err = wladx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
