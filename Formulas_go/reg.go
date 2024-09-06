package main

// Reg - Docker registry v2 command-line client
// Homepage: https://r.j3ss.co

import (
	"fmt"
	
	"os/exec"
)

func installReg() {
	// Método 1: Descargar y extraer .tar.gz
	reg_tar_url := "https://github.com/genuinetools/reg/archive/refs/tags/v0.16.1.tar.gz"
	reg_cmd_tar := exec.Command("curl", "-L", reg_tar_url, "-o", "package.tar.gz")
	err := reg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reg_zip_url := "https://github.com/genuinetools/reg/archive/refs/tags/v0.16.1.zip"
	reg_cmd_zip := exec.Command("curl", "-L", reg_zip_url, "-o", "package.zip")
	err = reg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reg_bin_url := "https://github.com/genuinetools/reg/archive/refs/tags/v0.16.1.bin"
	reg_cmd_bin := exec.Command("curl", "-L", reg_bin_url, "-o", "binary.bin")
	err = reg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reg_src_url := "https://github.com/genuinetools/reg/archive/refs/tags/v0.16.1.src.tar.gz"
	reg_cmd_src := exec.Command("curl", "-L", reg_src_url, "-o", "source.tar.gz")
	err = reg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reg_cmd_direct := exec.Command("./binary")
	err = reg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
