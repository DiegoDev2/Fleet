package main

// Cyme - List system USB buses and devices
// Homepage: https://github.com/tuna-f1sh/cyme

import (
	"fmt"
	
	"os/exec"
)

func installCyme() {
	// Método 1: Descargar y extraer .tar.gz
	cyme_tar_url := "https://github.com/tuna-f1sh/cyme/archive/refs/tags/v1.8.2.tar.gz"
	cyme_cmd_tar := exec.Command("curl", "-L", cyme_tar_url, "-o", "package.tar.gz")
	err := cyme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cyme_zip_url := "https://github.com/tuna-f1sh/cyme/archive/refs/tags/v1.8.2.zip"
	cyme_cmd_zip := exec.Command("curl", "-L", cyme_zip_url, "-o", "package.zip")
	err = cyme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cyme_bin_url := "https://github.com/tuna-f1sh/cyme/archive/refs/tags/v1.8.2.bin"
	cyme_cmd_bin := exec.Command("curl", "-L", cyme_bin_url, "-o", "binary.bin")
	err = cyme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cyme_src_url := "https://github.com/tuna-f1sh/cyme/archive/refs/tags/v1.8.2.src.tar.gz"
	cyme_cmd_src := exec.Command("curl", "-L", cyme_src_url, "-o", "source.tar.gz")
	err = cyme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cyme_cmd_direct := exec.Command("./binary")
	err = cyme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
