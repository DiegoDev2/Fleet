package main

// LadspaSdk - Linux Audio Developer's Simple Plugin
// Homepage: https://www.ladspa.org

import (
	"fmt"
	
	"os/exec"
)

func installLadspaSdk() {
	// Método 1: Descargar y extraer .tar.gz
	ladspasdk_tar_url := "https://www.ladspa.org/download/ladspa_sdk_1.17.tgz"
	ladspasdk_cmd_tar := exec.Command("curl", "-L", ladspasdk_tar_url, "-o", "package.tar.gz")
	err := ladspasdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ladspasdk_zip_url := "https://www.ladspa.org/download/ladspa_sdk_1.17.tgz"
	ladspasdk_cmd_zip := exec.Command("curl", "-L", ladspasdk_zip_url, "-o", "package.zip")
	err = ladspasdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ladspasdk_bin_url := "https://www.ladspa.org/download/ladspa_sdk_1.17.tgz"
	ladspasdk_cmd_bin := exec.Command("curl", "-L", ladspasdk_bin_url, "-o", "binary.bin")
	err = ladspasdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ladspasdk_src_url := "https://www.ladspa.org/download/ladspa_sdk_1.17.tgz"
	ladspasdk_cmd_src := exec.Command("curl", "-L", ladspasdk_src_url, "-o", "source.tar.gz")
	err = ladspasdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ladspasdk_cmd_direct := exec.Command("./binary")
	err = ladspasdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
}
