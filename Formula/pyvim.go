package main

// Pyvim - Pure Python Vim clone
// Homepage: https://github.com/prompt-toolkit/pyvim

import (
	"fmt"
	
	"os/exec"
)

func installPyvim() {
	// Método 1: Descargar y extraer .tar.gz
	pyvim_tar_url := "https://files.pythonhosted.org/packages/c3/31/04e144ec3a3a0303e3ef1ef9c6c1ec8a3b5ba9e88b98d21442d9152783c1/pyvim-3.0.3.tar.gz"
	pyvim_cmd_tar := exec.Command("curl", "-L", pyvim_tar_url, "-o", "package.tar.gz")
	err := pyvim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyvim_zip_url := "https://files.pythonhosted.org/packages/c3/31/04e144ec3a3a0303e3ef1ef9c6c1ec8a3b5ba9e88b98d21442d9152783c1/pyvim-3.0.3.zip"
	pyvim_cmd_zip := exec.Command("curl", "-L", pyvim_zip_url, "-o", "package.zip")
	err = pyvim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyvim_bin_url := "https://files.pythonhosted.org/packages/c3/31/04e144ec3a3a0303e3ef1ef9c6c1ec8a3b5ba9e88b98d21442d9152783c1/pyvim-3.0.3.bin"
	pyvim_cmd_bin := exec.Command("curl", "-L", pyvim_bin_url, "-o", "binary.bin")
	err = pyvim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyvim_src_url := "https://files.pythonhosted.org/packages/c3/31/04e144ec3a3a0303e3ef1ef9c6c1ec8a3b5ba9e88b98d21442d9152783c1/pyvim-3.0.3.src.tar.gz"
	pyvim_cmd_src := exec.Command("curl", "-L", pyvim_src_url, "-o", "source.tar.gz")
	err = pyvim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyvim_cmd_direct := exec.Command("./binary")
	err = pyvim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
