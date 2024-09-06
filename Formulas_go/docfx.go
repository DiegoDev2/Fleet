package main

// Docfx - Tools for building and publishing API documentation for .NET projects
// Homepage: https://dotnet.github.io/docfx/

import (
	"fmt"
	
	"os/exec"
)

func installDocfx() {
	// Método 1: Descargar y extraer .tar.gz
	docfx_tar_url := "https://github.com/dotnet/docfx/archive/refs/tags/v2.77.0.tar.gz"
	docfx_cmd_tar := exec.Command("curl", "-L", docfx_tar_url, "-o", "package.tar.gz")
	err := docfx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	docfx_zip_url := "https://github.com/dotnet/docfx/archive/refs/tags/v2.77.0.zip"
	docfx_cmd_zip := exec.Command("curl", "-L", docfx_zip_url, "-o", "package.zip")
	err = docfx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	docfx_bin_url := "https://github.com/dotnet/docfx/archive/refs/tags/v2.77.0.bin"
	docfx_cmd_bin := exec.Command("curl", "-L", docfx_bin_url, "-o", "binary.bin")
	err = docfx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	docfx_src_url := "https://github.com/dotnet/docfx/archive/refs/tags/v2.77.0.src.tar.gz"
	docfx_cmd_src := exec.Command("curl", "-L", docfx_src_url, "-o", "source.tar.gz")
	err = docfx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	docfx_cmd_direct := exec.Command("./binary")
	err = docfx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dotnet")
exec.Command("latte", "install", "dotnet")
}
