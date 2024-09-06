package main

// Sgr - Command-line client for Splitgraph, a version control system for data
// Homepage: https://www.splitgraph.com/docs/sgr-advanced/getting-started/introduction

import (
	"fmt"
	
	"os/exec"
)

func installSgr() {
	// Método 1: Descargar y extraer .tar.gz
	sgr_tar_url := "https://files.pythonhosted.org/packages/dd/61/7d6cf822edb39d2426f6f185c7fc4de0ad4b80e0da3e5f50d94952795c11/splitgraph-0.3.12.tar.gz"
	sgr_cmd_tar := exec.Command("curl", "-L", sgr_tar_url, "-o", "package.tar.gz")
	err := sgr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sgr_zip_url := "https://files.pythonhosted.org/packages/dd/61/7d6cf822edb39d2426f6f185c7fc4de0ad4b80e0da3e5f50d94952795c11/splitgraph-0.3.12.zip"
	sgr_cmd_zip := exec.Command("curl", "-L", sgr_zip_url, "-o", "package.zip")
	err = sgr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sgr_bin_url := "https://files.pythonhosted.org/packages/dd/61/7d6cf822edb39d2426f6f185c7fc4de0ad4b80e0da3e5f50d94952795c11/splitgraph-0.3.12.bin"
	sgr_cmd_bin := exec.Command("curl", "-L", sgr_bin_url, "-o", "binary.bin")
	err = sgr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sgr_src_url := "https://files.pythonhosted.org/packages/dd/61/7d6cf822edb39d2426f6f185c7fc4de0ad4b80e0da3e5f50d94952795c11/splitgraph-0.3.12.src.tar.gz"
	sgr_cmd_src := exec.Command("curl", "-L", sgr_src_url, "-o", "source.tar.gz")
	err = sgr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sgr_cmd_direct := exec.Command("./binary")
	err = sgr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cython")
	exec.Command("latte", "install", "cython").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
