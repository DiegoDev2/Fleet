package main

// Tfenv - Terraform version manager inspired by rbenv
// Homepage: https://github.com/tfutils/tfenv

import (
	"fmt"
	
	"os/exec"
)

func installTfenv() {
	// Método 1: Descargar y extraer .tar.gz
	tfenv_tar_url := "https://github.com/tfutils/tfenv/archive/refs/tags/v3.0.0.tar.gz"
	tfenv_cmd_tar := exec.Command("curl", "-L", tfenv_tar_url, "-o", "package.tar.gz")
	err := tfenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfenv_zip_url := "https://github.com/tfutils/tfenv/archive/refs/tags/v3.0.0.zip"
	tfenv_cmd_zip := exec.Command("curl", "-L", tfenv_zip_url, "-o", "package.zip")
	err = tfenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfenv_bin_url := "https://github.com/tfutils/tfenv/archive/refs/tags/v3.0.0.bin"
	tfenv_cmd_bin := exec.Command("curl", "-L", tfenv_bin_url, "-o", "binary.bin")
	err = tfenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfenv_src_url := "https://github.com/tfutils/tfenv/archive/refs/tags/v3.0.0.src.tar.gz"
	tfenv_cmd_src := exec.Command("curl", "-L", tfenv_src_url, "-o", "source.tar.gz")
	err = tfenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfenv_cmd_direct := exec.Command("./binary")
	err = tfenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: grep")
exec.Command("latte", "install", "grep")
}
