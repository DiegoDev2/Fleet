package main

// Reproc - Cross-platform (C99/C++11) process library
// Homepage: https://github.com/DaanDeMeyer/reproc

import (
	"fmt"
	
	"os/exec"
)

func installReproc() {
	// Método 1: Descargar y extraer .tar.gz
	reproc_tar_url := "https://github.com/DaanDeMeyer/reproc/archive/refs/tags/v14.2.5.tar.gz"
	reproc_cmd_tar := exec.Command("curl", "-L", reproc_tar_url, "-o", "package.tar.gz")
	err := reproc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reproc_zip_url := "https://github.com/DaanDeMeyer/reproc/archive/refs/tags/v14.2.5.zip"
	reproc_cmd_zip := exec.Command("curl", "-L", reproc_zip_url, "-o", "package.zip")
	err = reproc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reproc_bin_url := "https://github.com/DaanDeMeyer/reproc/archive/refs/tags/v14.2.5.bin"
	reproc_cmd_bin := exec.Command("curl", "-L", reproc_bin_url, "-o", "binary.bin")
	err = reproc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reproc_src_url := "https://github.com/DaanDeMeyer/reproc/archive/refs/tags/v14.2.5.src.tar.gz"
	reproc_cmd_src := exec.Command("curl", "-L", reproc_src_url, "-o", "source.tar.gz")
	err = reproc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reproc_cmd_direct := exec.Command("./binary")
	err = reproc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
