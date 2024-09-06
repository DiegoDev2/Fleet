package main

// Orcania - Potluck with different functions for different purposes in C
// Homepage: https://babelouest.github.io/orcania/

import (
	"fmt"
	
	"os/exec"
)

func installOrcania() {
	// Método 1: Descargar y extraer .tar.gz
	orcania_tar_url := "https://github.com/babelouest/orcania/archive/refs/tags/v2.3.3.tar.gz"
	orcania_cmd_tar := exec.Command("curl", "-L", orcania_tar_url, "-o", "package.tar.gz")
	err := orcania_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orcania_zip_url := "https://github.com/babelouest/orcania/archive/refs/tags/v2.3.3.zip"
	orcania_cmd_zip := exec.Command("curl", "-L", orcania_zip_url, "-o", "package.zip")
	err = orcania_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orcania_bin_url := "https://github.com/babelouest/orcania/archive/refs/tags/v2.3.3.bin"
	orcania_cmd_bin := exec.Command("curl", "-L", orcania_bin_url, "-o", "binary.bin")
	err = orcania_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orcania_src_url := "https://github.com/babelouest/orcania/archive/refs/tags/v2.3.3.src.tar.gz"
	orcania_cmd_src := exec.Command("curl", "-L", orcania_src_url, "-o", "source.tar.gz")
	err = orcania_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orcania_cmd_direct := exec.Command("./binary")
	err = orcania_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
}
