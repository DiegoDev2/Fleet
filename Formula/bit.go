package main

// Bit - Distributed Code Component Manager
// Homepage: https://bit.dev

import (
	"fmt"
	
	"os/exec"
)

func installBit() {
	// Método 1: Descargar y extraer .tar.gz
	bit_tar_url := "https://registry.npmjs.org/bit-bin/-/bit-bin-14.8.8.tgz"
	bit_cmd_tar := exec.Command("curl", "-L", bit_tar_url, "-o", "package.tar.gz")
	err := bit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bit_zip_url := "https://registry.npmjs.org/bit-bin/-/bit-bin-14.8.8.tgz"
	bit_cmd_zip := exec.Command("curl", "-L", bit_zip_url, "-o", "package.zip")
	err = bit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bit_bin_url := "https://registry.npmjs.org/bit-bin/-/bit-bin-14.8.8.tgz"
	bit_cmd_bin := exec.Command("curl", "-L", bit_bin_url, "-o", "binary.bin")
	err = bit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bit_src_url := "https://registry.npmjs.org/bit-bin/-/bit-bin-14.8.8.tgz"
	bit_cmd_src := exec.Command("curl", "-L", bit_src_url, "-o", "source.tar.gz")
	err = bit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bit_cmd_direct := exec.Command("./binary")
	err = bit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: terminal-notifier")
	exec.Command("latte", "install", "terminal-notifier").Run()
}
