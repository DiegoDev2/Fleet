package main

// Inframap - Read your tfstate or HCL to generate a graph
// Homepage: https://github.com/cycloidio/inframap

import (
	"fmt"
	
	"os/exec"
)

func installInframap() {
	// Método 1: Descargar y extraer .tar.gz
	inframap_tar_url := "https://github.com/cycloidio/inframap/archive/refs/tags/v0.7.0.tar.gz"
	inframap_cmd_tar := exec.Command("curl", "-L", inframap_tar_url, "-o", "package.tar.gz")
	err := inframap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inframap_zip_url := "https://github.com/cycloidio/inframap/archive/refs/tags/v0.7.0.zip"
	inframap_cmd_zip := exec.Command("curl", "-L", inframap_zip_url, "-o", "package.zip")
	err = inframap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inframap_bin_url := "https://github.com/cycloidio/inframap/archive/refs/tags/v0.7.0.bin"
	inframap_cmd_bin := exec.Command("curl", "-L", inframap_bin_url, "-o", "binary.bin")
	err = inframap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inframap_src_url := "https://github.com/cycloidio/inframap/archive/refs/tags/v0.7.0.src.tar.gz"
	inframap_cmd_src := exec.Command("curl", "-L", inframap_src_url, "-o", "source.tar.gz")
	err = inframap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inframap_cmd_direct := exec.Command("./binary")
	err = inframap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
