package main

// Tmuxp - Tmux session manager. Built on libtmux
// Homepage: https://tmuxp.git-pull.com/

import (
	"fmt"
	
	"os/exec"
)

func installTmuxp() {
	// Método 1: Descargar y extraer .tar.gz
	tmuxp_tar_url := "https://files.pythonhosted.org/packages/2e/5b/66ceb26459b9324bbd425c673170172066f08a6c180c8da166de6a44cf45/tmuxp-1.47.0.tar.gz"
	tmuxp_cmd_tar := exec.Command("curl", "-L", tmuxp_tar_url, "-o", "package.tar.gz")
	err := tmuxp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmuxp_zip_url := "https://files.pythonhosted.org/packages/2e/5b/66ceb26459b9324bbd425c673170172066f08a6c180c8da166de6a44cf45/tmuxp-1.47.0.zip"
	tmuxp_cmd_zip := exec.Command("curl", "-L", tmuxp_zip_url, "-o", "package.zip")
	err = tmuxp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmuxp_bin_url := "https://files.pythonhosted.org/packages/2e/5b/66ceb26459b9324bbd425c673170172066f08a6c180c8da166de6a44cf45/tmuxp-1.47.0.bin"
	tmuxp_cmd_bin := exec.Command("curl", "-L", tmuxp_bin_url, "-o", "binary.bin")
	err = tmuxp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmuxp_src_url := "https://files.pythonhosted.org/packages/2e/5b/66ceb26459b9324bbd425c673170172066f08a6c180c8da166de6a44cf45/tmuxp-1.47.0.src.tar.gz"
	tmuxp_cmd_src := exec.Command("curl", "-L", tmuxp_src_url, "-o", "source.tar.gz")
	err = tmuxp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmuxp_cmd_direct := exec.Command("./binary")
	err = tmuxp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: tmux")
	exec.Command("latte", "install", "tmux").Run()
}
