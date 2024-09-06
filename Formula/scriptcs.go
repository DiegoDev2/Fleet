package main

// Scriptcs - Tools to write and execute C#
// Homepage: https://github.com/scriptcs/scriptcs

import (
	"fmt"
	
	"os/exec"
)

func installScriptcs() {
	// Método 1: Descargar y extraer .tar.gz
	scriptcs_tar_url := "https://github.com/scriptcs/scriptcs/archive/refs/tags/v0.17.1.tar.gz"
	scriptcs_cmd_tar := exec.Command("curl", "-L", scriptcs_tar_url, "-o", "package.tar.gz")
	err := scriptcs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scriptcs_zip_url := "https://github.com/scriptcs/scriptcs/archive/refs/tags/v0.17.1.zip"
	scriptcs_cmd_zip := exec.Command("curl", "-L", scriptcs_zip_url, "-o", "package.zip")
	err = scriptcs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scriptcs_bin_url := "https://github.com/scriptcs/scriptcs/archive/refs/tags/v0.17.1.bin"
	scriptcs_cmd_bin := exec.Command("curl", "-L", scriptcs_bin_url, "-o", "binary.bin")
	err = scriptcs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scriptcs_src_url := "https://github.com/scriptcs/scriptcs/archive/refs/tags/v0.17.1.src.tar.gz"
	scriptcs_cmd_src := exec.Command("curl", "-L", scriptcs_src_url, "-o", "source.tar.gz")
	err = scriptcs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scriptcs_cmd_direct := exec.Command("./binary")
	err = scriptcs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mono")
	exec.Command("latte", "install", "mono").Run()
}
