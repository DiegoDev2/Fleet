package main

// Juliaup - Julia installer and version multiplexer
// Homepage: https://github.com/JuliaLang/juliaup

import (
	"fmt"
	
	"os/exec"
)

func installJuliaup() {
	// Método 1: Descargar y extraer .tar.gz
	juliaup_tar_url := "https://github.com/JuliaLang/juliaup/archive/refs/tags/v1.17.4.tar.gz"
	juliaup_cmd_tar := exec.Command("curl", "-L", juliaup_tar_url, "-o", "package.tar.gz")
	err := juliaup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	juliaup_zip_url := "https://github.com/JuliaLang/juliaup/archive/refs/tags/v1.17.4.zip"
	juliaup_cmd_zip := exec.Command("curl", "-L", juliaup_zip_url, "-o", "package.zip")
	err = juliaup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	juliaup_bin_url := "https://github.com/JuliaLang/juliaup/archive/refs/tags/v1.17.4.bin"
	juliaup_cmd_bin := exec.Command("curl", "-L", juliaup_bin_url, "-o", "binary.bin")
	err = juliaup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	juliaup_src_url := "https://github.com/JuliaLang/juliaup/archive/refs/tags/v1.17.4.src.tar.gz"
	juliaup_cmd_src := exec.Command("curl", "-L", juliaup_src_url, "-o", "source.tar.gz")
	err = juliaup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	juliaup_cmd_direct := exec.Command("./binary")
	err = juliaup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
