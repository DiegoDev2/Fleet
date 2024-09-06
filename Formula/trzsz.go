package main

// Trzsz - Simple file transfer tools, similar to lrzsz (rz/sz), and compatible with tmux
// Homepage: https://trzsz.github.io

import (
	"fmt"
	
	"os/exec"
)

func installTrzsz() {
	// Método 1: Descargar y extraer .tar.gz
	trzsz_tar_url := "https://files.pythonhosted.org/packages/22/1e/40a495c84a0dc625a4d97638c5cae308306718c493f480ee5ac64801947b/trzsz-1.1.5.tar.gz"
	trzsz_cmd_tar := exec.Command("curl", "-L", trzsz_tar_url, "-o", "package.tar.gz")
	err := trzsz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trzsz_zip_url := "https://files.pythonhosted.org/packages/22/1e/40a495c84a0dc625a4d97638c5cae308306718c493f480ee5ac64801947b/trzsz-1.1.5.zip"
	trzsz_cmd_zip := exec.Command("curl", "-L", trzsz_zip_url, "-o", "package.zip")
	err = trzsz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trzsz_bin_url := "https://files.pythonhosted.org/packages/22/1e/40a495c84a0dc625a4d97638c5cae308306718c493f480ee5ac64801947b/trzsz-1.1.5.bin"
	trzsz_cmd_bin := exec.Command("curl", "-L", trzsz_bin_url, "-o", "binary.bin")
	err = trzsz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trzsz_src_url := "https://files.pythonhosted.org/packages/22/1e/40a495c84a0dc625a4d97638c5cae308306718c493f480ee5ac64801947b/trzsz-1.1.5.src.tar.gz"
	trzsz_cmd_src := exec.Command("curl", "-L", trzsz_src_url, "-o", "source.tar.gz")
	err = trzsz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trzsz_cmd_direct := exec.Command("./binary")
	err = trzsz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
