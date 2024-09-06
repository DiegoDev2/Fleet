package main

// Clipper - Share macOS clipboard with tmux and other local and remote apps
// Homepage: https://github.com/wincent/clipper

import (
	"fmt"
	
	"os/exec"
)

func installClipper() {
	// Método 1: Descargar y extraer .tar.gz
	clipper_tar_url := "https://github.com/wincent/clipper/archive/refs/tags/2.1.0.tar.gz"
	clipper_cmd_tar := exec.Command("curl", "-L", clipper_tar_url, "-o", "package.tar.gz")
	err := clipper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clipper_zip_url := "https://github.com/wincent/clipper/archive/refs/tags/2.1.0.zip"
	clipper_cmd_zip := exec.Command("curl", "-L", clipper_zip_url, "-o", "package.zip")
	err = clipper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clipper_bin_url := "https://github.com/wincent/clipper/archive/refs/tags/2.1.0.bin"
	clipper_cmd_bin := exec.Command("curl", "-L", clipper_bin_url, "-o", "binary.bin")
	err = clipper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clipper_src_url := "https://github.com/wincent/clipper/archive/refs/tags/2.1.0.src.tar.gz"
	clipper_cmd_src := exec.Command("curl", "-L", clipper_src_url, "-o", "source.tar.gz")
	err = clipper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clipper_cmd_direct := exec.Command("./binary")
	err = clipper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
