package main

// Lazygit - Simple terminal UI for git commands
// Homepage: https://github.com/jesseduffield/lazygit/

import (
	"fmt"
	
	"os/exec"
)

func installLazygit() {
	// Método 1: Descargar y extraer .tar.gz
	lazygit_tar_url := "https://github.com/jesseduffield/lazygit/archive/refs/tags/v0.43.1.tar.gz"
	lazygit_cmd_tar := exec.Command("curl", "-L", lazygit_tar_url, "-o", "package.tar.gz")
	err := lazygit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lazygit_zip_url := "https://github.com/jesseduffield/lazygit/archive/refs/tags/v0.43.1.zip"
	lazygit_cmd_zip := exec.Command("curl", "-L", lazygit_zip_url, "-o", "package.zip")
	err = lazygit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lazygit_bin_url := "https://github.com/jesseduffield/lazygit/archive/refs/tags/v0.43.1.bin"
	lazygit_cmd_bin := exec.Command("curl", "-L", lazygit_bin_url, "-o", "binary.bin")
	err = lazygit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lazygit_src_url := "https://github.com/jesseduffield/lazygit/archive/refs/tags/v0.43.1.src.tar.gz"
	lazygit_cmd_src := exec.Command("curl", "-L", lazygit_src_url, "-o", "source.tar.gz")
	err = lazygit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lazygit_cmd_direct := exec.Command("./binary")
	err = lazygit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
