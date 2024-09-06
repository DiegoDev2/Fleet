package main

// Graphqxl - Language for creating big and scalable GraphQL server-side schemas
// Homepage: https://gabotechs.github.io/graphqxl

import (
	"fmt"
	
	"os/exec"
)

func installGraphqxl() {
	// Método 1: Descargar y extraer .tar.gz
	graphqxl_tar_url := "https://github.com/gabotechs/graphqxl/archive/refs/tags/v0.40.2.tar.gz"
	graphqxl_cmd_tar := exec.Command("curl", "-L", graphqxl_tar_url, "-o", "package.tar.gz")
	err := graphqxl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphqxl_zip_url := "https://github.com/gabotechs/graphqxl/archive/refs/tags/v0.40.2.zip"
	graphqxl_cmd_zip := exec.Command("curl", "-L", graphqxl_zip_url, "-o", "package.zip")
	err = graphqxl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphqxl_bin_url := "https://github.com/gabotechs/graphqxl/archive/refs/tags/v0.40.2.bin"
	graphqxl_cmd_bin := exec.Command("curl", "-L", graphqxl_bin_url, "-o", "binary.bin")
	err = graphqxl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphqxl_src_url := "https://github.com/gabotechs/graphqxl/archive/refs/tags/v0.40.2.src.tar.gz"
	graphqxl_cmd_src := exec.Command("curl", "-L", graphqxl_src_url, "-o", "source.tar.gz")
	err = graphqxl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphqxl_cmd_direct := exec.Command("./binary")
	err = graphqxl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
