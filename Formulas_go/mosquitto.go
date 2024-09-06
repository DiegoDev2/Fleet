package main

// Mosquitto - Message broker implementing the MQTT protocol
// Homepage: https://mosquitto.org/

import (
	"fmt"
	
	"os/exec"
)

func installMosquitto() {
	// Método 1: Descargar y extraer .tar.gz
	mosquitto_tar_url := "https://mosquitto.org/files/source/mosquitto-2.0.18.tar.gz"
	mosquitto_cmd_tar := exec.Command("curl", "-L", mosquitto_tar_url, "-o", "package.tar.gz")
	err := mosquitto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mosquitto_zip_url := "https://mosquitto.org/files/source/mosquitto-2.0.18.zip"
	mosquitto_cmd_zip := exec.Command("curl", "-L", mosquitto_zip_url, "-o", "package.zip")
	err = mosquitto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mosquitto_bin_url := "https://mosquitto.org/files/source/mosquitto-2.0.18.bin"
	mosquitto_cmd_bin := exec.Command("curl", "-L", mosquitto_bin_url, "-o", "binary.bin")
	err = mosquitto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mosquitto_src_url := "https://mosquitto.org/files/source/mosquitto-2.0.18.src.tar.gz"
	mosquitto_cmd_src := exec.Command("curl", "-L", mosquitto_src_url, "-o", "source.tar.gz")
	err = mosquitto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mosquitto_cmd_direct := exec.Command("./binary")
	err = mosquitto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cjson")
exec.Command("latte", "install", "cjson")
	fmt.Println("Instalando dependencia: libwebsockets")
exec.Command("latte", "install", "libwebsockets")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
