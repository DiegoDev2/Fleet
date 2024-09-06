package main

// Run - Easily manage and invoke small scripts and wrappers
// Homepage: https://github.com/TekWizely/run

import (
	"fmt"
	
	"os/exec"
)

func installRun() {
	// Método 1: Descargar y extraer .tar.gz
	run_tar_url := "https://github.com/TekWizely/run/archive/refs/tags/v0.11.2.tar.gz"
	run_cmd_tar := exec.Command("curl", "-L", run_tar_url, "-o", "package.tar.gz")
	err := run_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	run_zip_url := "https://github.com/TekWizely/run/archive/refs/tags/v0.11.2.zip"
	run_cmd_zip := exec.Command("curl", "-L", run_zip_url, "-o", "package.zip")
	err = run_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	run_bin_url := "https://github.com/TekWizely/run/archive/refs/tags/v0.11.2.bin"
	run_cmd_bin := exec.Command("curl", "-L", run_bin_url, "-o", "binary.bin")
	err = run_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	run_src_url := "https://github.com/TekWizely/run/archive/refs/tags/v0.11.2.src.tar.gz"
	run_cmd_src := exec.Command("curl", "-L", run_src_url, "-o", "source.tar.gz")
	err = run_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	run_cmd_direct := exec.Command("./binary")
	err = run_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
