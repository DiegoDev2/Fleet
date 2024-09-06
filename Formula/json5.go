package main

// Json5 - JSON enhanced with usability features
// Homepage: https://json5.org/

import (
	"fmt"
	
	"os/exec"
)

func installJson5() {
	// Método 1: Descargar y extraer .tar.gz
	json5_tar_url := "https://github.com/json5/json5/archive/refs/tags/v2.2.3.tar.gz"
	json5_cmd_tar := exec.Command("curl", "-L", json5_tar_url, "-o", "package.tar.gz")
	err := json5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	json5_zip_url := "https://github.com/json5/json5/archive/refs/tags/v2.2.3.zip"
	json5_cmd_zip := exec.Command("curl", "-L", json5_zip_url, "-o", "package.zip")
	err = json5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	json5_bin_url := "https://github.com/json5/json5/archive/refs/tags/v2.2.3.bin"
	json5_cmd_bin := exec.Command("curl", "-L", json5_bin_url, "-o", "binary.bin")
	err = json5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	json5_src_url := "https://github.com/json5/json5/archive/refs/tags/v2.2.3.src.tar.gz"
	json5_cmd_src := exec.Command("curl", "-L", json5_src_url, "-o", "source.tar.gz")
	err = json5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	json5_cmd_direct := exec.Command("./binary")
	err = json5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
