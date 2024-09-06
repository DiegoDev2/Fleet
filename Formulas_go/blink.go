package main

// Blink - Tiniest x86-64-linux emulator
// Homepage: https://github.com/jart/blink

import (
	"fmt"
	
	"os/exec"
)

func installBlink() {
	// Método 1: Descargar y extraer .tar.gz
	blink_tar_url := "https://github.com/jart/blink/releases/download/1.1.0/blink-1.1.0.tar.gz"
	blink_cmd_tar := exec.Command("curl", "-L", blink_tar_url, "-o", "package.tar.gz")
	err := blink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blink_zip_url := "https://github.com/jart/blink/releases/download/1.1.0/blink-1.1.0.zip"
	blink_cmd_zip := exec.Command("curl", "-L", blink_zip_url, "-o", "package.zip")
	err = blink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blink_bin_url := "https://github.com/jart/blink/releases/download/1.1.0/blink-1.1.0.bin"
	blink_cmd_bin := exec.Command("curl", "-L", blink_bin_url, "-o", "binary.bin")
	err = blink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blink_src_url := "https://github.com/jart/blink/releases/download/1.1.0/blink-1.1.0.src.tar.gz"
	blink_cmd_src := exec.Command("curl", "-L", blink_src_url, "-o", "source.tar.gz")
	err = blink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blink_cmd_direct := exec.Command("./binary")
	err = blink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: make")
exec.Command("latte", "install", "make")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
