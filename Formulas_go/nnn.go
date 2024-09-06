package main

// Nnn - Tiny, lightning fast, feature-packed file manager
// Homepage: https://github.com/jarun/nnn

import (
	"fmt"
	
	"os/exec"
)

func installNnn() {
	// Método 1: Descargar y extraer .tar.gz
	nnn_tar_url := "https://github.com/jarun/nnn/archive/refs/tags/v5.0.tar.gz"
	nnn_cmd_tar := exec.Command("curl", "-L", nnn_tar_url, "-o", "package.tar.gz")
	err := nnn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nnn_zip_url := "https://github.com/jarun/nnn/archive/refs/tags/v5.0.zip"
	nnn_cmd_zip := exec.Command("curl", "-L", nnn_zip_url, "-o", "package.zip")
	err = nnn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nnn_bin_url := "https://github.com/jarun/nnn/archive/refs/tags/v5.0.bin"
	nnn_cmd_bin := exec.Command("curl", "-L", nnn_bin_url, "-o", "binary.bin")
	err = nnn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nnn_src_url := "https://github.com/jarun/nnn/archive/refs/tags/v5.0.src.tar.gz"
	nnn_cmd_src := exec.Command("curl", "-L", nnn_src_url, "-o", "source.tar.gz")
	err = nnn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nnn_cmd_direct := exec.Command("./binary")
	err = nnn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
