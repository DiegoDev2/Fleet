package main

// Benthos - Stream processor for mundane tasks written in Go
// Homepage: https://github.com/redpanda-data/benthos

import (
	"fmt"
	
	"os/exec"
)

func installBenthos() {
	// Método 1: Descargar y extraer .tar.gz
	benthos_tar_url := "https://github.com/redpanda-data/benthos/archive/refs/tags/v4.36.0.tar.gz"
	benthos_cmd_tar := exec.Command("curl", "-L", benthos_tar_url, "-o", "package.tar.gz")
	err := benthos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	benthos_zip_url := "https://github.com/redpanda-data/benthos/archive/refs/tags/v4.36.0.zip"
	benthos_cmd_zip := exec.Command("curl", "-L", benthos_zip_url, "-o", "package.zip")
	err = benthos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	benthos_bin_url := "https://github.com/redpanda-data/benthos/archive/refs/tags/v4.36.0.bin"
	benthos_cmd_bin := exec.Command("curl", "-L", benthos_bin_url, "-o", "binary.bin")
	err = benthos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	benthos_src_url := "https://github.com/redpanda-data/benthos/archive/refs/tags/v4.36.0.src.tar.gz"
	benthos_cmd_src := exec.Command("curl", "-L", benthos_src_url, "-o", "source.tar.gz")
	err = benthos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	benthos_cmd_direct := exec.Command("./binary")
	err = benthos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
