package main

// Tgenv - Terragrunt version manager inspired by tfenv
// Homepage: https://github.com/cunymatthieu/tgenv

import (
	"fmt"
	
	"os/exec"
)

func installTgenv() {
	// Método 1: Descargar y extraer .tar.gz
	tgenv_tar_url := "https://github.com/cunymatthieu/tgenv/archive/refs/tags/v0.0.3.tar.gz"
	tgenv_cmd_tar := exec.Command("curl", "-L", tgenv_tar_url, "-o", "package.tar.gz")
	err := tgenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tgenv_zip_url := "https://github.com/cunymatthieu/tgenv/archive/refs/tags/v0.0.3.zip"
	tgenv_cmd_zip := exec.Command("curl", "-L", tgenv_zip_url, "-o", "package.zip")
	err = tgenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tgenv_bin_url := "https://github.com/cunymatthieu/tgenv/archive/refs/tags/v0.0.3.bin"
	tgenv_cmd_bin := exec.Command("curl", "-L", tgenv_bin_url, "-o", "binary.bin")
	err = tgenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tgenv_src_url := "https://github.com/cunymatthieu/tgenv/archive/refs/tags/v0.0.3.src.tar.gz"
	tgenv_cmd_src := exec.Command("curl", "-L", tgenv_src_url, "-o", "source.tar.gz")
	err = tgenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tgenv_cmd_direct := exec.Command("./binary")
	err = tgenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
