package main

// Helib - Implementation of homomorphic encryption
// Homepage: https://github.com/homenc/HElib

import (
	"fmt"
	
	"os/exec"
)

func installHelib() {
	// Método 1: Descargar y extraer .tar.gz
	helib_tar_url := "https://github.com/homenc/HElib/archive/refs/tags/v2.3.0.tar.gz"
	helib_cmd_tar := exec.Command("curl", "-L", helib_tar_url, "-o", "package.tar.gz")
	err := helib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	helib_zip_url := "https://github.com/homenc/HElib/archive/refs/tags/v2.3.0.zip"
	helib_cmd_zip := exec.Command("curl", "-L", helib_zip_url, "-o", "package.zip")
	err = helib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	helib_bin_url := "https://github.com/homenc/HElib/archive/refs/tags/v2.3.0.bin"
	helib_cmd_bin := exec.Command("curl", "-L", helib_bin_url, "-o", "binary.bin")
	err = helib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	helib_src_url := "https://github.com/homenc/HElib/archive/refs/tags/v2.3.0.src.tar.gz"
	helib_cmd_src := exec.Command("curl", "-L", helib_src_url, "-o", "source.tar.gz")
	err = helib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	helib_cmd_direct := exec.Command("./binary")
	err = helib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: bats-core")
	exec.Command("latte", "install", "bats-core").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: ntl")
	exec.Command("latte", "install", "ntl").Run()
}
