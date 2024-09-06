package main

// Cffi - C Foreign Function Interface for Python
// Homepage: https://cffi.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installCffi() {
	// Método 1: Descargar y extraer .tar.gz
	cffi_tar_url := "https://files.pythonhosted.org/packages/fc/97/c783634659c2920c3fc70419e3af40972dbaf758daa229a7d6ea6135c90d/cffi-1.17.1.tar.gz"
	cffi_cmd_tar := exec.Command("curl", "-L", cffi_tar_url, "-o", "package.tar.gz")
	err := cffi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cffi_zip_url := "https://files.pythonhosted.org/packages/fc/97/c783634659c2920c3fc70419e3af40972dbaf758daa229a7d6ea6135c90d/cffi-1.17.1.zip"
	cffi_cmd_zip := exec.Command("curl", "-L", cffi_zip_url, "-o", "package.zip")
	err = cffi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cffi_bin_url := "https://files.pythonhosted.org/packages/fc/97/c783634659c2920c3fc70419e3af40972dbaf758daa229a7d6ea6135c90d/cffi-1.17.1.bin"
	cffi_cmd_bin := exec.Command("curl", "-L", cffi_bin_url, "-o", "binary.bin")
	err = cffi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cffi_src_url := "https://files.pythonhosted.org/packages/fc/97/c783634659c2920c3fc70419e3af40972dbaf758daa229a7d6ea6135c90d/cffi-1.17.1.src.tar.gz"
	cffi_cmd_src := exec.Command("curl", "-L", cffi_src_url, "-o", "source.tar.gz")
	err = cffi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cffi_cmd_direct := exec.Command("./binary")
	err = cffi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pycparser")
	exec.Command("latte", "install", "pycparser").Run()
}
