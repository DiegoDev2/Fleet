package main

// Asimov - Automatically exclude development dependencies from Time Machine backups
// Homepage: https://github.com/stevegrunwell/asimov

import (
	"fmt"
	
	"os/exec"
)

func installAsimov() {
	// Método 1: Descargar y extraer .tar.gz
	asimov_tar_url := "https://github.com/stevegrunwell/asimov/archive/refs/tags/v0.3.0.tar.gz"
	asimov_cmd_tar := exec.Command("curl", "-L", asimov_tar_url, "-o", "package.tar.gz")
	err := asimov_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asimov_zip_url := "https://github.com/stevegrunwell/asimov/archive/refs/tags/v0.3.0.zip"
	asimov_cmd_zip := exec.Command("curl", "-L", asimov_zip_url, "-o", "package.zip")
	err = asimov_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asimov_bin_url := "https://github.com/stevegrunwell/asimov/archive/refs/tags/v0.3.0.bin"
	asimov_cmd_bin := exec.Command("curl", "-L", asimov_bin_url, "-o", "binary.bin")
	err = asimov_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asimov_src_url := "https://github.com/stevegrunwell/asimov/archive/refs/tags/v0.3.0.src.tar.gz"
	asimov_cmd_src := exec.Command("curl", "-L", asimov_src_url, "-o", "source.tar.gz")
	err = asimov_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asimov_cmd_direct := exec.Command("./binary")
	err = asimov_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
