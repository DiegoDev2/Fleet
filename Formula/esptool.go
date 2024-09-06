package main

// Esptool - ESP8266 and ESP32 serial bootloader utility
// Homepage: https://docs.espressif.com/projects/esptool/en/latest/esp32/

import (
	"fmt"
	
	"os/exec"
)

func installEsptool() {
	// Método 1: Descargar y extraer .tar.gz
	esptool_tar_url := "https://files.pythonhosted.org/packages/1b/8b/f0d1e75879dee053874a4f955ed1e9ad97275485f51cb4bc2cb4e9b24479/esptool-4.7.0.tar.gz"
	esptool_cmd_tar := exec.Command("curl", "-L", esptool_tar_url, "-o", "package.tar.gz")
	err := esptool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	esptool_zip_url := "https://files.pythonhosted.org/packages/1b/8b/f0d1e75879dee053874a4f955ed1e9ad97275485f51cb4bc2cb4e9b24479/esptool-4.7.0.zip"
	esptool_cmd_zip := exec.Command("curl", "-L", esptool_zip_url, "-o", "package.zip")
	err = esptool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	esptool_bin_url := "https://files.pythonhosted.org/packages/1b/8b/f0d1e75879dee053874a4f955ed1e9ad97275485f51cb4bc2cb4e9b24479/esptool-4.7.0.bin"
	esptool_cmd_bin := exec.Command("curl", "-L", esptool_bin_url, "-o", "binary.bin")
	err = esptool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	esptool_src_url := "https://files.pythonhosted.org/packages/1b/8b/f0d1e75879dee053874a4f955ed1e9ad97275485f51cb4bc2cb4e9b24479/esptool-4.7.0.src.tar.gz"
	esptool_cmd_src := exec.Command("curl", "-L", esptool_src_url, "-o", "source.tar.gz")
	err = esptool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	esptool_cmd_direct := exec.Command("./binary")
	err = esptool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
