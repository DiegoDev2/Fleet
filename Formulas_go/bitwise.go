package main

// Bitwise - Terminal based bit manipulator in ncurses
// Homepage: https://github.com/mellowcandle/bitwise

import (
	"fmt"
	
	"os/exec"
)

func installBitwise() {
	// Método 1: Descargar y extraer .tar.gz
	bitwise_tar_url := "https://github.com/mellowcandle/bitwise/releases/download/v0.50/bitwise-v0.50.tar.gz"
	bitwise_cmd_tar := exec.Command("curl", "-L", bitwise_tar_url, "-o", "package.tar.gz")
	err := bitwise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bitwise_zip_url := "https://github.com/mellowcandle/bitwise/releases/download/v0.50/bitwise-v0.50.zip"
	bitwise_cmd_zip := exec.Command("curl", "-L", bitwise_zip_url, "-o", "package.zip")
	err = bitwise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bitwise_bin_url := "https://github.com/mellowcandle/bitwise/releases/download/v0.50/bitwise-v0.50.bin"
	bitwise_cmd_bin := exec.Command("curl", "-L", bitwise_bin_url, "-o", "binary.bin")
	err = bitwise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bitwise_src_url := "https://github.com/mellowcandle/bitwise/releases/download/v0.50/bitwise-v0.50.src.tar.gz"
	bitwise_cmd_src := exec.Command("curl", "-L", bitwise_src_url, "-o", "source.tar.gz")
	err = bitwise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bitwise_cmd_direct := exec.Command("./binary")
	err = bitwise_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
