package main

// TrzszGo - Simple file transfer tools, similar to lrzsz (rz/sz), and compatible with tmux
// Homepage: https://trzsz.github.io

import (
	"fmt"
	
	"os/exec"
)

func installTrzszGo() {
	// Método 1: Descargar y extraer .tar.gz
	trzszgo_tar_url := "https://github.com/trzsz/trzsz-go/archive/refs/tags/v1.1.8.tar.gz"
	trzszgo_cmd_tar := exec.Command("curl", "-L", trzszgo_tar_url, "-o", "package.tar.gz")
	err := trzszgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trzszgo_zip_url := "https://github.com/trzsz/trzsz-go/archive/refs/tags/v1.1.8.zip"
	trzszgo_cmd_zip := exec.Command("curl", "-L", trzszgo_zip_url, "-o", "package.zip")
	err = trzszgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trzszgo_bin_url := "https://github.com/trzsz/trzsz-go/archive/refs/tags/v1.1.8.bin"
	trzszgo_cmd_bin := exec.Command("curl", "-L", trzszgo_bin_url, "-o", "binary.bin")
	err = trzszgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trzszgo_src_url := "https://github.com/trzsz/trzsz-go/archive/refs/tags/v1.1.8.src.tar.gz"
	trzszgo_cmd_src := exec.Command("curl", "-L", trzszgo_src_url, "-o", "source.tar.gz")
	err = trzszgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trzszgo_cmd_direct := exec.Command("./binary")
	err = trzszgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
