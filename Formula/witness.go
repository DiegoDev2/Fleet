package main

// Witness - Automates, normalizes, and verifies software artifact provenance
// Homepage: https://witness.dev

import (
	"fmt"
	
	"os/exec"
)

func installWitness() {
	// Método 1: Descargar y extraer .tar.gz
	witness_tar_url := "https://github.com/in-toto/witness/archive/refs/tags/v0.6.0.tar.gz"
	witness_cmd_tar := exec.Command("curl", "-L", witness_tar_url, "-o", "package.tar.gz")
	err := witness_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	witness_zip_url := "https://github.com/in-toto/witness/archive/refs/tags/v0.6.0.zip"
	witness_cmd_zip := exec.Command("curl", "-L", witness_zip_url, "-o", "package.zip")
	err = witness_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	witness_bin_url := "https://github.com/in-toto/witness/archive/refs/tags/v0.6.0.bin"
	witness_cmd_bin := exec.Command("curl", "-L", witness_bin_url, "-o", "binary.bin")
	err = witness_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	witness_src_url := "https://github.com/in-toto/witness/archive/refs/tags/v0.6.0.src.tar.gz"
	witness_cmd_src := exec.Command("curl", "-L", witness_src_url, "-o", "source.tar.gz")
	err = witness_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	witness_cmd_direct := exec.Command("./binary")
	err = witness_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
