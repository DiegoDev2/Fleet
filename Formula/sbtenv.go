package main

// Sbtenv - Command-line tool for managing sbt environments
// Homepage: https://github.com/sbtenv/sbtenv

import (
	"fmt"
	
	"os/exec"
)

func installSbtenv() {
	// Método 1: Descargar y extraer .tar.gz
	sbtenv_tar_url := "https://github.com/sbtenv/sbtenv/archive/refs/tags/version/0.0.24.tar.gz"
	sbtenv_cmd_tar := exec.Command("curl", "-L", sbtenv_tar_url, "-o", "package.tar.gz")
	err := sbtenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sbtenv_zip_url := "https://github.com/sbtenv/sbtenv/archive/refs/tags/version/0.0.24.zip"
	sbtenv_cmd_zip := exec.Command("curl", "-L", sbtenv_zip_url, "-o", "package.zip")
	err = sbtenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sbtenv_bin_url := "https://github.com/sbtenv/sbtenv/archive/refs/tags/version/0.0.24.bin"
	sbtenv_cmd_bin := exec.Command("curl", "-L", sbtenv_bin_url, "-o", "binary.bin")
	err = sbtenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sbtenv_src_url := "https://github.com/sbtenv/sbtenv/archive/refs/tags/version/0.0.24.src.tar.gz"
	sbtenv_cmd_src := exec.Command("curl", "-L", sbtenv_src_url, "-o", "source.tar.gz")
	err = sbtenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sbtenv_cmd_direct := exec.Command("./binary")
	err = sbtenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
