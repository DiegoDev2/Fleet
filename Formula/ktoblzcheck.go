package main

// Ktoblzcheck - Library for German banks
// Homepage: https://ktoblzcheck.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installKtoblzcheck() {
	// Método 1: Descargar y extraer .tar.gz
	ktoblzcheck_tar_url := "https://downloads.sourceforge.net/project/ktoblzcheck/ktoblzcheck-1.57.tar.gz"
	ktoblzcheck_cmd_tar := exec.Command("curl", "-L", ktoblzcheck_tar_url, "-o", "package.tar.gz")
	err := ktoblzcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ktoblzcheck_zip_url := "https://downloads.sourceforge.net/project/ktoblzcheck/ktoblzcheck-1.57.zip"
	ktoblzcheck_cmd_zip := exec.Command("curl", "-L", ktoblzcheck_zip_url, "-o", "package.zip")
	err = ktoblzcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ktoblzcheck_bin_url := "https://downloads.sourceforge.net/project/ktoblzcheck/ktoblzcheck-1.57.bin"
	ktoblzcheck_cmd_bin := exec.Command("curl", "-L", ktoblzcheck_bin_url, "-o", "binary.bin")
	err = ktoblzcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ktoblzcheck_src_url := "https://downloads.sourceforge.net/project/ktoblzcheck/ktoblzcheck-1.57.src.tar.gz"
	ktoblzcheck_cmd_src := exec.Command("curl", "-L", ktoblzcheck_src_url, "-o", "source.tar.gz")
	err = ktoblzcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ktoblzcheck_cmd_direct := exec.Command("./binary")
	err = ktoblzcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
