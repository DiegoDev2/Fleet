package main

// Bumpversion - Increase version numbers with SemVer terms
// Homepage: https://pypi.python.org/pypi/bumpversion

import (
	"fmt"
	
	"os/exec"
)

func installBumpversion() {
	// Método 1: Descargar y extraer .tar.gz
	bumpversion_tar_url := "https://files.pythonhosted.org/packages/29/2a/688aca6eeebfe8941235be53f4da780c6edee05dbbea5d7abaa3aab6fad2/bump2version-1.0.1.tar.gz"
	bumpversion_cmd_tar := exec.Command("curl", "-L", bumpversion_tar_url, "-o", "package.tar.gz")
	err := bumpversion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bumpversion_zip_url := "https://files.pythonhosted.org/packages/29/2a/688aca6eeebfe8941235be53f4da780c6edee05dbbea5d7abaa3aab6fad2/bump2version-1.0.1.zip"
	bumpversion_cmd_zip := exec.Command("curl", "-L", bumpversion_zip_url, "-o", "package.zip")
	err = bumpversion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bumpversion_bin_url := "https://files.pythonhosted.org/packages/29/2a/688aca6eeebfe8941235be53f4da780c6edee05dbbea5d7abaa3aab6fad2/bump2version-1.0.1.bin"
	bumpversion_cmd_bin := exec.Command("curl", "-L", bumpversion_bin_url, "-o", "binary.bin")
	err = bumpversion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bumpversion_src_url := "https://files.pythonhosted.org/packages/29/2a/688aca6eeebfe8941235be53f4da780c6edee05dbbea5d7abaa3aab6fad2/bump2version-1.0.1.src.tar.gz"
	bumpversion_cmd_src := exec.Command("curl", "-L", bumpversion_src_url, "-o", "source.tar.gz")
	err = bumpversion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bumpversion_cmd_direct := exec.Command("./binary")
	err = bumpversion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
