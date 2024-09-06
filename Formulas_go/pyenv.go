package main

// Pyenv - Python version management
// Homepage: https://github.com/pyenv/pyenv

import (
	"fmt"
	
	"os/exec"
)

func installPyenv() {
	// Método 1: Descargar y extraer .tar.gz
	pyenv_tar_url := "https://github.com/pyenv/pyenv/archive/refs/tags/v2.4.11.tar.gz"
	pyenv_cmd_tar := exec.Command("curl", "-L", pyenv_tar_url, "-o", "package.tar.gz")
	err := pyenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyenv_zip_url := "https://github.com/pyenv/pyenv/archive/refs/tags/v2.4.11.zip"
	pyenv_cmd_zip := exec.Command("curl", "-L", pyenv_zip_url, "-o", "package.zip")
	err = pyenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyenv_bin_url := "https://github.com/pyenv/pyenv/archive/refs/tags/v2.4.11.bin"
	pyenv_cmd_bin := exec.Command("curl", "-L", pyenv_bin_url, "-o", "binary.bin")
	err = pyenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyenv_src_url := "https://github.com/pyenv/pyenv/archive/refs/tags/v2.4.11.src.tar.gz"
	pyenv_cmd_src := exec.Command("curl", "-L", pyenv_src_url, "-o", "source.tar.gz")
	err = pyenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyenv_cmd_direct := exec.Command("./binary")
	err = pyenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
