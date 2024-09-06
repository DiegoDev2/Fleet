package main

// Coconut - Simple, elegant, Pythonic functional programming
// Homepage: https://coconut-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installCoconut() {
	// Método 1: Descargar y extraer .tar.gz
	coconut_tar_url := "https://files.pythonhosted.org/packages/93/75/414f33186846444da53b4e834d5ccfb0577d0e09b997819c183fa509f70a/coconut-3.1.2.tar.gz"
	coconut_cmd_tar := exec.Command("curl", "-L", coconut_tar_url, "-o", "package.tar.gz")
	err := coconut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coconut_zip_url := "https://files.pythonhosted.org/packages/93/75/414f33186846444da53b4e834d5ccfb0577d0e09b997819c183fa509f70a/coconut-3.1.2.zip"
	coconut_cmd_zip := exec.Command("curl", "-L", coconut_zip_url, "-o", "package.zip")
	err = coconut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coconut_bin_url := "https://files.pythonhosted.org/packages/93/75/414f33186846444da53b4e834d5ccfb0577d0e09b997819c183fa509f70a/coconut-3.1.2.bin"
	coconut_cmd_bin := exec.Command("curl", "-L", coconut_bin_url, "-o", "binary.bin")
	err = coconut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coconut_src_url := "https://files.pythonhosted.org/packages/93/75/414f33186846444da53b4e834d5ccfb0577d0e09b997819c183fa509f70a/coconut-3.1.2.src.tar.gz"
	coconut_cmd_src := exec.Command("curl", "-L", coconut_src_url, "-o", "source.tar.gz")
	err = coconut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coconut_cmd_direct := exec.Command("./binary")
	err = coconut_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
