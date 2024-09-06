package main

// Jsonlint - JSON parser and validator with a CLI
// Homepage: https://github.com/zaach/jsonlint

import (
	"fmt"
	
	"os/exec"
)

func installJsonlint() {
	// Método 1: Descargar y extraer .tar.gz
	jsonlint_tar_url := "https://github.com/zaach/jsonlint/archive/refs/tags/v1.6.0.tar.gz"
	jsonlint_cmd_tar := exec.Command("curl", "-L", jsonlint_tar_url, "-o", "package.tar.gz")
	err := jsonlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonlint_zip_url := "https://github.com/zaach/jsonlint/archive/refs/tags/v1.6.0.zip"
	jsonlint_cmd_zip := exec.Command("curl", "-L", jsonlint_zip_url, "-o", "package.zip")
	err = jsonlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonlint_bin_url := "https://github.com/zaach/jsonlint/archive/refs/tags/v1.6.0.bin"
	jsonlint_cmd_bin := exec.Command("curl", "-L", jsonlint_bin_url, "-o", "binary.bin")
	err = jsonlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonlint_src_url := "https://github.com/zaach/jsonlint/archive/refs/tags/v1.6.0.src.tar.gz"
	jsonlint_cmd_src := exec.Command("curl", "-L", jsonlint_src_url, "-o", "source.tar.gz")
	err = jsonlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonlint_cmd_direct := exec.Command("./binary")
	err = jsonlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
