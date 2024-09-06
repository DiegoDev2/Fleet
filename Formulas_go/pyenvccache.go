package main

// PyenvCcache - Make Python build faster, using the leverage of `ccache`
// Homepage: https://github.com/pyenv/pyenv-ccache

import (
	"fmt"
	
	"os/exec"
)

func installPyenvCcache() {
	// Método 1: Descargar y extraer .tar.gz
	pyenvccache_tar_url := "https://github.com/pyenv/pyenv-ccache/archive/refs/tags/v0.0.2.tar.gz"
	pyenvccache_cmd_tar := exec.Command("curl", "-L", pyenvccache_tar_url, "-o", "package.tar.gz")
	err := pyenvccache_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyenvccache_zip_url := "https://github.com/pyenv/pyenv-ccache/archive/refs/tags/v0.0.2.zip"
	pyenvccache_cmd_zip := exec.Command("curl", "-L", pyenvccache_zip_url, "-o", "package.zip")
	err = pyenvccache_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyenvccache_bin_url := "https://github.com/pyenv/pyenv-ccache/archive/refs/tags/v0.0.2.bin"
	pyenvccache_cmd_bin := exec.Command("curl", "-L", pyenvccache_bin_url, "-o", "binary.bin")
	err = pyenvccache_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyenvccache_src_url := "https://github.com/pyenv/pyenv-ccache/archive/refs/tags/v0.0.2.src.tar.gz"
	pyenvccache_cmd_src := exec.Command("curl", "-L", pyenvccache_src_url, "-o", "source.tar.gz")
	err = pyenvccache_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyenvccache_cmd_direct := exec.Command("./binary")
	err = pyenvccache_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ccache")
exec.Command("latte", "install", "ccache")
	fmt.Println("Instalando dependencia: pyenv")
exec.Command("latte", "install", "pyenv")
}
