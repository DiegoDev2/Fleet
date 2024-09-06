package main

// Dnsviz - Tools for analyzing and visualizing DNS and DNSSEC behavior
// Homepage: https://github.com/dnsviz/dnsviz/

import (
	"fmt"
	
	"os/exec"
)

func installDnsviz() {
	// Método 1: Descargar y extraer .tar.gz
	dnsviz_tar_url := "https://files.pythonhosted.org/packages/cd/6e/8e285523108cc91b32f0584c2b4a7b006348af597cdc84e728206df15b3b/dnsviz-0.10.0.tar.gz"
	dnsviz_cmd_tar := exec.Command("curl", "-L", dnsviz_tar_url, "-o", "package.tar.gz")
	err := dnsviz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsviz_zip_url := "https://files.pythonhosted.org/packages/cd/6e/8e285523108cc91b32f0584c2b4a7b006348af597cdc84e728206df15b3b/dnsviz-0.10.0.zip"
	dnsviz_cmd_zip := exec.Command("curl", "-L", dnsviz_zip_url, "-o", "package.zip")
	err = dnsviz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsviz_bin_url := "https://files.pythonhosted.org/packages/cd/6e/8e285523108cc91b32f0584c2b4a7b006348af597cdc84e728206df15b3b/dnsviz-0.10.0.bin"
	dnsviz_cmd_bin := exec.Command("curl", "-L", dnsviz_bin_url, "-o", "binary.bin")
	err = dnsviz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsviz_src_url := "https://files.pythonhosted.org/packages/cd/6e/8e285523108cc91b32f0584c2b4a7b006348af597cdc84e728206df15b3b/dnsviz-0.10.0.src.tar.gz"
	dnsviz_cmd_src := exec.Command("curl", "-L", dnsviz_src_url, "-o", "source.tar.gz")
	err = dnsviz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsviz_cmd_direct := exec.Command("./binary")
	err = dnsviz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bind")
	exec.Command("latte", "install", "bind").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
}
