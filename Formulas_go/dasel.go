package main

// Dasel - JSON, YAML, TOML, XML, and CSV query and modification tool
// Homepage: https://github.com/TomWright/dasel

import (
	"fmt"
	
	"os/exec"
)

func installDasel() {
	// Método 1: Descargar y extraer .tar.gz
	dasel_tar_url := "https://github.com/TomWright/dasel/archive/refs/tags/v2.8.1.tar.gz"
	dasel_cmd_tar := exec.Command("curl", "-L", dasel_tar_url, "-o", "package.tar.gz")
	err := dasel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dasel_zip_url := "https://github.com/TomWright/dasel/archive/refs/tags/v2.8.1.zip"
	dasel_cmd_zip := exec.Command("curl", "-L", dasel_zip_url, "-o", "package.zip")
	err = dasel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dasel_bin_url := "https://github.com/TomWright/dasel/archive/refs/tags/v2.8.1.bin"
	dasel_cmd_bin := exec.Command("curl", "-L", dasel_bin_url, "-o", "binary.bin")
	err = dasel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dasel_src_url := "https://github.com/TomWright/dasel/archive/refs/tags/v2.8.1.src.tar.gz"
	dasel_cmd_src := exec.Command("curl", "-L", dasel_src_url, "-o", "source.tar.gz")
	err = dasel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dasel_cmd_direct := exec.Command("./binary")
	err = dasel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
