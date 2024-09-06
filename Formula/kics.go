package main

// Kics - Detect vulnerabilities, compliance issues, and misconfigurations
// Homepage: https://kics.io/

import (
	"fmt"
	
	"os/exec"
)

func installKics() {
	// Método 1: Descargar y extraer .tar.gz
	kics_tar_url := "https://github.com/Checkmarx/kics/archive/refs/tags/v2.1.2.tar.gz"
	kics_cmd_tar := exec.Command("curl", "-L", kics_tar_url, "-o", "package.tar.gz")
	err := kics_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kics_zip_url := "https://github.com/Checkmarx/kics/archive/refs/tags/v2.1.2.zip"
	kics_cmd_zip := exec.Command("curl", "-L", kics_zip_url, "-o", "package.zip")
	err = kics_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kics_bin_url := "https://github.com/Checkmarx/kics/archive/refs/tags/v2.1.2.bin"
	kics_cmd_bin := exec.Command("curl", "-L", kics_bin_url, "-o", "binary.bin")
	err = kics_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kics_src_url := "https://github.com/Checkmarx/kics/archive/refs/tags/v2.1.2.src.tar.gz"
	kics_cmd_src := exec.Command("curl", "-L", kics_src_url, "-o", "source.tar.gz")
	err = kics_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kics_cmd_direct := exec.Command("./binary")
	err = kics_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
