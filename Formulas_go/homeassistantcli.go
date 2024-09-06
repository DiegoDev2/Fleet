package main

// HomeassistantCli - Command-line utility for Home Assistant
// Homepage: https://github.com/home-assistant-ecosystem/home-assistant-cli

import (
	"fmt"
	
	"os/exec"
)

func installHomeassistantCli() {
	// Método 1: Descargar y extraer .tar.gz
	homeassistantcli_tar_url := "https://files.pythonhosted.org/packages/b2/98/fd5e7beb7cc135f80d78b32c85ac15f3ba9219063b794b1d184fb07fd84b/homeassistant-cli-0.9.6.tar.gz"
	homeassistantcli_cmd_tar := exec.Command("curl", "-L", homeassistantcli_tar_url, "-o", "package.tar.gz")
	err := homeassistantcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	homeassistantcli_zip_url := "https://files.pythonhosted.org/packages/b2/98/fd5e7beb7cc135f80d78b32c85ac15f3ba9219063b794b1d184fb07fd84b/homeassistant-cli-0.9.6.zip"
	homeassistantcli_cmd_zip := exec.Command("curl", "-L", homeassistantcli_zip_url, "-o", "package.zip")
	err = homeassistantcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	homeassistantcli_bin_url := "https://files.pythonhosted.org/packages/b2/98/fd5e7beb7cc135f80d78b32c85ac15f3ba9219063b794b1d184fb07fd84b/homeassistant-cli-0.9.6.bin"
	homeassistantcli_cmd_bin := exec.Command("curl", "-L", homeassistantcli_bin_url, "-o", "binary.bin")
	err = homeassistantcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	homeassistantcli_src_url := "https://files.pythonhosted.org/packages/b2/98/fd5e7beb7cc135f80d78b32c85ac15f3ba9219063b794b1d184fb07fd84b/homeassistant-cli-0.9.6.src.tar.gz"
	homeassistantcli_cmd_src := exec.Command("curl", "-L", homeassistantcli_src_url, "-o", "source.tar.gz")
	err = homeassistantcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	homeassistantcli_cmd_direct := exec.Command("./binary")
	err = homeassistantcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
