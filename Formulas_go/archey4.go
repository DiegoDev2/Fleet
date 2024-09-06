package main

// Archey4 - Simple system information tool written in Python
// Homepage: https://github.com/HorlogeSkynet/archey4

import (
	"fmt"
	
	"os/exec"
)

func installArchey4() {
	// Método 1: Descargar y extraer .tar.gz
	archey4_tar_url := "https://files.pythonhosted.org/packages/dc/9a/8a0921ed2df4da26197f98b0f6a8b3af1fa232cc898a13e719e9ed396e95/archey4-4.14.3.0.tar.gz"
	archey4_cmd_tar := exec.Command("curl", "-L", archey4_tar_url, "-o", "package.tar.gz")
	err := archey4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	archey4_zip_url := "https://files.pythonhosted.org/packages/dc/9a/8a0921ed2df4da26197f98b0f6a8b3af1fa232cc898a13e719e9ed396e95/archey4-4.14.3.0.zip"
	archey4_cmd_zip := exec.Command("curl", "-L", archey4_zip_url, "-o", "package.zip")
	err = archey4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	archey4_bin_url := "https://files.pythonhosted.org/packages/dc/9a/8a0921ed2df4da26197f98b0f6a8b3af1fa232cc898a13e719e9ed396e95/archey4-4.14.3.0.bin"
	archey4_cmd_bin := exec.Command("curl", "-L", archey4_bin_url, "-o", "binary.bin")
	err = archey4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	archey4_src_url := "https://files.pythonhosted.org/packages/dc/9a/8a0921ed2df4da26197f98b0f6a8b3af1fa232cc898a13e719e9ed396e95/archey4-4.14.3.0.src.tar.gz"
	archey4_cmd_src := exec.Command("curl", "-L", archey4_src_url, "-o", "source.tar.gz")
	err = archey4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	archey4_cmd_direct := exec.Command("./binary")
	err = archey4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
