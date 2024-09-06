package main

// ConanAT1 - Distributed, open source, package manager for C/C++
// Homepage: https://conan.io

import (
	"fmt"
	
	"os/exec"
)

func installConanAT1() {
	// Método 1: Descargar y extraer .tar.gz
	conanat1_tar_url := "https://files.pythonhosted.org/packages/7b/00/40fa960783926ccc478707cc294f645636df5ecd7eb29edc87f6280a4898/conan-1.65.0.tar.gz"
	conanat1_cmd_tar := exec.Command("curl", "-L", conanat1_tar_url, "-o", "package.tar.gz")
	err := conanat1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	conanat1_zip_url := "https://files.pythonhosted.org/packages/7b/00/40fa960783926ccc478707cc294f645636df5ecd7eb29edc87f6280a4898/conan-1.65.0.zip"
	conanat1_cmd_zip := exec.Command("curl", "-L", conanat1_zip_url, "-o", "package.zip")
	err = conanat1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	conanat1_bin_url := "https://files.pythonhosted.org/packages/7b/00/40fa960783926ccc478707cc294f645636df5ecd7eb29edc87f6280a4898/conan-1.65.0.bin"
	conanat1_cmd_bin := exec.Command("curl", "-L", conanat1_bin_url, "-o", "binary.bin")
	err = conanat1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	conanat1_src_url := "https://files.pythonhosted.org/packages/7b/00/40fa960783926ccc478707cc294f645636df5ecd7eb29edc87f6280a4898/conan-1.65.0.src.tar.gz"
	conanat1_cmd_src := exec.Command("curl", "-L", conanat1_src_url, "-o", "source.tar.gz")
	err = conanat1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	conanat1_cmd_direct := exec.Command("./binary")
	err = conanat1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
