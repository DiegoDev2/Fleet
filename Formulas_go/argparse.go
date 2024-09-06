package main

// Argparse - Argument Parser for Modern C++
// Homepage: https://github.com/p-ranav/argparse

import (
	"fmt"
	
	"os/exec"
)

func installArgparse() {
	// Método 1: Descargar y extraer .tar.gz
	argparse_tar_url := "https://github.com/p-ranav/argparse/archive/refs/tags/v3.1.tar.gz"
	argparse_cmd_tar := exec.Command("curl", "-L", argparse_tar_url, "-o", "package.tar.gz")
	err := argparse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argparse_zip_url := "https://github.com/p-ranav/argparse/archive/refs/tags/v3.1.zip"
	argparse_cmd_zip := exec.Command("curl", "-L", argparse_zip_url, "-o", "package.zip")
	err = argparse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argparse_bin_url := "https://github.com/p-ranav/argparse/archive/refs/tags/v3.1.bin"
	argparse_cmd_bin := exec.Command("curl", "-L", argparse_bin_url, "-o", "binary.bin")
	err = argparse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argparse_src_url := "https://github.com/p-ranav/argparse/archive/refs/tags/v3.1.src.tar.gz"
	argparse_cmd_src := exec.Command("curl", "-L", argparse_src_url, "-o", "source.tar.gz")
	err = argparse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argparse_cmd_direct := exec.Command("./binary")
	err = argparse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
