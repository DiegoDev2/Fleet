package main

// Ddgr - DuckDuckGo from the terminal
// Homepage: https://github.com/jarun/ddgr

import (
	"fmt"
	
	"os/exec"
)

func installDdgr() {
	// Método 1: Descargar y extraer .tar.gz
	ddgr_tar_url := "https://github.com/jarun/ddgr/archive/refs/tags/v2.2.tar.gz"
	ddgr_cmd_tar := exec.Command("curl", "-L", ddgr_tar_url, "-o", "package.tar.gz")
	err := ddgr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddgr_zip_url := "https://github.com/jarun/ddgr/archive/refs/tags/v2.2.zip"
	ddgr_cmd_zip := exec.Command("curl", "-L", ddgr_zip_url, "-o", "package.zip")
	err = ddgr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddgr_bin_url := "https://github.com/jarun/ddgr/archive/refs/tags/v2.2.bin"
	ddgr_cmd_bin := exec.Command("curl", "-L", ddgr_bin_url, "-o", "binary.bin")
	err = ddgr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddgr_src_url := "https://github.com/jarun/ddgr/archive/refs/tags/v2.2.src.tar.gz"
	ddgr_cmd_src := exec.Command("curl", "-L", ddgr_src_url, "-o", "source.tar.gz")
	err = ddgr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddgr_cmd_direct := exec.Command("./binary")
	err = ddgr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
