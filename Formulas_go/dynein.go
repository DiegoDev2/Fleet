package main

// Dynein - DynamoDB CLI
// Homepage: https://github.com/awslabs/dynein

import (
	"fmt"
	
	"os/exec"
)

func installDynein() {
	// Método 1: Descargar y extraer .tar.gz
	dynein_tar_url := "https://github.com/awslabs/dynein/archive/refs/tags/v0.3.0.tar.gz"
	dynein_cmd_tar := exec.Command("curl", "-L", dynein_tar_url, "-o", "package.tar.gz")
	err := dynein_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dynein_zip_url := "https://github.com/awslabs/dynein/archive/refs/tags/v0.3.0.zip"
	dynein_cmd_zip := exec.Command("curl", "-L", dynein_zip_url, "-o", "package.zip")
	err = dynein_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dynein_bin_url := "https://github.com/awslabs/dynein/archive/refs/tags/v0.3.0.bin"
	dynein_cmd_bin := exec.Command("curl", "-L", dynein_bin_url, "-o", "binary.bin")
	err = dynein_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dynein_src_url := "https://github.com/awslabs/dynein/archive/refs/tags/v0.3.0.src.tar.gz"
	dynein_cmd_src := exec.Command("curl", "-L", dynein_src_url, "-o", "source.tar.gz")
	err = dynein_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dynein_cmd_direct := exec.Command("./binary")
	err = dynein_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
