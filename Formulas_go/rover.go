package main

// Rover - CLI for managing and maintaining data graphs with Apollo Studio
// Homepage: https://www.apollographql.com/docs/rover/

import (
	"fmt"
	
	"os/exec"
)

func installRover() {
	// Método 1: Descargar y extraer .tar.gz
	rover_tar_url := "https://github.com/apollographql/rover/archive/refs/tags/v0.26.1.tar.gz"
	rover_cmd_tar := exec.Command("curl", "-L", rover_tar_url, "-o", "package.tar.gz")
	err := rover_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rover_zip_url := "https://github.com/apollographql/rover/archive/refs/tags/v0.26.1.zip"
	rover_cmd_zip := exec.Command("curl", "-L", rover_zip_url, "-o", "package.zip")
	err = rover_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rover_bin_url := "https://github.com/apollographql/rover/archive/refs/tags/v0.26.1.bin"
	rover_cmd_bin := exec.Command("curl", "-L", rover_bin_url, "-o", "binary.bin")
	err = rover_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rover_src_url := "https://github.com/apollographql/rover/archive/refs/tags/v0.26.1.src.tar.gz"
	rover_cmd_src := exec.Command("curl", "-L", rover_src_url, "-o", "source.tar.gz")
	err = rover_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rover_cmd_direct := exec.Command("./binary")
	err = rover_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
