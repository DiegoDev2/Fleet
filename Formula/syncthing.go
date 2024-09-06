package main

// Syncthing - Open source continuous file synchronization application
// Homepage: https://syncthing.net/

import (
	"fmt"
	
	"os/exec"
)

func installSyncthing() {
	// Método 1: Descargar y extraer .tar.gz
	syncthing_tar_url := "https://github.com/syncthing/syncthing/archive/refs/tags/v1.27.11.tar.gz"
	syncthing_cmd_tar := exec.Command("curl", "-L", syncthing_tar_url, "-o", "package.tar.gz")
	err := syncthing_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	syncthing_zip_url := "https://github.com/syncthing/syncthing/archive/refs/tags/v1.27.11.zip"
	syncthing_cmd_zip := exec.Command("curl", "-L", syncthing_zip_url, "-o", "package.zip")
	err = syncthing_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	syncthing_bin_url := "https://github.com/syncthing/syncthing/archive/refs/tags/v1.27.11.bin"
	syncthing_cmd_bin := exec.Command("curl", "-L", syncthing_bin_url, "-o", "binary.bin")
	err = syncthing_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	syncthing_src_url := "https://github.com/syncthing/syncthing/archive/refs/tags/v1.27.11.src.tar.gz"
	syncthing_cmd_src := exec.Command("curl", "-L", syncthing_src_url, "-o", "source.tar.gz")
	err = syncthing_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	syncthing_cmd_direct := exec.Command("./binary")
	err = syncthing_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
