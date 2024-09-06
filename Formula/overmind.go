package main

// Overmind - Process manager for Procfile-based applications and tmux
// Homepage: https://github.com/DarthSim/overmind

import (
	"fmt"
	
	"os/exec"
)

func installOvermind() {
	// Método 1: Descargar y extraer .tar.gz
	overmind_tar_url := "https://github.com/DarthSim/overmind/archive/refs/tags/v2.5.1.tar.gz"
	overmind_cmd_tar := exec.Command("curl", "-L", overmind_tar_url, "-o", "package.tar.gz")
	err := overmind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	overmind_zip_url := "https://github.com/DarthSim/overmind/archive/refs/tags/v2.5.1.zip"
	overmind_cmd_zip := exec.Command("curl", "-L", overmind_zip_url, "-o", "package.zip")
	err = overmind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	overmind_bin_url := "https://github.com/DarthSim/overmind/archive/refs/tags/v2.5.1.bin"
	overmind_cmd_bin := exec.Command("curl", "-L", overmind_bin_url, "-o", "binary.bin")
	err = overmind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	overmind_src_url := "https://github.com/DarthSim/overmind/archive/refs/tags/v2.5.1.src.tar.gz"
	overmind_cmd_src := exec.Command("curl", "-L", overmind_src_url, "-o", "source.tar.gz")
	err = overmind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	overmind_cmd_direct := exec.Command("./binary")
	err = overmind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: tmux")
	exec.Command("latte", "install", "tmux").Run()
}
