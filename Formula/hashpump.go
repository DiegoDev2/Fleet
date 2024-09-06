package main

// Hashpump - Tool to exploit hash length extension attack
// Homepage: https://github.com/bwall/HashPump

import (
	"fmt"
	
	"os/exec"
)

func installHashpump() {
	// Método 1: Descargar y extraer .tar.gz
	hashpump_tar_url := "https://github.com/bwall/HashPump/archive/refs/tags/v1.2.0.tar.gz"
	hashpump_cmd_tar := exec.Command("curl", "-L", hashpump_tar_url, "-o", "package.tar.gz")
	err := hashpump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hashpump_zip_url := "https://github.com/bwall/HashPump/archive/refs/tags/v1.2.0.zip"
	hashpump_cmd_zip := exec.Command("curl", "-L", hashpump_zip_url, "-o", "package.zip")
	err = hashpump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hashpump_bin_url := "https://github.com/bwall/HashPump/archive/refs/tags/v1.2.0.bin"
	hashpump_cmd_bin := exec.Command("curl", "-L", hashpump_bin_url, "-o", "binary.bin")
	err = hashpump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hashpump_src_url := "https://github.com/bwall/HashPump/archive/refs/tags/v1.2.0.src.tar.gz"
	hashpump_cmd_src := exec.Command("curl", "-L", hashpump_src_url, "-o", "source.tar.gz")
	err = hashpump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hashpump_cmd_direct := exec.Command("./binary")
	err = hashpump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
}
