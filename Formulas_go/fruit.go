package main

// Fruit - Dependency injection framework for C++
// Homepage: https://github.com/google/fruit/wiki

import (
	"fmt"
	
	"os/exec"
)

func installFruit() {
	// Método 1: Descargar y extraer .tar.gz
	fruit_tar_url := "https://github.com/google/fruit/archive/refs/tags/v3.7.1.tar.gz"
	fruit_cmd_tar := exec.Command("curl", "-L", fruit_tar_url, "-o", "package.tar.gz")
	err := fruit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fruit_zip_url := "https://github.com/google/fruit/archive/refs/tags/v3.7.1.zip"
	fruit_cmd_zip := exec.Command("curl", "-L", fruit_zip_url, "-o", "package.zip")
	err = fruit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fruit_bin_url := "https://github.com/google/fruit/archive/refs/tags/v3.7.1.bin"
	fruit_cmd_bin := exec.Command("curl", "-L", fruit_bin_url, "-o", "binary.bin")
	err = fruit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fruit_src_url := "https://github.com/google/fruit/archive/refs/tags/v3.7.1.src.tar.gz"
	fruit_cmd_src := exec.Command("curl", "-L", fruit_src_url, "-o", "source.tar.gz")
	err = fruit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fruit_cmd_direct := exec.Command("./binary")
	err = fruit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
