package main

// PyenvVirtualenvwrapper - Alternative to pyenv for managing virtualenvs
// Homepage: https://github.com/pyenv/pyenv-virtualenvwrapper

import (
	"fmt"
	
	"os/exec"
)

func installPyenvVirtualenvwrapper() {
	// Método 1: Descargar y extraer .tar.gz
	pyenvvirtualenvwrapper_tar_url := "https://github.com/pyenv/pyenv-virtualenvwrapper/archive/refs/tags/v20140609.tar.gz"
	pyenvvirtualenvwrapper_cmd_tar := exec.Command("curl", "-L", pyenvvirtualenvwrapper_tar_url, "-o", "package.tar.gz")
	err := pyenvvirtualenvwrapper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyenvvirtualenvwrapper_zip_url := "https://github.com/pyenv/pyenv-virtualenvwrapper/archive/refs/tags/v20140609.zip"
	pyenvvirtualenvwrapper_cmd_zip := exec.Command("curl", "-L", pyenvvirtualenvwrapper_zip_url, "-o", "package.zip")
	err = pyenvvirtualenvwrapper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyenvvirtualenvwrapper_bin_url := "https://github.com/pyenv/pyenv-virtualenvwrapper/archive/refs/tags/v20140609.bin"
	pyenvvirtualenvwrapper_cmd_bin := exec.Command("curl", "-L", pyenvvirtualenvwrapper_bin_url, "-o", "binary.bin")
	err = pyenvvirtualenvwrapper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyenvvirtualenvwrapper_src_url := "https://github.com/pyenv/pyenv-virtualenvwrapper/archive/refs/tags/v20140609.src.tar.gz"
	pyenvvirtualenvwrapper_cmd_src := exec.Command("curl", "-L", pyenvvirtualenvwrapper_src_url, "-o", "source.tar.gz")
	err = pyenvvirtualenvwrapper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyenvvirtualenvwrapper_cmd_direct := exec.Command("./binary")
	err = pyenvvirtualenvwrapper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyenv")
	exec.Command("latte", "install", "pyenv").Run()
}
