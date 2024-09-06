package main

// Localtunnel - Exposes your localhost to the world for easy testing and sharing
// Homepage: https://theboroer.github.io/localtunnel-www/

import (
	"fmt"
	
	"os/exec"
)

func installLocaltunnel() {
	// Método 1: Descargar y extraer .tar.gz
	localtunnel_tar_url := "https://registry.npmjs.org/localtunnel/-/localtunnel-2.0.2.tgz"
	localtunnel_cmd_tar := exec.Command("curl", "-L", localtunnel_tar_url, "-o", "package.tar.gz")
	err := localtunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	localtunnel_zip_url := "https://registry.npmjs.org/localtunnel/-/localtunnel-2.0.2.tgz"
	localtunnel_cmd_zip := exec.Command("curl", "-L", localtunnel_zip_url, "-o", "package.zip")
	err = localtunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	localtunnel_bin_url := "https://registry.npmjs.org/localtunnel/-/localtunnel-2.0.2.tgz"
	localtunnel_cmd_bin := exec.Command("curl", "-L", localtunnel_bin_url, "-o", "binary.bin")
	err = localtunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	localtunnel_src_url := "https://registry.npmjs.org/localtunnel/-/localtunnel-2.0.2.tgz"
	localtunnel_cmd_src := exec.Command("curl", "-L", localtunnel_src_url, "-o", "source.tar.gz")
	err = localtunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	localtunnel_cmd_direct := exec.Command("./binary")
	err = localtunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
