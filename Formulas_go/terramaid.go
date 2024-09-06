package main

// Terramaid - Utility for generating Mermaid diagrams from Terraform configurations
// Homepage: https://github.com/RoseSecurity/Terramaid

import (
	"fmt"
	
	"os/exec"
)

func installTerramaid() {
	// Método 1: Descargar y extraer .tar.gz
	terramaid_tar_url := "https://github.com/RoseSecurity/Terramaid/archive/refs/tags/v1.12.0.tar.gz"
	terramaid_cmd_tar := exec.Command("curl", "-L", terramaid_tar_url, "-o", "package.tar.gz")
	err := terramaid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terramaid_zip_url := "https://github.com/RoseSecurity/Terramaid/archive/refs/tags/v1.12.0.zip"
	terramaid_cmd_zip := exec.Command("curl", "-L", terramaid_zip_url, "-o", "package.zip")
	err = terramaid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terramaid_bin_url := "https://github.com/RoseSecurity/Terramaid/archive/refs/tags/v1.12.0.bin"
	terramaid_cmd_bin := exec.Command("curl", "-L", terramaid_bin_url, "-o", "binary.bin")
	err = terramaid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terramaid_src_url := "https://github.com/RoseSecurity/Terramaid/archive/refs/tags/v1.12.0.src.tar.gz"
	terramaid_cmd_src := exec.Command("curl", "-L", terramaid_src_url, "-o", "source.tar.gz")
	err = terramaid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terramaid_cmd_direct := exec.Command("./binary")
	err = terramaid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
