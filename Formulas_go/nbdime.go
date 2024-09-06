package main

// Nbdime - Jupyter Notebook Diff and Merge tools
// Homepage: https://nbdime.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installNbdime() {
	// Método 1: Descargar y extraer .tar.gz
	nbdime_tar_url := "https://files.pythonhosted.org/packages/a6/f1/4be57ecea4d55d322f05a0f89e0b73d7a8d90a16dbf01168eab3e7bf5939/nbdime-4.0.2.tar.gz"
	nbdime_cmd_tar := exec.Command("curl", "-L", nbdime_tar_url, "-o", "package.tar.gz")
	err := nbdime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nbdime_zip_url := "https://files.pythonhosted.org/packages/a6/f1/4be57ecea4d55d322f05a0f89e0b73d7a8d90a16dbf01168eab3e7bf5939/nbdime-4.0.2.zip"
	nbdime_cmd_zip := exec.Command("curl", "-L", nbdime_zip_url, "-o", "package.zip")
	err = nbdime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nbdime_bin_url := "https://files.pythonhosted.org/packages/a6/f1/4be57ecea4d55d322f05a0f89e0b73d7a8d90a16dbf01168eab3e7bf5939/nbdime-4.0.2.bin"
	nbdime_cmd_bin := exec.Command("curl", "-L", nbdime_bin_url, "-o", "binary.bin")
	err = nbdime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nbdime_src_url := "https://files.pythonhosted.org/packages/a6/f1/4be57ecea4d55d322f05a0f89e0b73d7a8d90a16dbf01168eab3e7bf5939/nbdime-4.0.2.src.tar.gz"
	nbdime_cmd_src := exec.Command("curl", "-L", nbdime_src_url, "-o", "source.tar.gz")
	err = nbdime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nbdime_cmd_direct := exec.Command("./binary")
	err = nbdime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: jupyterlab")
exec.Command("latte", "install", "jupyterlab")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
