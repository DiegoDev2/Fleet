package main

// Avra - Assembler for the Atmel AVR microcontroller family
// Homepage: https://github.com/Ro5bert/avra

import (
	"fmt"
	
	"os/exec"
)

func installAvra() {
	// Método 1: Descargar y extraer .tar.gz
	avra_tar_url := "https://github.com/Ro5bert/avra/archive/refs/tags/1.4.2.tar.gz"
	avra_cmd_tar := exec.Command("curl", "-L", avra_tar_url, "-o", "package.tar.gz")
	err := avra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avra_zip_url := "https://github.com/Ro5bert/avra/archive/refs/tags/1.4.2.zip"
	avra_cmd_zip := exec.Command("curl", "-L", avra_zip_url, "-o", "package.zip")
	err = avra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avra_bin_url := "https://github.com/Ro5bert/avra/archive/refs/tags/1.4.2.bin"
	avra_cmd_bin := exec.Command("curl", "-L", avra_bin_url, "-o", "binary.bin")
	err = avra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avra_src_url := "https://github.com/Ro5bert/avra/archive/refs/tags/1.4.2.src.tar.gz"
	avra_cmd_src := exec.Command("curl", "-L", avra_src_url, "-o", "source.tar.gz")
	err = avra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avra_cmd_direct := exec.Command("./binary")
	err = avra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
