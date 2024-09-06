package main

// StructurizrCli - Command-line utility for Structurizr
// Homepage: https://structurizr.com

import (
	"fmt"
	
	"os/exec"
)

func installStructurizrCli() {
	// Método 1: Descargar y extraer .tar.gz
	structurizrcli_tar_url := "https://github.com/structurizr/cli/releases/download/v2024.07.03/structurizr-cli.zip"
	structurizrcli_cmd_tar := exec.Command("curl", "-L", structurizrcli_tar_url, "-o", "package.tar.gz")
	err := structurizrcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	structurizrcli_zip_url := "https://github.com/structurizr/cli/releases/download/v2024.07.03/structurizr-cli.zip"
	structurizrcli_cmd_zip := exec.Command("curl", "-L", structurizrcli_zip_url, "-o", "package.zip")
	err = structurizrcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	structurizrcli_bin_url := "https://github.com/structurizr/cli/releases/download/v2024.07.03/structurizr-cli.zip"
	structurizrcli_cmd_bin := exec.Command("curl", "-L", structurizrcli_bin_url, "-o", "binary.bin")
	err = structurizrcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	structurizrcli_src_url := "https://github.com/structurizr/cli/releases/download/v2024.07.03/structurizr-cli.zip"
	structurizrcli_cmd_src := exec.Command("curl", "-L", structurizrcli_src_url, "-o", "source.tar.gz")
	err = structurizrcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	structurizrcli_cmd_direct := exec.Command("./binary")
	err = structurizrcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
