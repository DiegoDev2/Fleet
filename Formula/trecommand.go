package main

// TreCommand - Tree command, improved
// Homepage: https://github.com/dduan/tre

import (
	"fmt"
	
	"os/exec"
)

func installTreCommand() {
	// Método 1: Descargar y extraer .tar.gz
	trecommand_tar_url := "https://github.com/dduan/tre/archive/refs/tags/v0.4.0.tar.gz"
	trecommand_cmd_tar := exec.Command("curl", "-L", trecommand_tar_url, "-o", "package.tar.gz")
	err := trecommand_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trecommand_zip_url := "https://github.com/dduan/tre/archive/refs/tags/v0.4.0.zip"
	trecommand_cmd_zip := exec.Command("curl", "-L", trecommand_zip_url, "-o", "package.zip")
	err = trecommand_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trecommand_bin_url := "https://github.com/dduan/tre/archive/refs/tags/v0.4.0.bin"
	trecommand_cmd_bin := exec.Command("curl", "-L", trecommand_bin_url, "-o", "binary.bin")
	err = trecommand_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trecommand_src_url := "https://github.com/dduan/tre/archive/refs/tags/v0.4.0.src.tar.gz"
	trecommand_cmd_src := exec.Command("curl", "-L", trecommand_src_url, "-o", "source.tar.gz")
	err = trecommand_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trecommand_cmd_direct := exec.Command("./binary")
	err = trecommand_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
