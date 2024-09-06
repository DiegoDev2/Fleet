package main

// Ubertooth - Host tools for Project Ubertooth
// Homepage: https://github.com/greatscottgadgets/ubertooth/wiki

import (
	"fmt"
	
	"os/exec"
)

func installUbertooth() {
	// Método 1: Descargar y extraer .tar.gz
	ubertooth_tar_url := "https://github.com/greatscottgadgets/ubertooth/releases/download/2020-12-R1/ubertooth-2020-12-R1.tar.xz"
	ubertooth_cmd_tar := exec.Command("curl", "-L", ubertooth_tar_url, "-o", "package.tar.gz")
	err := ubertooth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ubertooth_zip_url := "https://github.com/greatscottgadgets/ubertooth/releases/download/2020-12-R1/ubertooth-2020-12-R1.tar.xz"
	ubertooth_cmd_zip := exec.Command("curl", "-L", ubertooth_zip_url, "-o", "package.zip")
	err = ubertooth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ubertooth_bin_url := "https://github.com/greatscottgadgets/ubertooth/releases/download/2020-12-R1/ubertooth-2020-12-R1.tar.xz"
	ubertooth_cmd_bin := exec.Command("curl", "-L", ubertooth_bin_url, "-o", "binary.bin")
	err = ubertooth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ubertooth_src_url := "https://github.com/greatscottgadgets/ubertooth/releases/download/2020-12-R1/ubertooth-2020-12-R1.tar.xz"
	ubertooth_cmd_src := exec.Command("curl", "-L", ubertooth_src_url, "-o", "source.tar.gz")
	err = ubertooth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ubertooth_cmd_direct := exec.Command("./binary")
	err = ubertooth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libbtbb")
	exec.Command("latte", "install", "libbtbb").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
