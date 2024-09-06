package main

// Cahute - Library and set of utilities to interact with Casio calculators
// Homepage: https://cahuteproject.org/

import (
	"fmt"
	
	"os/exec"
)

func installCahute() {
	// Método 1: Descargar y extraer .tar.gz
	cahute_tar_url := "https://ftp.cahuteproject.org/releases/cahute-0.5.tar.gz"
	cahute_cmd_tar := exec.Command("curl", "-L", cahute_tar_url, "-o", "package.tar.gz")
	err := cahute_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cahute_zip_url := "https://ftp.cahuteproject.org/releases/cahute-0.5.zip"
	cahute_cmd_zip := exec.Command("curl", "-L", cahute_zip_url, "-o", "package.zip")
	err = cahute_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cahute_bin_url := "https://ftp.cahuteproject.org/releases/cahute-0.5.bin"
	cahute_cmd_bin := exec.Command("curl", "-L", cahute_bin_url, "-o", "binary.bin")
	err = cahute_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cahute_src_url := "https://ftp.cahuteproject.org/releases/cahute-0.5.src.tar.gz"
	cahute_cmd_src := exec.Command("curl", "-L", cahute_src_url, "-o", "source.tar.gz")
	err = cahute_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cahute_cmd_direct := exec.Command("./binary")
	err = cahute_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
