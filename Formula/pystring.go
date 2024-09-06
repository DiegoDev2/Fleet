package main

// Pystring - Collection of C++ functions for the interface of Python's string class methods
// Homepage: https://github.com/imageworks/pystring

import (
	"fmt"
	
	"os/exec"
)

func installPystring() {
	// Método 1: Descargar y extraer .tar.gz
	pystring_tar_url := "https://github.com/imageworks/pystring/archive/refs/tags/v1.1.4.tar.gz"
	pystring_cmd_tar := exec.Command("curl", "-L", pystring_tar_url, "-o", "package.tar.gz")
	err := pystring_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pystring_zip_url := "https://github.com/imageworks/pystring/archive/refs/tags/v1.1.4.zip"
	pystring_cmd_zip := exec.Command("curl", "-L", pystring_zip_url, "-o", "package.zip")
	err = pystring_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pystring_bin_url := "https://github.com/imageworks/pystring/archive/refs/tags/v1.1.4.bin"
	pystring_cmd_bin := exec.Command("curl", "-L", pystring_bin_url, "-o", "binary.bin")
	err = pystring_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pystring_src_url := "https://github.com/imageworks/pystring/archive/refs/tags/v1.1.4.src.tar.gz"
	pystring_cmd_src := exec.Command("curl", "-L", pystring_src_url, "-o", "source.tar.gz")
	err = pystring_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pystring_cmd_direct := exec.Command("./binary")
	err = pystring_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
