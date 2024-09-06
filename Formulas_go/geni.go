package main

// Geni - Standalone database migration tool
// Homepage: https://github.com/emilpriver/geni

import (
	"fmt"
	
	"os/exec"
)

func installGeni() {
	// Método 1: Descargar y extraer .tar.gz
	geni_tar_url := "https://github.com/emilpriver/geni/archive/refs/tags/v1.0.14.tar.gz"
	geni_cmd_tar := exec.Command("curl", "-L", geni_tar_url, "-o", "package.tar.gz")
	err := geni_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geni_zip_url := "https://github.com/emilpriver/geni/archive/refs/tags/v1.0.14.zip"
	geni_cmd_zip := exec.Command("curl", "-L", geni_zip_url, "-o", "package.zip")
	err = geni_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geni_bin_url := "https://github.com/emilpriver/geni/archive/refs/tags/v1.0.14.bin"
	geni_cmd_bin := exec.Command("curl", "-L", geni_bin_url, "-o", "binary.bin")
	err = geni_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geni_src_url := "https://github.com/emilpriver/geni/archive/refs/tags/v1.0.14.src.tar.gz"
	geni_cmd_src := exec.Command("curl", "-L", geni_src_url, "-o", "source.tar.gz")
	err = geni_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geni_cmd_direct := exec.Command("./binary")
	err = geni_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
