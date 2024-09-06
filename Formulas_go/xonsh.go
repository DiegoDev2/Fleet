package main

// Xonsh - Python-powered, cross-platform, Unix-gazing shell language and command prompt
// Homepage: https://xon.sh/

import (
	"fmt"
	
	"os/exec"
)

func installXonsh() {
	// Método 1: Descargar y extraer .tar.gz
	xonsh_tar_url := "https://files.pythonhosted.org/packages/73/82/c39c24a07daf22877f2ffa0c53dca5cd7c84ccc2d647b6fd6cd134d7f022/xonsh-0.18.3.tar.gz"
	xonsh_cmd_tar := exec.Command("curl", "-L", xonsh_tar_url, "-o", "package.tar.gz")
	err := xonsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xonsh_zip_url := "https://files.pythonhosted.org/packages/73/82/c39c24a07daf22877f2ffa0c53dca5cd7c84ccc2d647b6fd6cd134d7f022/xonsh-0.18.3.zip"
	xonsh_cmd_zip := exec.Command("curl", "-L", xonsh_zip_url, "-o", "package.zip")
	err = xonsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xonsh_bin_url := "https://files.pythonhosted.org/packages/73/82/c39c24a07daf22877f2ffa0c53dca5cd7c84ccc2d647b6fd6cd134d7f022/xonsh-0.18.3.bin"
	xonsh_cmd_bin := exec.Command("curl", "-L", xonsh_bin_url, "-o", "binary.bin")
	err = xonsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xonsh_src_url := "https://files.pythonhosted.org/packages/73/82/c39c24a07daf22877f2ffa0c53dca5cd7c84ccc2d647b6fd6cd134d7f022/xonsh-0.18.3.src.tar.gz"
	xonsh_cmd_src := exec.Command("curl", "-L", xonsh_src_url, "-o", "source.tar.gz")
	err = xonsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xonsh_cmd_direct := exec.Command("./binary")
	err = xonsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
