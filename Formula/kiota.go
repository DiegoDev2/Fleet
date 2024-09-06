package main

// Kiota - OpenAPI based HTTP Client code generator
// Homepage: https://aka.ms/kiota/docs

import (
	"fmt"
	
	"os/exec"
)

func installKiota() {
	// Método 1: Descargar y extraer .tar.gz
	kiota_tar_url := "https://github.com/microsoft/kiota/archive/refs/tags/v1.11.1.tar.gz"
	kiota_cmd_tar := exec.Command("curl", "-L", kiota_tar_url, "-o", "package.tar.gz")
	err := kiota_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kiota_zip_url := "https://github.com/microsoft/kiota/archive/refs/tags/v1.11.1.zip"
	kiota_cmd_zip := exec.Command("curl", "-L", kiota_zip_url, "-o", "package.zip")
	err = kiota_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kiota_bin_url := "https://github.com/microsoft/kiota/archive/refs/tags/v1.11.1.bin"
	kiota_cmd_bin := exec.Command("curl", "-L", kiota_bin_url, "-o", "binary.bin")
	err = kiota_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kiota_src_url := "https://github.com/microsoft/kiota/archive/refs/tags/v1.11.1.src.tar.gz"
	kiota_cmd_src := exec.Command("curl", "-L", kiota_src_url, "-o", "source.tar.gz")
	err = kiota_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kiota_cmd_direct := exec.Command("./binary")
	err = kiota_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dotnet")
	exec.Command("latte", "install", "dotnet").Run()
}
