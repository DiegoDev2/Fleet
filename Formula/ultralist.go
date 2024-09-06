package main

// Ultralist - Simple GTD-style task management for the command-line
// Homepage: https://ultralist.io

import (
	"fmt"
	
	"os/exec"
)

func installUltralist() {
	// Método 1: Descargar y extraer .tar.gz
	ultralist_tar_url := "https://github.com/ultralist/ultralist/archive/refs/tags/1.7.0.tar.gz"
	ultralist_cmd_tar := exec.Command("curl", "-L", ultralist_tar_url, "-o", "package.tar.gz")
	err := ultralist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ultralist_zip_url := "https://github.com/ultralist/ultralist/archive/refs/tags/1.7.0.zip"
	ultralist_cmd_zip := exec.Command("curl", "-L", ultralist_zip_url, "-o", "package.zip")
	err = ultralist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ultralist_bin_url := "https://github.com/ultralist/ultralist/archive/refs/tags/1.7.0.bin"
	ultralist_cmd_bin := exec.Command("curl", "-L", ultralist_bin_url, "-o", "binary.bin")
	err = ultralist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ultralist_src_url := "https://github.com/ultralist/ultralist/archive/refs/tags/1.7.0.src.tar.gz"
	ultralist_cmd_src := exec.Command("curl", "-L", ultralist_src_url, "-o", "source.tar.gz")
	err = ultralist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ultralist_cmd_direct := exec.Command("./binary")
	err = ultralist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
