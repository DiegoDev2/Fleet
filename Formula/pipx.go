package main

// Pipx - Execute binaries from Python packages in isolated environments
// Homepage: https://pipx.pypa.io

import (
	"fmt"
	
	"os/exec"
)

func installPipx() {
	// Método 1: Descargar y extraer .tar.gz
	pipx_tar_url := "https://files.pythonhosted.org/packages/17/21/dd6b9a9c4f0cb659ce3dad991f0e8dde852b2c81922224ef77df4222ab7a/pipx-1.7.1.tar.gz"
	pipx_cmd_tar := exec.Command("curl", "-L", pipx_tar_url, "-o", "package.tar.gz")
	err := pipx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipx_zip_url := "https://files.pythonhosted.org/packages/17/21/dd6b9a9c4f0cb659ce3dad991f0e8dde852b2c81922224ef77df4222ab7a/pipx-1.7.1.zip"
	pipx_cmd_zip := exec.Command("curl", "-L", pipx_zip_url, "-o", "package.zip")
	err = pipx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipx_bin_url := "https://files.pythonhosted.org/packages/17/21/dd6b9a9c4f0cb659ce3dad991f0e8dde852b2c81922224ef77df4222ab7a/pipx-1.7.1.bin"
	pipx_cmd_bin := exec.Command("curl", "-L", pipx_bin_url, "-o", "binary.bin")
	err = pipx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipx_src_url := "https://files.pythonhosted.org/packages/17/21/dd6b9a9c4f0cb659ce3dad991f0e8dde852b2c81922224ef77df4222ab7a/pipx-1.7.1.src.tar.gz"
	pipx_cmd_src := exec.Command("curl", "-L", pipx_src_url, "-o", "source.tar.gz")
	err = pipx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipx_cmd_direct := exec.Command("./binary")
	err = pipx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
