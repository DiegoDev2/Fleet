package main

// Dotbot - Tool that bootstraps your dotfiles
// Homepage: https://github.com/anishathalye/dotbot

import (
	"fmt"
	
	"os/exec"
)

func installDotbot() {
	// Método 1: Descargar y extraer .tar.gz
	dotbot_tar_url := "https://files.pythonhosted.org/packages/04/8b/0899638625ff6443b627294b10f3fa95b84da330d7caf9936ba991baf504/dotbot-1.20.1.tar.gz"
	dotbot_cmd_tar := exec.Command("curl", "-L", dotbot_tar_url, "-o", "package.tar.gz")
	err := dotbot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dotbot_zip_url := "https://files.pythonhosted.org/packages/04/8b/0899638625ff6443b627294b10f3fa95b84da330d7caf9936ba991baf504/dotbot-1.20.1.zip"
	dotbot_cmd_zip := exec.Command("curl", "-L", dotbot_zip_url, "-o", "package.zip")
	err = dotbot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dotbot_bin_url := "https://files.pythonhosted.org/packages/04/8b/0899638625ff6443b627294b10f3fa95b84da330d7caf9936ba991baf504/dotbot-1.20.1.bin"
	dotbot_cmd_bin := exec.Command("curl", "-L", dotbot_bin_url, "-o", "binary.bin")
	err = dotbot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dotbot_src_url := "https://files.pythonhosted.org/packages/04/8b/0899638625ff6443b627294b10f3fa95b84da330d7caf9936ba991baf504/dotbot-1.20.1.src.tar.gz"
	dotbot_cmd_src := exec.Command("curl", "-L", dotbot_src_url, "-o", "source.tar.gz")
	err = dotbot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dotbot_cmd_direct := exec.Command("./binary")
	err = dotbot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
