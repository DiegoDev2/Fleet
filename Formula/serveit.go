package main

// Serveit - Synchronous server and rebuilder of static content
// Homepage: https://github.com/garybernhardt/serveit

import (
	"fmt"
	
	"os/exec"
)

func installServeit() {
	// Método 1: Descargar y extraer .tar.gz
	serveit_tar_url := "https://github.com/garybernhardt/serveit/archive/refs/tags/v0.0.3.tar.gz"
	serveit_cmd_tar := exec.Command("curl", "-L", serveit_tar_url, "-o", "package.tar.gz")
	err := serveit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	serveit_zip_url := "https://github.com/garybernhardt/serveit/archive/refs/tags/v0.0.3.zip"
	serveit_cmd_zip := exec.Command("curl", "-L", serveit_zip_url, "-o", "package.zip")
	err = serveit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	serveit_bin_url := "https://github.com/garybernhardt/serveit/archive/refs/tags/v0.0.3.bin"
	serveit_cmd_bin := exec.Command("curl", "-L", serveit_bin_url, "-o", "binary.bin")
	err = serveit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	serveit_src_url := "https://github.com/garybernhardt/serveit/archive/refs/tags/v0.0.3.src.tar.gz"
	serveit_cmd_src := exec.Command("curl", "-L", serveit_src_url, "-o", "source.tar.gz")
	err = serveit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	serveit_cmd_direct := exec.Command("./binary")
	err = serveit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
