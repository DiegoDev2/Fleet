package main

// Nox - Flexible test automation for Python
// Homepage: https://nox.thea.codes/

import (
	"fmt"
	
	"os/exec"
)

func installNox() {
	// Método 1: Descargar y extraer .tar.gz
	nox_tar_url := "https://files.pythonhosted.org/packages/1e/86/b86fc26784d2f63d038b4efc9e18d4d807ec025569da66c6d032b8f717df/nox-2024.4.15.tar.gz"
	nox_cmd_tar := exec.Command("curl", "-L", nox_tar_url, "-o", "package.tar.gz")
	err := nox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nox_zip_url := "https://files.pythonhosted.org/packages/1e/86/b86fc26784d2f63d038b4efc9e18d4d807ec025569da66c6d032b8f717df/nox-2024.4.15.zip"
	nox_cmd_zip := exec.Command("curl", "-L", nox_zip_url, "-o", "package.zip")
	err = nox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nox_bin_url := "https://files.pythonhosted.org/packages/1e/86/b86fc26784d2f63d038b4efc9e18d4d807ec025569da66c6d032b8f717df/nox-2024.4.15.bin"
	nox_cmd_bin := exec.Command("curl", "-L", nox_bin_url, "-o", "binary.bin")
	err = nox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nox_src_url := "https://files.pythonhosted.org/packages/1e/86/b86fc26784d2f63d038b4efc9e18d4d807ec025569da66c6d032b8f717df/nox-2024.4.15.src.tar.gz"
	nox_cmd_src := exec.Command("curl", "-L", nox_src_url, "-o", "source.tar.gz")
	err = nox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nox_cmd_direct := exec.Command("./binary")
	err = nox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
