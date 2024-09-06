package main

// GitCola - Highly caffeinated git GUI
// Homepage: https://git-cola.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installGitCola() {
	// Método 1: Descargar y extraer .tar.gz
	gitcola_tar_url := "https://files.pythonhosted.org/packages/f0/9b/81c8279f27c52aabfcf92206c9b5124b36f55d14558d9126e97fb1c7a6da/git-cola-4.8.2.tar.gz"
	gitcola_cmd_tar := exec.Command("curl", "-L", gitcola_tar_url, "-o", "package.tar.gz")
	err := gitcola_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitcola_zip_url := "https://files.pythonhosted.org/packages/f0/9b/81c8279f27c52aabfcf92206c9b5124b36f55d14558d9126e97fb1c7a6da/git-cola-4.8.2.zip"
	gitcola_cmd_zip := exec.Command("curl", "-L", gitcola_zip_url, "-o", "package.zip")
	err = gitcola_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitcola_bin_url := "https://files.pythonhosted.org/packages/f0/9b/81c8279f27c52aabfcf92206c9b5124b36f55d14558d9126e97fb1c7a6da/git-cola-4.8.2.bin"
	gitcola_cmd_bin := exec.Command("curl", "-L", gitcola_bin_url, "-o", "binary.bin")
	err = gitcola_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitcola_src_url := "https://files.pythonhosted.org/packages/f0/9b/81c8279f27c52aabfcf92206c9b5124b36f55d14558d9126e97fb1c7a6da/git-cola-4.8.2.src.tar.gz"
	gitcola_cmd_src := exec.Command("curl", "-L", gitcola_src_url, "-o", "source.tar.gz")
	err = gitcola_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitcola_cmd_direct := exec.Command("./binary")
	err = gitcola_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyqt")
exec.Command("latte", "install", "pyqt")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
