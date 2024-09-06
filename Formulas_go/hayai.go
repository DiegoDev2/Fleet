package main

// Hayai - C++ benchmarking framework inspired by the googletest framework
// Homepage: https://bruun.co/2012/02/07/easy-cpp-benchmarking

import (
	"fmt"
	
	"os/exec"
)

func installHayai() {
	// Método 1: Descargar y extraer .tar.gz
	hayai_tar_url := "https://github.com/nickbruun/hayai/archive/refs/tags/v1.0.2.tar.gz"
	hayai_cmd_tar := exec.Command("curl", "-L", hayai_tar_url, "-o", "package.tar.gz")
	err := hayai_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hayai_zip_url := "https://github.com/nickbruun/hayai/archive/refs/tags/v1.0.2.zip"
	hayai_cmd_zip := exec.Command("curl", "-L", hayai_zip_url, "-o", "package.zip")
	err = hayai_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hayai_bin_url := "https://github.com/nickbruun/hayai/archive/refs/tags/v1.0.2.bin"
	hayai_cmd_bin := exec.Command("curl", "-L", hayai_bin_url, "-o", "binary.bin")
	err = hayai_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hayai_src_url := "https://github.com/nickbruun/hayai/archive/refs/tags/v1.0.2.src.tar.gz"
	hayai_cmd_src := exec.Command("curl", "-L", hayai_src_url, "-o", "source.tar.gz")
	err = hayai_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hayai_cmd_direct := exec.Command("./binary")
	err = hayai_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
