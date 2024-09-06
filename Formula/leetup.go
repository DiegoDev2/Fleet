package main

// Leetup - Command-line tool to solve Leetcode problems
// Homepage: https://github.com/dragfire/leetup

import (
	"fmt"
	
	"os/exec"
)

func installLeetup() {
	// Método 1: Descargar y extraer .tar.gz
	leetup_tar_url := "https://github.com/dragfire/leetup/archive/refs/tags/v1.2.5.tar.gz"
	leetup_cmd_tar := exec.Command("curl", "-L", leetup_tar_url, "-o", "package.tar.gz")
	err := leetup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leetup_zip_url := "https://github.com/dragfire/leetup/archive/refs/tags/v1.2.5.zip"
	leetup_cmd_zip := exec.Command("curl", "-L", leetup_zip_url, "-o", "package.zip")
	err = leetup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leetup_bin_url := "https://github.com/dragfire/leetup/archive/refs/tags/v1.2.5.bin"
	leetup_cmd_bin := exec.Command("curl", "-L", leetup_bin_url, "-o", "binary.bin")
	err = leetup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leetup_src_url := "https://github.com/dragfire/leetup/archive/refs/tags/v1.2.5.src.tar.gz"
	leetup_cmd_src := exec.Command("curl", "-L", leetup_src_url, "-o", "source.tar.gz")
	err = leetup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leetup_cmd_direct := exec.Command("./binary")
	err = leetup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
