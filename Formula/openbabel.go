package main

// OpenBabel - Chemical toolbox
// Homepage: https://github.com/openbabel/openbabel

import (
	"fmt"
	
	"os/exec"
)

func installOpenBabel() {
	// Método 1: Descargar y extraer .tar.gz
	openbabel_tar_url := "https://github.com/openbabel/openbabel/archive/refs/tags/openbabel-3-1-1.tar.gz"
	openbabel_cmd_tar := exec.Command("curl", "-L", openbabel_tar_url, "-o", "package.tar.gz")
	err := openbabel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openbabel_zip_url := "https://github.com/openbabel/openbabel/archive/refs/tags/openbabel-3-1-1.zip"
	openbabel_cmd_zip := exec.Command("curl", "-L", openbabel_zip_url, "-o", "package.zip")
	err = openbabel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openbabel_bin_url := "https://github.com/openbabel/openbabel/archive/refs/tags/openbabel-3-1-1.bin"
	openbabel_cmd_bin := exec.Command("curl", "-L", openbabel_bin_url, "-o", "binary.bin")
	err = openbabel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openbabel_src_url := "https://github.com/openbabel/openbabel/archive/refs/tags/openbabel-3-1-1.src.tar.gz"
	openbabel_cmd_src := exec.Command("curl", "-L", openbabel_src_url, "-o", "source.tar.gz")
	err = openbabel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openbabel_cmd_direct := exec.Command("./binary")
	err = openbabel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: rapidjson")
	exec.Command("latte", "install", "rapidjson").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
