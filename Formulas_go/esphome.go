package main

// Esphome - Make creating custom firmwares for ESP32/ESP8266 super easy
// Homepage: https://github.com/esphome/esphome

import (
	"fmt"
	
	"os/exec"
)

func installEsphome() {
	// Método 1: Descargar y extraer .tar.gz
	esphome_tar_url := "https://files.pythonhosted.org/packages/54/2a/b57e21c5529c65a15f893b57a339fb75f6bec8c7e7c815b2b5268180e310/esphome-2024.8.3.tar.gz"
	esphome_cmd_tar := exec.Command("curl", "-L", esphome_tar_url, "-o", "package.tar.gz")
	err := esphome_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	esphome_zip_url := "https://files.pythonhosted.org/packages/54/2a/b57e21c5529c65a15f893b57a339fb75f6bec8c7e7c815b2b5268180e310/esphome-2024.8.3.zip"
	esphome_cmd_zip := exec.Command("curl", "-L", esphome_zip_url, "-o", "package.zip")
	err = esphome_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	esphome_bin_url := "https://files.pythonhosted.org/packages/54/2a/b57e21c5529c65a15f893b57a339fb75f6bec8c7e7c815b2b5268180e310/esphome-2024.8.3.bin"
	esphome_cmd_bin := exec.Command("curl", "-L", esphome_bin_url, "-o", "binary.bin")
	err = esphome_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	esphome_src_url := "https://files.pythonhosted.org/packages/54/2a/b57e21c5529c65a15f893b57a339fb75f6bec8c7e7c815b2b5268180e310/esphome-2024.8.3.src.tar.gz"
	esphome_cmd_src := exec.Command("curl", "-L", esphome_src_url, "-o", "source.tar.gz")
	err = esphome_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	esphome_cmd_direct := exec.Command("./binary")
	err = esphome_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
