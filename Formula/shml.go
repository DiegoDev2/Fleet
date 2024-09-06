package main

// Shml - Style Framework for The Terminal
// Homepage: https://odb.github.io/shml/

import (
	"fmt"
	
	"os/exec"
)

func installShml() {
	// Método 1: Descargar y extraer .tar.gz
	shml_tar_url := "https://github.com/odb/shml/archive/refs/tags/1.1.0.tar.gz"
	shml_cmd_tar := exec.Command("curl", "-L", shml_tar_url, "-o", "package.tar.gz")
	err := shml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shml_zip_url := "https://github.com/odb/shml/archive/refs/tags/1.1.0.zip"
	shml_cmd_zip := exec.Command("curl", "-L", shml_zip_url, "-o", "package.zip")
	err = shml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shml_bin_url := "https://github.com/odb/shml/archive/refs/tags/1.1.0.bin"
	shml_cmd_bin := exec.Command("curl", "-L", shml_bin_url, "-o", "binary.bin")
	err = shml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shml_src_url := "https://github.com/odb/shml/archive/refs/tags/1.1.0.src.tar.gz"
	shml_cmd_src := exec.Command("curl", "-L", shml_src_url, "-o", "source.tar.gz")
	err = shml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shml_cmd_direct := exec.Command("./binary")
	err = shml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
