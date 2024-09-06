package main

// Cppman - C++ 98/11/14/17/20 manual pages from cplusplus.com and cppreference.com
// Homepage: https://github.com/aitjcize/cppman

import (
	"fmt"
	
	"os/exec"
)

func installCppman() {
	// Método 1: Descargar y extraer .tar.gz
	cppman_tar_url := "https://files.pythonhosted.org/packages/f7/c1/0ee5b360b7e5941fac6b3e4749e0f02c45154b1747f097ca925e8f605ea2/cppman-0.5.7.tar.gz"
	cppman_cmd_tar := exec.Command("curl", "-L", cppman_tar_url, "-o", "package.tar.gz")
	err := cppman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppman_zip_url := "https://files.pythonhosted.org/packages/f7/c1/0ee5b360b7e5941fac6b3e4749e0f02c45154b1747f097ca925e8f605ea2/cppman-0.5.7.zip"
	cppman_cmd_zip := exec.Command("curl", "-L", cppman_zip_url, "-o", "package.zip")
	err = cppman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppman_bin_url := "https://files.pythonhosted.org/packages/f7/c1/0ee5b360b7e5941fac6b3e4749e0f02c45154b1747f097ca925e8f605ea2/cppman-0.5.7.bin"
	cppman_cmd_bin := exec.Command("curl", "-L", cppman_bin_url, "-o", "binary.bin")
	err = cppman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppman_src_url := "https://files.pythonhosted.org/packages/f7/c1/0ee5b360b7e5941fac6b3e4749e0f02c45154b1747f097ca925e8f605ea2/cppman-0.5.7.src.tar.gz"
	cppman_cmd_src := exec.Command("curl", "-L", cppman_src_url, "-o", "source.tar.gz")
	err = cppman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppman_cmd_direct := exec.Command("./binary")
	err = cppman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: groff")
	exec.Command("latte", "install", "groff").Run()
}
