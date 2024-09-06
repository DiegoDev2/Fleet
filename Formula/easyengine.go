package main

// Easyengine - Command-line control panel to manage WordPress sites
// Homepage: https://easyengine.io/

import (
	"fmt"
	
	"os/exec"
)

func installEasyengine() {
	// Método 1: Descargar y extraer .tar.gz
	easyengine_tar_url := "https://github.com/EasyEngine/easyengine/releases/download/v4.7.3/easyengine.phar"
	easyengine_cmd_tar := exec.Command("curl", "-L", easyengine_tar_url, "-o", "package.tar.gz")
	err := easyengine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	easyengine_zip_url := "https://github.com/EasyEngine/easyengine/releases/download/v4.7.3/easyengine.phar"
	easyengine_cmd_zip := exec.Command("curl", "-L", easyengine_zip_url, "-o", "package.zip")
	err = easyengine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	easyengine_bin_url := "https://github.com/EasyEngine/easyengine/releases/download/v4.7.3/easyengine.phar"
	easyengine_cmd_bin := exec.Command("curl", "-L", easyengine_bin_url, "-o", "binary.bin")
	err = easyengine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	easyengine_src_url := "https://github.com/EasyEngine/easyengine/releases/download/v4.7.3/easyengine.phar"
	easyengine_cmd_src := exec.Command("curl", "-L", easyengine_src_url, "-o", "source.tar.gz")
	err = easyengine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	easyengine_cmd_direct := exec.Command("./binary")
	err = easyengine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dnsmasq")
	exec.Command("latte", "install", "dnsmasq").Run()
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
