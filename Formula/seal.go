package main

// Seal - Easy-to-use homomorphic encryption library
// Homepage: https://github.com/microsoft/SEAL

import (
	"fmt"
	
	"os/exec"
)

func installSeal() {
	// Método 1: Descargar y extraer .tar.gz
	seal_tar_url := "https://github.com/microsoft/SEAL/archive/refs/tags/v4.1.2.tar.gz"
	seal_cmd_tar := exec.Command("curl", "-L", seal_tar_url, "-o", "package.tar.gz")
	err := seal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	seal_zip_url := "https://github.com/microsoft/SEAL/archive/refs/tags/v4.1.2.zip"
	seal_cmd_zip := exec.Command("curl", "-L", seal_zip_url, "-o", "package.zip")
	err = seal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	seal_bin_url := "https://github.com/microsoft/SEAL/archive/refs/tags/v4.1.2.bin"
	seal_cmd_bin := exec.Command("curl", "-L", seal_bin_url, "-o", "binary.bin")
	err = seal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	seal_src_url := "https://github.com/microsoft/SEAL/archive/refs/tags/v4.1.2.src.tar.gz"
	seal_cmd_src := exec.Command("curl", "-L", seal_src_url, "-o", "source.tar.gz")
	err = seal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	seal_cmd_direct := exec.Command("./binary")
	err = seal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: cpp-gsl")
	exec.Command("latte", "install", "cpp-gsl").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
