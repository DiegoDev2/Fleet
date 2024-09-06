package main

// LibpahoMqtt - Eclipse Paho C client library for MQTT
// Homepage: https://eclipse.github.io/paho.mqtt.c/

import (
	"fmt"
	
	"os/exec"
)

func installLibpahoMqtt() {
	// Método 1: Descargar y extraer .tar.gz
	libpahomqtt_tar_url := "https://github.com/eclipse/paho.mqtt.c/archive/refs/tags/v1.3.13.tar.gz"
	libpahomqtt_cmd_tar := exec.Command("curl", "-L", libpahomqtt_tar_url, "-o", "package.tar.gz")
	err := libpahomqtt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpahomqtt_zip_url := "https://github.com/eclipse/paho.mqtt.c/archive/refs/tags/v1.3.13.zip"
	libpahomqtt_cmd_zip := exec.Command("curl", "-L", libpahomqtt_zip_url, "-o", "package.zip")
	err = libpahomqtt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpahomqtt_bin_url := "https://github.com/eclipse/paho.mqtt.c/archive/refs/tags/v1.3.13.bin"
	libpahomqtt_cmd_bin := exec.Command("curl", "-L", libpahomqtt_bin_url, "-o", "binary.bin")
	err = libpahomqtt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpahomqtt_src_url := "https://github.com/eclipse/paho.mqtt.c/archive/refs/tags/v1.3.13.src.tar.gz"
	libpahomqtt_cmd_src := exec.Command("curl", "-L", libpahomqtt_src_url, "-o", "source.tar.gz")
	err = libpahomqtt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpahomqtt_cmd_direct := exec.Command("./binary")
	err = libpahomqtt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
