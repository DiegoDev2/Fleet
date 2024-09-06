package main

// TeensyLoaderCli - Command-line integration for Teensy USB development boards
// Homepage: https://www.pjrc.com/teensy/loader_cli.html

import (
	"fmt"
	
	"os/exec"
)

func installTeensyLoaderCli() {
	// Método 1: Descargar y extraer .tar.gz
	teensyloadercli_tar_url := "https://github.com/PaulStoffregen/teensy_loader_cli/archive/refs/tags/2.3.tar.gz"
	teensyloadercli_cmd_tar := exec.Command("curl", "-L", teensyloadercli_tar_url, "-o", "package.tar.gz")
	err := teensyloadercli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	teensyloadercli_zip_url := "https://github.com/PaulStoffregen/teensy_loader_cli/archive/refs/tags/2.3.zip"
	teensyloadercli_cmd_zip := exec.Command("curl", "-L", teensyloadercli_zip_url, "-o", "package.zip")
	err = teensyloadercli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	teensyloadercli_bin_url := "https://github.com/PaulStoffregen/teensy_loader_cli/archive/refs/tags/2.3.bin"
	teensyloadercli_cmd_bin := exec.Command("curl", "-L", teensyloadercli_bin_url, "-o", "binary.bin")
	err = teensyloadercli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	teensyloadercli_src_url := "https://github.com/PaulStoffregen/teensy_loader_cli/archive/refs/tags/2.3.src.tar.gz"
	teensyloadercli_cmd_src := exec.Command("curl", "-L", teensyloadercli_src_url, "-o", "source.tar.gz")
	err = teensyloadercli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	teensyloadercli_cmd_direct := exec.Command("./binary")
	err = teensyloadercli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
