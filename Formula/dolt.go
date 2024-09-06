package main

// Dolt - Git for Data
// Homepage: https://github.com/dolthub/dolt

import (
	"fmt"
	
	"os/exec"
)

func installDolt() {
	// Método 1: Descargar y extraer .tar.gz
	dolt_tar_url := "https://github.com/dolthub/dolt/archive/refs/tags/v1.42.18.tar.gz"
	dolt_cmd_tar := exec.Command("curl", "-L", dolt_tar_url, "-o", "package.tar.gz")
	err := dolt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dolt_zip_url := "https://github.com/dolthub/dolt/archive/refs/tags/v1.42.18.zip"
	dolt_cmd_zip := exec.Command("curl", "-L", dolt_zip_url, "-o", "package.zip")
	err = dolt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dolt_bin_url := "https://github.com/dolthub/dolt/archive/refs/tags/v1.42.18.bin"
	dolt_cmd_bin := exec.Command("curl", "-L", dolt_bin_url, "-o", "binary.bin")
	err = dolt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dolt_src_url := "https://github.com/dolthub/dolt/archive/refs/tags/v1.42.18.src.tar.gz"
	dolt_cmd_src := exec.Command("curl", "-L", dolt_src_url, "-o", "source.tar.gz")
	err = dolt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dolt_cmd_direct := exec.Command("./binary")
	err = dolt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
