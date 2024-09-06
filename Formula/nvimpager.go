package main

// Nvimpager - Use NeoVim as a pager to view manpages, diffs, etc.
// Homepage: https://github.com/lucc/nvimpager

import (
	"fmt"
	
	"os/exec"
)

func installNvimpager() {
	// Método 1: Descargar y extraer .tar.gz
	nvimpager_tar_url := "https://github.com/lucc/nvimpager/archive/refs/tags/v0.13.0.tar.gz"
	nvimpager_cmd_tar := exec.Command("curl", "-L", nvimpager_tar_url, "-o", "package.tar.gz")
	err := nvimpager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nvimpager_zip_url := "https://github.com/lucc/nvimpager/archive/refs/tags/v0.13.0.zip"
	nvimpager_cmd_zip := exec.Command("curl", "-L", nvimpager_zip_url, "-o", "package.zip")
	err = nvimpager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nvimpager_bin_url := "https://github.com/lucc/nvimpager/archive/refs/tags/v0.13.0.bin"
	nvimpager_cmd_bin := exec.Command("curl", "-L", nvimpager_bin_url, "-o", "binary.bin")
	err = nvimpager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nvimpager_src_url := "https://github.com/lucc/nvimpager/archive/refs/tags/v0.13.0.src.tar.gz"
	nvimpager_cmd_src := exec.Command("curl", "-L", nvimpager_src_url, "-o", "source.tar.gz")
	err = nvimpager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nvimpager_cmd_direct := exec.Command("./binary")
	err = nvimpager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
	fmt.Println("Instalando dependencia: neovim")
	exec.Command("latte", "install", "neovim").Run()
}
