package main

// TerrapinScanner - Vulnerability scanner for the Terrapin attack
// Homepage: https://terrapin-attack.com/

import (
	"fmt"
	
	"os/exec"
)

func installTerrapinScanner() {
	// Método 1: Descargar y extraer .tar.gz
	terrapinscanner_tar_url := "https://github.com/RUB-NDS/Terrapin-Scanner/archive/refs/tags/v1.1.3.tar.gz"
	terrapinscanner_cmd_tar := exec.Command("curl", "-L", terrapinscanner_tar_url, "-o", "package.tar.gz")
	err := terrapinscanner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terrapinscanner_zip_url := "https://github.com/RUB-NDS/Terrapin-Scanner/archive/refs/tags/v1.1.3.zip"
	terrapinscanner_cmd_zip := exec.Command("curl", "-L", terrapinscanner_zip_url, "-o", "package.zip")
	err = terrapinscanner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terrapinscanner_bin_url := "https://github.com/RUB-NDS/Terrapin-Scanner/archive/refs/tags/v1.1.3.bin"
	terrapinscanner_cmd_bin := exec.Command("curl", "-L", terrapinscanner_bin_url, "-o", "binary.bin")
	err = terrapinscanner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terrapinscanner_src_url := "https://github.com/RUB-NDS/Terrapin-Scanner/archive/refs/tags/v1.1.3.src.tar.gz"
	terrapinscanner_cmd_src := exec.Command("curl", "-L", terrapinscanner_src_url, "-o", "source.tar.gz")
	err = terrapinscanner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terrapinscanner_cmd_direct := exec.Command("./binary")
	err = terrapinscanner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
