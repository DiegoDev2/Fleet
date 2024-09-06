package main

// Pygit2 - Bindings to the libgit2 shared library
// Homepage: https://github.com/libgit2/pygit2

import (
	"fmt"
	
	"os/exec"
)

func installPygit2() {
	// Método 1: Descargar y extraer .tar.gz
	pygit2_tar_url := "https://files.pythonhosted.org/packages/53/77/d33e2c619478d0daea4a50f9ffdd588db2ca55817c7e9a6c796fca3b80ef/pygit2-1.15.1.tar.gz"
	pygit2_cmd_tar := exec.Command("curl", "-L", pygit2_tar_url, "-o", "package.tar.gz")
	err := pygit2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pygit2_zip_url := "https://files.pythonhosted.org/packages/53/77/d33e2c619478d0daea4a50f9ffdd588db2ca55817c7e9a6c796fca3b80ef/pygit2-1.15.1.zip"
	pygit2_cmd_zip := exec.Command("curl", "-L", pygit2_zip_url, "-o", "package.zip")
	err = pygit2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pygit2_bin_url := "https://files.pythonhosted.org/packages/53/77/d33e2c619478d0daea4a50f9ffdd588db2ca55817c7e9a6c796fca3b80ef/pygit2-1.15.1.bin"
	pygit2_cmd_bin := exec.Command("curl", "-L", pygit2_bin_url, "-o", "binary.bin")
	err = pygit2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pygit2_src_url := "https://files.pythonhosted.org/packages/53/77/d33e2c619478d0daea4a50f9ffdd588db2ca55817c7e9a6c796fca3b80ef/pygit2-1.15.1.src.tar.gz"
	pygit2_cmd_src := exec.Command("curl", "-L", pygit2_src_url, "-o", "source.tar.gz")
	err = pygit2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pygit2_cmd_direct := exec.Command("./binary")
	err = pygit2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: cffi")
	exec.Command("latte", "install", "cffi").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
}
