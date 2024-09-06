package main

// Mqttui - Subscribe to a MQTT Topic or publish something quickly from the terminal
// Homepage: https://github.com/EdJoPaTo/mqttui

import (
	"fmt"
	
	"os/exec"
)

func installMqttui() {
	// Método 1: Descargar y extraer .tar.gz
	mqttui_tar_url := "https://github.com/EdJoPaTo/mqttui/archive/refs/tags/v0.21.1.tar.gz"
	mqttui_cmd_tar := exec.Command("curl", "-L", mqttui_tar_url, "-o", "package.tar.gz")
	err := mqttui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mqttui_zip_url := "https://github.com/EdJoPaTo/mqttui/archive/refs/tags/v0.21.1.zip"
	mqttui_cmd_zip := exec.Command("curl", "-L", mqttui_zip_url, "-o", "package.zip")
	err = mqttui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mqttui_bin_url := "https://github.com/EdJoPaTo/mqttui/archive/refs/tags/v0.21.1.bin"
	mqttui_cmd_bin := exec.Command("curl", "-L", mqttui_bin_url, "-o", "binary.bin")
	err = mqttui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mqttui_src_url := "https://github.com/EdJoPaTo/mqttui/archive/refs/tags/v0.21.1.src.tar.gz"
	mqttui_cmd_src := exec.Command("curl", "-L", mqttui_src_url, "-o", "source.tar.gz")
	err = mqttui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mqttui_cmd_direct := exec.Command("./binary")
	err = mqttui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
