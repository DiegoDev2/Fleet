package main

// Websocketd - WebSockets the Unix way
// Homepage: http://websocketd.com

import (
	"fmt"
	
	"os/exec"
)

func installWebsocketd() {
	// Método 1: Descargar y extraer .tar.gz
	websocketd_tar_url := "https://github.com/joewalnes/websocketd/archive/refs/tags/v0.4.1.tar.gz"
	websocketd_cmd_tar := exec.Command("curl", "-L", websocketd_tar_url, "-o", "package.tar.gz")
	err := websocketd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	websocketd_zip_url := "https://github.com/joewalnes/websocketd/archive/refs/tags/v0.4.1.zip"
	websocketd_cmd_zip := exec.Command("curl", "-L", websocketd_zip_url, "-o", "package.zip")
	err = websocketd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	websocketd_bin_url := "https://github.com/joewalnes/websocketd/archive/refs/tags/v0.4.1.bin"
	websocketd_cmd_bin := exec.Command("curl", "-L", websocketd_bin_url, "-o", "binary.bin")
	err = websocketd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	websocketd_src_url := "https://github.com/joewalnes/websocketd/archive/refs/tags/v0.4.1.src.tar.gz"
	websocketd_cmd_src := exec.Command("curl", "-L", websocketd_src_url, "-o", "source.tar.gz")
	err = websocketd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	websocketd_cmd_direct := exec.Command("./binary")
	err = websocketd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
