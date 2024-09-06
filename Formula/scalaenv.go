package main

// Scalaenv - Command-line tool to manage Scala environments
// Homepage: https://github.com/scalaenv/scalaenv

import (
	"fmt"
	
	"os/exec"
)

func installScalaenv() {
	// Método 1: Descargar y extraer .tar.gz
	scalaenv_tar_url := "https://github.com/scalaenv/scalaenv/archive/refs/tags/version/0.1.14.tar.gz"
	scalaenv_cmd_tar := exec.Command("curl", "-L", scalaenv_tar_url, "-o", "package.tar.gz")
	err := scalaenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scalaenv_zip_url := "https://github.com/scalaenv/scalaenv/archive/refs/tags/version/0.1.14.zip"
	scalaenv_cmd_zip := exec.Command("curl", "-L", scalaenv_zip_url, "-o", "package.zip")
	err = scalaenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scalaenv_bin_url := "https://github.com/scalaenv/scalaenv/archive/refs/tags/version/0.1.14.bin"
	scalaenv_cmd_bin := exec.Command("curl", "-L", scalaenv_bin_url, "-o", "binary.bin")
	err = scalaenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scalaenv_src_url := "https://github.com/scalaenv/scalaenv/archive/refs/tags/version/0.1.14.src.tar.gz"
	scalaenv_cmd_src := exec.Command("curl", "-L", scalaenv_src_url, "-o", "source.tar.gz")
	err = scalaenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scalaenv_cmd_direct := exec.Command("./binary")
	err = scalaenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
