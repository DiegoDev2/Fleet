package main

// Mlkit - Compiler for the Standard ML programming language
// Homepage: https://melsman.github.io/mlkit

import (
	"fmt"
	
	"os/exec"
)

func installMlkit() {
	// Método 1: Descargar y extraer .tar.gz
	mlkit_tar_url := "https://github.com/melsman/mlkit/archive/refs/tags/v4.7.11.tar.gz"
	mlkit_cmd_tar := exec.Command("curl", "-L", mlkit_tar_url, "-o", "package.tar.gz")
	err := mlkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mlkit_zip_url := "https://github.com/melsman/mlkit/archive/refs/tags/v4.7.11.zip"
	mlkit_cmd_zip := exec.Command("curl", "-L", mlkit_zip_url, "-o", "package.zip")
	err = mlkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mlkit_bin_url := "https://github.com/melsman/mlkit/archive/refs/tags/v4.7.11.bin"
	mlkit_cmd_bin := exec.Command("curl", "-L", mlkit_bin_url, "-o", "binary.bin")
	err = mlkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mlkit_src_url := "https://github.com/melsman/mlkit/archive/refs/tags/v4.7.11.src.tar.gz"
	mlkit_cmd_src := exec.Command("curl", "-L", mlkit_src_url, "-o", "source.tar.gz")
	err = mlkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mlkit_cmd_direct := exec.Command("./binary")
	err = mlkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: mlton")
exec.Command("latte", "install", "mlton")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
