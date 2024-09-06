package main

// Z3 - High-performance theorem prover
// Homepage: https://github.com/Z3Prover/z3

import (
	"fmt"
	
	"os/exec"
)

func installZ3() {
	// Método 1: Descargar y extraer .tar.gz
	z3_tar_url := "https://github.com/Z3Prover/z3/archive/refs/tags/z3-4.13.0.tar.gz"
	z3_cmd_tar := exec.Command("curl", "-L", z3_tar_url, "-o", "package.tar.gz")
	err := z3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	z3_zip_url := "https://github.com/Z3Prover/z3/archive/refs/tags/z3-4.13.0.zip"
	z3_cmd_zip := exec.Command("curl", "-L", z3_zip_url, "-o", "package.zip")
	err = z3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	z3_bin_url := "https://github.com/Z3Prover/z3/archive/refs/tags/z3-4.13.0.bin"
	z3_cmd_bin := exec.Command("curl", "-L", z3_bin_url, "-o", "binary.bin")
	err = z3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	z3_src_url := "https://github.com/Z3Prover/z3/archive/refs/tags/z3-4.13.0.src.tar.gz"
	z3_cmd_src := exec.Command("curl", "-L", z3_src_url, "-o", "source.tar.gz")
	err = z3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	z3_cmd_direct := exec.Command("./binary")
	err = z3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
