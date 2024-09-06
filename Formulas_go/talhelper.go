package main

// Talhelper - Configuration helper for talos clusters
// Homepage: https://budimanjojo.github.io/talhelper/latest/

import (
	"fmt"
	
	"os/exec"
)

func installTalhelper() {
	// Método 1: Descargar y extraer .tar.gz
	talhelper_tar_url := "https://github.com/budimanjojo/talhelper/archive/refs/tags/v3.0.5.tar.gz"
	talhelper_cmd_tar := exec.Command("curl", "-L", talhelper_tar_url, "-o", "package.tar.gz")
	err := talhelper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	talhelper_zip_url := "https://github.com/budimanjojo/talhelper/archive/refs/tags/v3.0.5.zip"
	talhelper_cmd_zip := exec.Command("curl", "-L", talhelper_zip_url, "-o", "package.zip")
	err = talhelper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	talhelper_bin_url := "https://github.com/budimanjojo/talhelper/archive/refs/tags/v3.0.5.bin"
	talhelper_cmd_bin := exec.Command("curl", "-L", talhelper_bin_url, "-o", "binary.bin")
	err = talhelper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	talhelper_src_url := "https://github.com/budimanjojo/talhelper/archive/refs/tags/v3.0.5.src.tar.gz"
	talhelper_cmd_src := exec.Command("curl", "-L", talhelper_src_url, "-o", "source.tar.gz")
	err = talhelper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	talhelper_cmd_direct := exec.Command("./binary")
	err = talhelper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
