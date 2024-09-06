package main

// Nifi - Easy to use, powerful, and reliable system to process and distribute data
// Homepage: https://nifi.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installNifi() {
	// Método 1: Descargar y extraer .tar.gz
	nifi_tar_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-1.27.0-bin.zip"
	nifi_cmd_tar := exec.Command("curl", "-L", nifi_tar_url, "-o", "package.tar.gz")
	err := nifi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nifi_zip_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-1.27.0-bin.zip"
	nifi_cmd_zip := exec.Command("curl", "-L", nifi_zip_url, "-o", "package.zip")
	err = nifi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nifi_bin_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-1.27.0-bin.zip"
	nifi_cmd_bin := exec.Command("curl", "-L", nifi_bin_url, "-o", "binary.bin")
	err = nifi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nifi_src_url := "https://www.apache.org/dyn/closer.lua?path=/nifi/1.27.0/nifi-1.27.0-bin.zip"
	nifi_cmd_src := exec.Command("curl", "-L", nifi_src_url, "-o", "source.tar.gz")
	err = nifi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nifi_cmd_direct := exec.Command("./binary")
	err = nifi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
