package main

// Chamber - CLI for managing secrets through AWS SSM Parameter Store
// Homepage: https://github.com/segmentio/chamber

import (
	"fmt"
	
	"os/exec"
)

func installChamber() {
	// Método 1: Descargar y extraer .tar.gz
	chamber_tar_url := "https://github.com/segmentio/chamber/archive/refs/tags/v3.1.0.tar.gz"
	chamber_cmd_tar := exec.Command("curl", "-L", chamber_tar_url, "-o", "package.tar.gz")
	err := chamber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chamber_zip_url := "https://github.com/segmentio/chamber/archive/refs/tags/v3.1.0.zip"
	chamber_cmd_zip := exec.Command("curl", "-L", chamber_zip_url, "-o", "package.zip")
	err = chamber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chamber_bin_url := "https://github.com/segmentio/chamber/archive/refs/tags/v3.1.0.bin"
	chamber_cmd_bin := exec.Command("curl", "-L", chamber_bin_url, "-o", "binary.bin")
	err = chamber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chamber_src_url := "https://github.com/segmentio/chamber/archive/refs/tags/v3.1.0.src.tar.gz"
	chamber_cmd_src := exec.Command("curl", "-L", chamber_src_url, "-o", "source.tar.gz")
	err = chamber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chamber_cmd_direct := exec.Command("./binary")
	err = chamber_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
