package main

// Polyml - Standard ML implementation
// Homepage: https://www.polyml.org/

import (
	"fmt"
	
	"os/exec"
)

func installPolyml() {
	// Método 1: Descargar y extraer .tar.gz
	polyml_tar_url := "https://github.com/polyml/polyml/archive/refs/tags/v5.9.1.tar.gz"
	polyml_cmd_tar := exec.Command("curl", "-L", polyml_tar_url, "-o", "package.tar.gz")
	err := polyml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	polyml_zip_url := "https://github.com/polyml/polyml/archive/refs/tags/v5.9.1.zip"
	polyml_cmd_zip := exec.Command("curl", "-L", polyml_zip_url, "-o", "package.zip")
	err = polyml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	polyml_bin_url := "https://github.com/polyml/polyml/archive/refs/tags/v5.9.1.bin"
	polyml_cmd_bin := exec.Command("curl", "-L", polyml_bin_url, "-o", "binary.bin")
	err = polyml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	polyml_src_url := "https://github.com/polyml/polyml/archive/refs/tags/v5.9.1.src.tar.gz"
	polyml_cmd_src := exec.Command("curl", "-L", polyml_src_url, "-o", "source.tar.gz")
	err = polyml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	polyml_cmd_direct := exec.Command("./binary")
	err = polyml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
