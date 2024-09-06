package main

// TerragruntAtlantisConfig - Generate Atlantis config for Terragrunt projects
// Homepage: https://github.com/transcend-io/terragrunt-atlantis-config

import (
	"fmt"
	
	"os/exec"
)

func installTerragruntAtlantisConfig() {
	// Método 1: Descargar y extraer .tar.gz
	terragruntatlantisconfig_tar_url := "https://github.com/transcend-io/terragrunt-atlantis-config/archive/refs/tags/v1.18.0.tar.gz"
	terragruntatlantisconfig_cmd_tar := exec.Command("curl", "-L", terragruntatlantisconfig_tar_url, "-o", "package.tar.gz")
	err := terragruntatlantisconfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terragruntatlantisconfig_zip_url := "https://github.com/transcend-io/terragrunt-atlantis-config/archive/refs/tags/v1.18.0.zip"
	terragruntatlantisconfig_cmd_zip := exec.Command("curl", "-L", terragruntatlantisconfig_zip_url, "-o", "package.zip")
	err = terragruntatlantisconfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terragruntatlantisconfig_bin_url := "https://github.com/transcend-io/terragrunt-atlantis-config/archive/refs/tags/v1.18.0.bin"
	terragruntatlantisconfig_cmd_bin := exec.Command("curl", "-L", terragruntatlantisconfig_bin_url, "-o", "binary.bin")
	err = terragruntatlantisconfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terragruntatlantisconfig_src_url := "https://github.com/transcend-io/terragrunt-atlantis-config/archive/refs/tags/v1.18.0.src.tar.gz"
	terragruntatlantisconfig_cmd_src := exec.Command("curl", "-L", terragruntatlantisconfig_src_url, "-o", "source.tar.gz")
	err = terragruntatlantisconfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terragruntatlantisconfig_cmd_direct := exec.Command("./binary")
	err = terragruntatlantisconfig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
