package main

// Marcli - Parse MARC (ISO 2709) files
// Homepage: https://github.com/hectorcorrea/marcli

import (
	"fmt"
	
	"os/exec"
)

func installMarcli() {
	// Método 1: Descargar y extraer .tar.gz
	marcli_tar_url := "https://github.com/hectorcorrea/marcli/archive/refs/tags/v1.1.0.tar.gz"
	marcli_cmd_tar := exec.Command("curl", "-L", marcli_tar_url, "-o", "package.tar.gz")
	err := marcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	marcli_zip_url := "https://github.com/hectorcorrea/marcli/archive/refs/tags/v1.1.0.zip"
	marcli_cmd_zip := exec.Command("curl", "-L", marcli_zip_url, "-o", "package.zip")
	err = marcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	marcli_bin_url := "https://github.com/hectorcorrea/marcli/archive/refs/tags/v1.1.0.bin"
	marcli_cmd_bin := exec.Command("curl", "-L", marcli_bin_url, "-o", "binary.bin")
	err = marcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	marcli_src_url := "https://github.com/hectorcorrea/marcli/archive/refs/tags/v1.1.0.src.tar.gz"
	marcli_cmd_src := exec.Command("curl", "-L", marcli_src_url, "-o", "source.tar.gz")
	err = marcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	marcli_cmd_direct := exec.Command("./binary")
	err = marcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
