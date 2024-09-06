package main

// Massdriver - Manage applications and infrastructure on Massdriver Cloud
// Homepage: https://www.massdriver.cloud/

import (
	"fmt"
	
	"os/exec"
)

func installMassdriver() {
	// Método 1: Descargar y extraer .tar.gz
	massdriver_tar_url := "https://github.com/massdriver-cloud/mass/archive/refs/tags/1.9.0.tar.gz"
	massdriver_cmd_tar := exec.Command("curl", "-L", massdriver_tar_url, "-o", "package.tar.gz")
	err := massdriver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	massdriver_zip_url := "https://github.com/massdriver-cloud/mass/archive/refs/tags/1.9.0.zip"
	massdriver_cmd_zip := exec.Command("curl", "-L", massdriver_zip_url, "-o", "package.zip")
	err = massdriver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	massdriver_bin_url := "https://github.com/massdriver-cloud/mass/archive/refs/tags/1.9.0.bin"
	massdriver_cmd_bin := exec.Command("curl", "-L", massdriver_bin_url, "-o", "binary.bin")
	err = massdriver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	massdriver_src_url := "https://github.com/massdriver-cloud/mass/archive/refs/tags/1.9.0.src.tar.gz"
	massdriver_cmd_src := exec.Command("curl", "-L", massdriver_src_url, "-o", "source.tar.gz")
	err = massdriver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	massdriver_cmd_direct := exec.Command("./binary")
	err = massdriver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
