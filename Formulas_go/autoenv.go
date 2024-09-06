package main

// Autoenv - Per-project, per-directory shell environments
// Homepage: https://github.com/hyperupcall/autoenv

import (
	"fmt"
	
	"os/exec"
)

func installAutoenv() {
	// Método 1: Descargar y extraer .tar.gz
	autoenv_tar_url := "https://github.com/hyperupcall/autoenv/archive/refs/tags/v0.3.0.tar.gz"
	autoenv_cmd_tar := exec.Command("curl", "-L", autoenv_tar_url, "-o", "package.tar.gz")
	err := autoenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autoenv_zip_url := "https://github.com/hyperupcall/autoenv/archive/refs/tags/v0.3.0.zip"
	autoenv_cmd_zip := exec.Command("curl", "-L", autoenv_zip_url, "-o", "package.zip")
	err = autoenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autoenv_bin_url := "https://github.com/hyperupcall/autoenv/archive/refs/tags/v0.3.0.bin"
	autoenv_cmd_bin := exec.Command("curl", "-L", autoenv_bin_url, "-o", "binary.bin")
	err = autoenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autoenv_src_url := "https://github.com/hyperupcall/autoenv/archive/refs/tags/v0.3.0.src.tar.gz"
	autoenv_cmd_src := exec.Command("curl", "-L", autoenv_src_url, "-o", "source.tar.gz")
	err = autoenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autoenv_cmd_direct := exec.Command("./binary")
	err = autoenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
