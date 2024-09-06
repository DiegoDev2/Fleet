package main

// Vis - Vim-like text editor
// Homepage: https://github.com/martanne/vis

import (
	"fmt"
	
	"os/exec"
)

func installVis() {
	// Método 1: Descargar y extraer .tar.gz
	vis_tar_url := "https://github.com/martanne/vis/archive/refs/tags/v0.9.tar.gz"
	vis_cmd_tar := exec.Command("curl", "-L", vis_tar_url, "-o", "package.tar.gz")
	err := vis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vis_zip_url := "https://github.com/martanne/vis/archive/refs/tags/v0.9.zip"
	vis_cmd_zip := exec.Command("curl", "-L", vis_zip_url, "-o", "package.zip")
	err = vis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vis_bin_url := "https://github.com/martanne/vis/archive/refs/tags/v0.9.bin"
	vis_cmd_bin := exec.Command("curl", "-L", vis_bin_url, "-o", "binary.bin")
	err = vis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vis_src_url := "https://github.com/martanne/vis/archive/refs/tags/v0.9.src.tar.gz"
	vis_cmd_src := exec.Command("curl", "-L", vis_src_url, "-o", "source.tar.gz")
	err = vis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vis_cmd_direct := exec.Command("./binary")
	err = vis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libtermkey")
	exec.Command("latte", "install", "libtermkey").Run()
	fmt.Println("Instalando dependencia: lpeg")
	exec.Command("latte", "install", "lpeg").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
}
