package main

// Appium - Automation for Apps
// Homepage: https://appium.io/

import (
	"fmt"
	
	"os/exec"
)

func installAppium() {
	// Método 1: Descargar y extraer .tar.gz
	appium_tar_url := "https://registry.npmjs.org/appium/-/appium-2.11.3.tgz"
	appium_cmd_tar := exec.Command("curl", "-L", appium_tar_url, "-o", "package.tar.gz")
	err := appium_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	appium_zip_url := "https://registry.npmjs.org/appium/-/appium-2.11.3.tgz"
	appium_cmd_zip := exec.Command("curl", "-L", appium_zip_url, "-o", "package.zip")
	err = appium_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	appium_bin_url := "https://registry.npmjs.org/appium/-/appium-2.11.3.tgz"
	appium_cmd_bin := exec.Command("curl", "-L", appium_bin_url, "-o", "binary.bin")
	err = appium_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	appium_src_url := "https://registry.npmjs.org/appium/-/appium-2.11.3.tgz"
	appium_cmd_src := exec.Command("curl", "-L", appium_src_url, "-o", "source.tar.gz")
	err = appium_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	appium_cmd_direct := exec.Command("./binary")
	err = appium_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: vips")
exec.Command("latte", "install", "vips")
}
