package main

// Cherrybomb - Tool designed to validate your spec
// Homepage: https://blstsecurity.com

import (
	"fmt"
	
	"os/exec"
)

func installCherrybomb() {
	// Método 1: Descargar y extraer .tar.gz
	cherrybomb_tar_url := "https://github.com/blst-security/cherrybomb/archive/refs/tags/v1.0.1.tar.gz"
	cherrybomb_cmd_tar := exec.Command("curl", "-L", cherrybomb_tar_url, "-o", "package.tar.gz")
	err := cherrybomb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cherrybomb_zip_url := "https://github.com/blst-security/cherrybomb/archive/refs/tags/v1.0.1.zip"
	cherrybomb_cmd_zip := exec.Command("curl", "-L", cherrybomb_zip_url, "-o", "package.zip")
	err = cherrybomb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cherrybomb_bin_url := "https://github.com/blst-security/cherrybomb/archive/refs/tags/v1.0.1.bin"
	cherrybomb_cmd_bin := exec.Command("curl", "-L", cherrybomb_bin_url, "-o", "binary.bin")
	err = cherrybomb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cherrybomb_src_url := "https://github.com/blst-security/cherrybomb/archive/refs/tags/v1.0.1.src.tar.gz"
	cherrybomb_cmd_src := exec.Command("curl", "-L", cherrybomb_src_url, "-o", "source.tar.gz")
	err = cherrybomb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cherrybomb_cmd_direct := exec.Command("./binary")
	err = cherrybomb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
