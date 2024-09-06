package main

// Avrdude - Atmel AVR MCU programmer
// Homepage: https://www.nongnu.org/avrdude/

import (
	"fmt"
	
	"os/exec"
)

func installAvrdude() {
	// Método 1: Descargar y extraer .tar.gz
	avrdude_tar_url := "https://github.com/avrdudes/avrdude/archive/refs/tags/v8.0.tar.gz"
	avrdude_cmd_tar := exec.Command("curl", "-L", avrdude_tar_url, "-o", "package.tar.gz")
	err := avrdude_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avrdude_zip_url := "https://github.com/avrdudes/avrdude/archive/refs/tags/v8.0.zip"
	avrdude_cmd_zip := exec.Command("curl", "-L", avrdude_zip_url, "-o", "package.zip")
	err = avrdude_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avrdude_bin_url := "https://github.com/avrdudes/avrdude/archive/refs/tags/v8.0.bin"
	avrdude_cmd_bin := exec.Command("curl", "-L", avrdude_bin_url, "-o", "binary.bin")
	err = avrdude_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avrdude_src_url := "https://github.com/avrdudes/avrdude/archive/refs/tags/v8.0.src.tar.gz"
	avrdude_cmd_src := exec.Command("curl", "-L", avrdude_src_url, "-o", "source.tar.gz")
	err = avrdude_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avrdude_cmd_direct := exec.Command("./binary")
	err = avrdude_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: hidapi")
exec.Command("latte", "install", "hidapi")
	fmt.Println("Instalando dependencia: libftdi")
exec.Command("latte", "install", "libftdi")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
	fmt.Println("Instalando dependencia: libelf")
exec.Command("latte", "install", "libelf")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
