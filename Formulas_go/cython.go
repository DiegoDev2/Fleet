package main

// Cython - Compiler for writing C extensions for the Python language
// Homepage: https://cython.org/

import (
	"fmt"
	
	"os/exec"
)

func installCython() {
	// Método 1: Descargar y extraer .tar.gz
	cython_tar_url := "https://files.pythonhosted.org/packages/84/4d/b720d6000f4ca77f030bd70f12550820f0766b568e43f11af7f7ad9061aa/cython-3.0.11.tar.gz"
	cython_cmd_tar := exec.Command("curl", "-L", cython_tar_url, "-o", "package.tar.gz")
	err := cython_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cython_zip_url := "https://files.pythonhosted.org/packages/84/4d/b720d6000f4ca77f030bd70f12550820f0766b568e43f11af7f7ad9061aa/cython-3.0.11.zip"
	cython_cmd_zip := exec.Command("curl", "-L", cython_zip_url, "-o", "package.zip")
	err = cython_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cython_bin_url := "https://files.pythonhosted.org/packages/84/4d/b720d6000f4ca77f030bd70f12550820f0766b568e43f11af7f7ad9061aa/cython-3.0.11.bin"
	cython_cmd_bin := exec.Command("curl", "-L", cython_bin_url, "-o", "binary.bin")
	err = cython_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cython_src_url := "https://files.pythonhosted.org/packages/84/4d/b720d6000f4ca77f030bd70f12550820f0766b568e43f11af7f7ad9061aa/cython-3.0.11.src.tar.gz"
	cython_cmd_src := exec.Command("curl", "-L", cython_src_url, "-o", "source.tar.gz")
	err = cython_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cython_cmd_direct := exec.Command("./binary")
	err = cython_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
