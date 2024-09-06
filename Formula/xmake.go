package main

// Xmake - Cross-platform build utility based on Lua
// Homepage: https://xmake.io/

import (
	"fmt"
	
	"os/exec"
)

func installXmake() {
	// Método 1: Descargar y extraer .tar.gz
	xmake_tar_url := "https://github.com/xmake-io/xmake/releases/download/v2.9.4/xmake-v2.9.4.tar.gz"
	xmake_cmd_tar := exec.Command("curl", "-L", xmake_tar_url, "-o", "package.tar.gz")
	err := xmake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmake_zip_url := "https://github.com/xmake-io/xmake/releases/download/v2.9.4/xmake-v2.9.4.zip"
	xmake_cmd_zip := exec.Command("curl", "-L", xmake_zip_url, "-o", "package.zip")
	err = xmake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmake_bin_url := "https://github.com/xmake-io/xmake/releases/download/v2.9.4/xmake-v2.9.4.bin"
	xmake_cmd_bin := exec.Command("curl", "-L", xmake_bin_url, "-o", "binary.bin")
	err = xmake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmake_src_url := "https://github.com/xmake-io/xmake/releases/download/v2.9.4/xmake-v2.9.4.src.tar.gz"
	xmake_cmd_src := exec.Command("curl", "-L", xmake_src_url, "-o", "source.tar.gz")
	err = xmake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmake_cmd_direct := exec.Command("./binary")
	err = xmake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
