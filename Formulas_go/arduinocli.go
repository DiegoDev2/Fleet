package main

// ArduinoCli - Arduino command-line interface
// Homepage: https://github.com/arduino/arduino-cli

import (
	"fmt"
	
	"os/exec"
)

func installArduinoCli() {
	// Método 1: Descargar y extraer .tar.gz
	arduinocli_tar_url := "https://github.com/arduino/arduino-cli/archive/refs/tags/v1.0.4.tar.gz"
	arduinocli_cmd_tar := exec.Command("curl", "-L", arduinocli_tar_url, "-o", "package.tar.gz")
	err := arduinocli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arduinocli_zip_url := "https://github.com/arduino/arduino-cli/archive/refs/tags/v1.0.4.zip"
	arduinocli_cmd_zip := exec.Command("curl", "-L", arduinocli_zip_url, "-o", "package.zip")
	err = arduinocli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arduinocli_bin_url := "https://github.com/arduino/arduino-cli/archive/refs/tags/v1.0.4.bin"
	arduinocli_cmd_bin := exec.Command("curl", "-L", arduinocli_bin_url, "-o", "binary.bin")
	err = arduinocli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arduinocli_src_url := "https://github.com/arduino/arduino-cli/archive/refs/tags/v1.0.4.src.tar.gz"
	arduinocli_cmd_src := exec.Command("curl", "-L", arduinocli_src_url, "-o", "source.tar.gz")
	err = arduinocli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arduinocli_cmd_direct := exec.Command("./binary")
	err = arduinocli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
