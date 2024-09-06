package main

// Boolector - SMT solver for fixed-size bit-vectors
// Homepage: https://boolector.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installBoolector() {
	// Método 1: Descargar y extraer .tar.gz
	boolector_tar_url := "https://github.com/Boolector/boolector/archive/refs/tags/3.2.4.tar.gz"
	boolector_cmd_tar := exec.Command("curl", "-L", boolector_tar_url, "-o", "package.tar.gz")
	err := boolector_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boolector_zip_url := "https://github.com/Boolector/boolector/archive/refs/tags/3.2.4.zip"
	boolector_cmd_zip := exec.Command("curl", "-L", boolector_zip_url, "-o", "package.zip")
	err = boolector_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boolector_bin_url := "https://github.com/Boolector/boolector/archive/refs/tags/3.2.4.bin"
	boolector_cmd_bin := exec.Command("curl", "-L", boolector_bin_url, "-o", "binary.bin")
	err = boolector_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boolector_src_url := "https://github.com/Boolector/boolector/archive/refs/tags/3.2.4.src.tar.gz"
	boolector_cmd_src := exec.Command("curl", "-L", boolector_src_url, "-o", "source.tar.gz")
	err = boolector_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boolector_cmd_direct := exec.Command("./binary")
	err = boolector_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
