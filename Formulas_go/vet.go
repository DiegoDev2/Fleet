package main

// Vet - Policy driven vetting of open source dependencies
// Homepage: https://github.com/safedep/vet

import (
	"fmt"
	
	"os/exec"
)

func installVet() {
	// Método 1: Descargar y extraer .tar.gz
	vet_tar_url := "https://github.com/safedep/vet/archive/refs/tags/v1.6.1.tar.gz"
	vet_cmd_tar := exec.Command("curl", "-L", vet_tar_url, "-o", "package.tar.gz")
	err := vet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vet_zip_url := "https://github.com/safedep/vet/archive/refs/tags/v1.6.1.zip"
	vet_cmd_zip := exec.Command("curl", "-L", vet_zip_url, "-o", "package.zip")
	err = vet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vet_bin_url := "https://github.com/safedep/vet/archive/refs/tags/v1.6.1.bin"
	vet_cmd_bin := exec.Command("curl", "-L", vet_bin_url, "-o", "binary.bin")
	err = vet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vet_src_url := "https://github.com/safedep/vet/archive/refs/tags/v1.6.1.src.tar.gz"
	vet_cmd_src := exec.Command("curl", "-L", vet_src_url, "-o", "source.tar.gz")
	err = vet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vet_cmd_direct := exec.Command("./binary")
	err = vet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
