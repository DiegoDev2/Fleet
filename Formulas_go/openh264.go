package main

// Openh264 - H.264 codec from Cisco
// Homepage: https://www.openh264.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenh264() {
	// Método 1: Descargar y extraer .tar.gz
	openh264_tar_url := "https://github.com/cisco/openh264/archive/refs/tags/v2.4.1.tar.gz"
	openh264_cmd_tar := exec.Command("curl", "-L", openh264_tar_url, "-o", "package.tar.gz")
	err := openh264_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openh264_zip_url := "https://github.com/cisco/openh264/archive/refs/tags/v2.4.1.zip"
	openh264_cmd_zip := exec.Command("curl", "-L", openh264_zip_url, "-o", "package.zip")
	err = openh264_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openh264_bin_url := "https://github.com/cisco/openh264/archive/refs/tags/v2.4.1.bin"
	openh264_cmd_bin := exec.Command("curl", "-L", openh264_bin_url, "-o", "binary.bin")
	err = openh264_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openh264_src_url := "https://github.com/cisco/openh264/archive/refs/tags/v2.4.1.src.tar.gz"
	openh264_cmd_src := exec.Command("curl", "-L", openh264_src_url, "-o", "source.tar.gz")
	err = openh264_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openh264_cmd_direct := exec.Command("./binary")
	err = openh264_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}
