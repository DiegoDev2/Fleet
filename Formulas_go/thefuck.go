package main

// Thefuck - Programmatically correct mistyped console commands
// Homepage: https://github.com/nvbn/thefuck

import (
	"fmt"
	
	"os/exec"
)

func installThefuck() {
	// Método 1: Descargar y extraer .tar.gz
	thefuck_tar_url := "https://files.pythonhosted.org/packages/ac/d0/0c256afd3ba1d05882154d16aa0685018f21c60a6769a496558da7d9d8f1/thefuck-3.32.tar.gz"
	thefuck_cmd_tar := exec.Command("curl", "-L", thefuck_tar_url, "-o", "package.tar.gz")
	err := thefuck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	thefuck_zip_url := "https://files.pythonhosted.org/packages/ac/d0/0c256afd3ba1d05882154d16aa0685018f21c60a6769a496558da7d9d8f1/thefuck-3.32.zip"
	thefuck_cmd_zip := exec.Command("curl", "-L", thefuck_zip_url, "-o", "package.zip")
	err = thefuck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	thefuck_bin_url := "https://files.pythonhosted.org/packages/ac/d0/0c256afd3ba1d05882154d16aa0685018f21c60a6769a496558da7d9d8f1/thefuck-3.32.bin"
	thefuck_cmd_bin := exec.Command("curl", "-L", thefuck_bin_url, "-o", "binary.bin")
	err = thefuck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	thefuck_src_url := "https://files.pythonhosted.org/packages/ac/d0/0c256afd3ba1d05882154d16aa0685018f21c60a6769a496558da7d9d8f1/thefuck-3.32.src.tar.gz"
	thefuck_cmd_src := exec.Command("curl", "-L", thefuck_src_url, "-o", "source.tar.gz")
	err = thefuck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	thefuck_cmd_direct := exec.Command("./binary")
	err = thefuck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
