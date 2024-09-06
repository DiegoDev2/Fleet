package main

// Virtualenv - Tool for creating isolated virtual python environments
// Homepage: https://virtualenv.pypa.io/

import (
	"fmt"
	
	"os/exec"
)

func installVirtualenv() {
	// Método 1: Descargar y extraer .tar.gz
	virtualenv_tar_url := "https://files.pythonhosted.org/packages/68/60/db9f95e6ad456f1872486769c55628c7901fb4de5a72c2f7bdd912abf0c1/virtualenv-20.26.3.tar.gz"
	virtualenv_cmd_tar := exec.Command("curl", "-L", virtualenv_tar_url, "-o", "package.tar.gz")
	err := virtualenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virtualenv_zip_url := "https://files.pythonhosted.org/packages/68/60/db9f95e6ad456f1872486769c55628c7901fb4de5a72c2f7bdd912abf0c1/virtualenv-20.26.3.zip"
	virtualenv_cmd_zip := exec.Command("curl", "-L", virtualenv_zip_url, "-o", "package.zip")
	err = virtualenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virtualenv_bin_url := "https://files.pythonhosted.org/packages/68/60/db9f95e6ad456f1872486769c55628c7901fb4de5a72c2f7bdd912abf0c1/virtualenv-20.26.3.bin"
	virtualenv_cmd_bin := exec.Command("curl", "-L", virtualenv_bin_url, "-o", "binary.bin")
	err = virtualenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virtualenv_src_url := "https://files.pythonhosted.org/packages/68/60/db9f95e6ad456f1872486769c55628c7901fb4de5a72c2f7bdd912abf0c1/virtualenv-20.26.3.src.tar.gz"
	virtualenv_cmd_src := exec.Command("curl", "-L", virtualenv_src_url, "-o", "source.tar.gz")
	err = virtualenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virtualenv_cmd_direct := exec.Command("./binary")
	err = virtualenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
