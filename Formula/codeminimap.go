package main

// CodeMinimap - High performance code minimap generator
// Homepage: https://github.com/wfxr/code-minimap

import (
	"fmt"
	
	"os/exec"
)

func installCodeMinimap() {
	// Método 1: Descargar y extraer .tar.gz
	codeminimap_tar_url := "https://github.com/wfxr/code-minimap/archive/refs/tags/v0.6.7.tar.gz"
	codeminimap_cmd_tar := exec.Command("curl", "-L", codeminimap_tar_url, "-o", "package.tar.gz")
	err := codeminimap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codeminimap_zip_url := "https://github.com/wfxr/code-minimap/archive/refs/tags/v0.6.7.zip"
	codeminimap_cmd_zip := exec.Command("curl", "-L", codeminimap_zip_url, "-o", "package.zip")
	err = codeminimap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codeminimap_bin_url := "https://github.com/wfxr/code-minimap/archive/refs/tags/v0.6.7.bin"
	codeminimap_cmd_bin := exec.Command("curl", "-L", codeminimap_bin_url, "-o", "binary.bin")
	err = codeminimap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codeminimap_src_url := "https://github.com/wfxr/code-minimap/archive/refs/tags/v0.6.7.src.tar.gz"
	codeminimap_cmd_src := exec.Command("curl", "-L", codeminimap_src_url, "-o", "source.tar.gz")
	err = codeminimap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codeminimap_cmd_direct := exec.Command("./binary")
	err = codeminimap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
