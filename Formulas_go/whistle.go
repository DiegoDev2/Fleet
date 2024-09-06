package main

// Whistle - HTTP, HTTP2, HTTPS, Websocket debugging proxy
// Homepage: https://github.com/avwo/whistle

import (
	"fmt"
	
	"os/exec"
)

func installWhistle() {
	// Método 1: Descargar y extraer .tar.gz
	whistle_tar_url := "https://registry.npmjs.org/whistle/-/whistle-2.9.84.tgz"
	whistle_cmd_tar := exec.Command("curl", "-L", whistle_tar_url, "-o", "package.tar.gz")
	err := whistle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	whistle_zip_url := "https://registry.npmjs.org/whistle/-/whistle-2.9.84.tgz"
	whistle_cmd_zip := exec.Command("curl", "-L", whistle_zip_url, "-o", "package.zip")
	err = whistle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	whistle_bin_url := "https://registry.npmjs.org/whistle/-/whistle-2.9.84.tgz"
	whistle_cmd_bin := exec.Command("curl", "-L", whistle_bin_url, "-o", "binary.bin")
	err = whistle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	whistle_src_url := "https://registry.npmjs.org/whistle/-/whistle-2.9.84.tgz"
	whistle_cmd_src := exec.Command("curl", "-L", whistle_src_url, "-o", "source.tar.gz")
	err = whistle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	whistle_cmd_direct := exec.Command("./binary")
	err = whistle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
