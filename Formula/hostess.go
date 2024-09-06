package main

// Hostess - Idempotent command-line utility for managing your /etc/hosts file
// Homepage: https://github.com/cbednarski/hostess

import (
	"fmt"
	
	"os/exec"
)

func installHostess() {
	// Método 1: Descargar y extraer .tar.gz
	hostess_tar_url := "https://github.com/cbednarski/hostess/archive/refs/tags/v0.5.2.tar.gz"
	hostess_cmd_tar := exec.Command("curl", "-L", hostess_tar_url, "-o", "package.tar.gz")
	err := hostess_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hostess_zip_url := "https://github.com/cbednarski/hostess/archive/refs/tags/v0.5.2.zip"
	hostess_cmd_zip := exec.Command("curl", "-L", hostess_zip_url, "-o", "package.zip")
	err = hostess_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hostess_bin_url := "https://github.com/cbednarski/hostess/archive/refs/tags/v0.5.2.bin"
	hostess_cmd_bin := exec.Command("curl", "-L", hostess_bin_url, "-o", "binary.bin")
	err = hostess_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hostess_src_url := "https://github.com/cbednarski/hostess/archive/refs/tags/v0.5.2.src.tar.gz"
	hostess_cmd_src := exec.Command("curl", "-L", hostess_src_url, "-o", "source.tar.gz")
	err = hostess_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hostess_cmd_direct := exec.Command("./binary")
	err = hostess_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
