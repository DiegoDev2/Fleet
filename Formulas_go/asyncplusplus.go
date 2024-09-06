package main

// Asyncplusplus - Concurrency framework for C++11
// Homepage: https://github.com/Amanieu/asyncplusplus

import (
	"fmt"
	
	"os/exec"
)

func installAsyncplusplus() {
	// Método 1: Descargar y extraer .tar.gz
	asyncplusplus_tar_url := "https://github.com/Amanieu/asyncplusplus/archive/refs/tags/v1.1.tar.gz"
	asyncplusplus_cmd_tar := exec.Command("curl", "-L", asyncplusplus_tar_url, "-o", "package.tar.gz")
	err := asyncplusplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asyncplusplus_zip_url := "https://github.com/Amanieu/asyncplusplus/archive/refs/tags/v1.1.zip"
	asyncplusplus_cmd_zip := exec.Command("curl", "-L", asyncplusplus_zip_url, "-o", "package.zip")
	err = asyncplusplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asyncplusplus_bin_url := "https://github.com/Amanieu/asyncplusplus/archive/refs/tags/v1.1.bin"
	asyncplusplus_cmd_bin := exec.Command("curl", "-L", asyncplusplus_bin_url, "-o", "binary.bin")
	err = asyncplusplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asyncplusplus_src_url := "https://github.com/Amanieu/asyncplusplus/archive/refs/tags/v1.1.src.tar.gz"
	asyncplusplus_cmd_src := exec.Command("curl", "-L", asyncplusplus_src_url, "-o", "source.tar.gz")
	err = asyncplusplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asyncplusplus_cmd_direct := exec.Command("./binary")
	err = asyncplusplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
