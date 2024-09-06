package main

// Oakc - Portable programming language with a compact intermediate representation
// Homepage: https://github.com/adam-mcdaniel/oakc

import (
	"fmt"
	
	"os/exec"
)

func installOakc() {
	// Método 1: Descargar y extraer .tar.gz
	oakc_tar_url := "https://static.crates.io/crates/oakc/oakc-0.6.1.crate"
	oakc_cmd_tar := exec.Command("curl", "-L", oakc_tar_url, "-o", "package.tar.gz")
	err := oakc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oakc_zip_url := "https://static.crates.io/crates/oakc/oakc-0.6.1.crate"
	oakc_cmd_zip := exec.Command("curl", "-L", oakc_zip_url, "-o", "package.zip")
	err = oakc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oakc_bin_url := "https://static.crates.io/crates/oakc/oakc-0.6.1.crate"
	oakc_cmd_bin := exec.Command("curl", "-L", oakc_bin_url, "-o", "binary.bin")
	err = oakc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oakc_src_url := "https://static.crates.io/crates/oakc/oakc-0.6.1.crate"
	oakc_cmd_src := exec.Command("curl", "-L", oakc_src_url, "-o", "source.tar.gz")
	err = oakc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oakc_cmd_direct := exec.Command("./binary")
	err = oakc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
