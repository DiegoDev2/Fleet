package main

// Analog - Logfile analyzer
// Homepage: https://www.c-amie.co.uk/software/analog/

import (
	"fmt"
	
	"os/exec"
)

func installAnalog() {
	// Método 1: Descargar y extraer .tar.gz
	analog_tar_url := "https://github.com/c-amie/analog-ce/archive/refs/tags/6.0.17.tar.gz"
	analog_cmd_tar := exec.Command("curl", "-L", analog_tar_url, "-o", "package.tar.gz")
	err := analog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	analog_zip_url := "https://github.com/c-amie/analog-ce/archive/refs/tags/6.0.17.zip"
	analog_cmd_zip := exec.Command("curl", "-L", analog_zip_url, "-o", "package.zip")
	err = analog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	analog_bin_url := "https://github.com/c-amie/analog-ce/archive/refs/tags/6.0.17.bin"
	analog_cmd_bin := exec.Command("curl", "-L", analog_bin_url, "-o", "binary.bin")
	err = analog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	analog_src_url := "https://github.com/c-amie/analog-ce/archive/refs/tags/6.0.17.src.tar.gz"
	analog_cmd_src := exec.Command("curl", "-L", analog_src_url, "-o", "source.tar.gz")
	err = analog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	analog_cmd_direct := exec.Command("./binary")
	err = analog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
