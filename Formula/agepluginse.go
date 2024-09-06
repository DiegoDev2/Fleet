package main

// AgePluginSe - Age plugin for Apple Secure Enclave
// Homepage: https://github.com/remko/age-plugin-se

import (
	"fmt"
	
	"os/exec"
)

func installAgePluginSe() {
	// Método 1: Descargar y extraer .tar.gz
	agepluginse_tar_url := "https://github.com/remko/age-plugin-se/archive/refs/tags/v0.1.4.tar.gz"
	agepluginse_cmd_tar := exec.Command("curl", "-L", agepluginse_tar_url, "-o", "package.tar.gz")
	err := agepluginse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	agepluginse_zip_url := "https://github.com/remko/age-plugin-se/archive/refs/tags/v0.1.4.zip"
	agepluginse_cmd_zip := exec.Command("curl", "-L", agepluginse_zip_url, "-o", "package.zip")
	err = agepluginse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	agepluginse_bin_url := "https://github.com/remko/age-plugin-se/archive/refs/tags/v0.1.4.bin"
	agepluginse_cmd_bin := exec.Command("curl", "-L", agepluginse_bin_url, "-o", "binary.bin")
	err = agepluginse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	agepluginse_src_url := "https://github.com/remko/age-plugin-se/archive/refs/tags/v0.1.4.src.tar.gz"
	agepluginse_cmd_src := exec.Command("curl", "-L", agepluginse_src_url, "-o", "source.tar.gz")
	err = agepluginse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	agepluginse_cmd_direct := exec.Command("./binary")
	err = agepluginse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
	fmt.Println("Instalando dependencia: age")
	exec.Command("latte", "install", "age").Run()
}
