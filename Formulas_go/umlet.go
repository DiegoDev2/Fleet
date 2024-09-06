package main

// Umlet - This UML tool aimed at providing a fast way of creating UML diagrams
// Homepage: https://www.umlet.com/

import (
	"fmt"
	
	"os/exec"
)

func installUmlet() {
	// Método 1: Descargar y extraer .tar.gz
	umlet_tar_url := "https://www.umlet.com/download/umlet_15_1/umlet-standalone-15.1.zip"
	umlet_cmd_tar := exec.Command("curl", "-L", umlet_tar_url, "-o", "package.tar.gz")
	err := umlet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	umlet_zip_url := "https://www.umlet.com/download/umlet_15_1/umlet-standalone-15.1.zip"
	umlet_cmd_zip := exec.Command("curl", "-L", umlet_zip_url, "-o", "package.zip")
	err = umlet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	umlet_bin_url := "https://www.umlet.com/download/umlet_15_1/umlet-standalone-15.1.zip"
	umlet_cmd_bin := exec.Command("curl", "-L", umlet_bin_url, "-o", "binary.bin")
	err = umlet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	umlet_src_url := "https://www.umlet.com/download/umlet_15_1/umlet-standalone-15.1.zip"
	umlet_cmd_src := exec.Command("curl", "-L", umlet_src_url, "-o", "source.tar.gz")
	err = umlet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	umlet_cmd_direct := exec.Command("./binary")
	err = umlet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
