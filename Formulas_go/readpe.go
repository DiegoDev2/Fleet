package main

// Readpe - PE analysis toolkit
// Homepage: https://github.com/mentebinaria/readpe

import (
	"fmt"
	
	"os/exec"
)

func installReadpe() {
	// Método 1: Descargar y extraer .tar.gz
	readpe_tar_url := "https://github.com/mentebinaria/readpe/archive/refs/tags/v0.84.tar.gz"
	readpe_cmd_tar := exec.Command("curl", "-L", readpe_tar_url, "-o", "package.tar.gz")
	err := readpe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	readpe_zip_url := "https://github.com/mentebinaria/readpe/archive/refs/tags/v0.84.zip"
	readpe_cmd_zip := exec.Command("curl", "-L", readpe_zip_url, "-o", "package.zip")
	err = readpe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	readpe_bin_url := "https://github.com/mentebinaria/readpe/archive/refs/tags/v0.84.bin"
	readpe_cmd_bin := exec.Command("curl", "-L", readpe_bin_url, "-o", "binary.bin")
	err = readpe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	readpe_src_url := "https://github.com/mentebinaria/readpe/archive/refs/tags/v0.84.src.tar.gz"
	readpe_cmd_src := exec.Command("curl", "-L", readpe_src_url, "-o", "source.tar.gz")
	err = readpe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	readpe_cmd_direct := exec.Command("./binary")
	err = readpe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
