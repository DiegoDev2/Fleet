package main

// Libbtbb - Bluetooth baseband decoding library
// Homepage: https://github.com/greatscottgadgets/libbtbb

import (
	"fmt"
	
	"os/exec"
)

func installLibbtbb() {
	// Método 1: Descargar y extraer .tar.gz
	libbtbb_tar_url := "https://github.com/greatscottgadgets/libbtbb/archive/refs/tags/2020-12-R1.tar.gz"
	libbtbb_cmd_tar := exec.Command("curl", "-L", libbtbb_tar_url, "-o", "package.tar.gz")
	err := libbtbb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbtbb_zip_url := "https://github.com/greatscottgadgets/libbtbb/archive/refs/tags/2020-12-R1.zip"
	libbtbb_cmd_zip := exec.Command("curl", "-L", libbtbb_zip_url, "-o", "package.zip")
	err = libbtbb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbtbb_bin_url := "https://github.com/greatscottgadgets/libbtbb/archive/refs/tags/2020-12-R1.bin"
	libbtbb_cmd_bin := exec.Command("curl", "-L", libbtbb_bin_url, "-o", "binary.bin")
	err = libbtbb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbtbb_src_url := "https://github.com/greatscottgadgets/libbtbb/archive/refs/tags/2020-12-R1.src.tar.gz"
	libbtbb_cmd_src := exec.Command("curl", "-L", libbtbb_src_url, "-o", "source.tar.gz")
	err = libbtbb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbtbb_cmd_direct := exec.Command("./binary")
	err = libbtbb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
