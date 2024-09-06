package main

// Conan - Distributed, open source, package manager for C/C++
// Homepage: https://conan.io

import (
	"fmt"
	
	"os/exec"
)

func installConan() {
	// Método 1: Descargar y extraer .tar.gz
	conan_tar_url := "https://files.pythonhosted.org/packages/bd/8d/1477148f54e3a57e30a96f17edadad7279d1740a0094edb7a5ae904d0a94/conan-2.7.0.tar.gz"
	conan_cmd_tar := exec.Command("curl", "-L", conan_tar_url, "-o", "package.tar.gz")
	err := conan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	conan_zip_url := "https://files.pythonhosted.org/packages/bd/8d/1477148f54e3a57e30a96f17edadad7279d1740a0094edb7a5ae904d0a94/conan-2.7.0.zip"
	conan_cmd_zip := exec.Command("curl", "-L", conan_zip_url, "-o", "package.zip")
	err = conan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	conan_bin_url := "https://files.pythonhosted.org/packages/bd/8d/1477148f54e3a57e30a96f17edadad7279d1740a0094edb7a5ae904d0a94/conan-2.7.0.bin"
	conan_cmd_bin := exec.Command("curl", "-L", conan_bin_url, "-o", "binary.bin")
	err = conan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	conan_src_url := "https://files.pythonhosted.org/packages/bd/8d/1477148f54e3a57e30a96f17edadad7279d1740a0094edb7a5ae904d0a94/conan-2.7.0.src.tar.gz"
	conan_cmd_src := exec.Command("curl", "-L", conan_src_url, "-o", "source.tar.gz")
	err = conan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	conan_cmd_direct := exec.Command("./binary")
	err = conan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
