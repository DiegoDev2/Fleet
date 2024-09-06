package main

// Vegeta - HTTP load testing tool and library
// Homepage: https://github.com/tsenart/vegeta

import (
	"fmt"
	
	"os/exec"
)

func installVegeta() {
	// Método 1: Descargar y extraer .tar.gz
	vegeta_tar_url := "https://github.com/tsenart/vegeta/archive/refs/tags/v12.12.0.tar.gz"
	vegeta_cmd_tar := exec.Command("curl", "-L", vegeta_tar_url, "-o", "package.tar.gz")
	err := vegeta_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vegeta_zip_url := "https://github.com/tsenart/vegeta/archive/refs/tags/v12.12.0.zip"
	vegeta_cmd_zip := exec.Command("curl", "-L", vegeta_zip_url, "-o", "package.zip")
	err = vegeta_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vegeta_bin_url := "https://github.com/tsenart/vegeta/archive/refs/tags/v12.12.0.bin"
	vegeta_cmd_bin := exec.Command("curl", "-L", vegeta_bin_url, "-o", "binary.bin")
	err = vegeta_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vegeta_src_url := "https://github.com/tsenart/vegeta/archive/refs/tags/v12.12.0.src.tar.gz"
	vegeta_cmd_src := exec.Command("curl", "-L", vegeta_src_url, "-o", "source.tar.gz")
	err = vegeta_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vegeta_cmd_direct := exec.Command("./binary")
	err = vegeta_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
