package main

// B2Tools - B2 Cloud Storage Command-Line Tools
// Homepage: https://github.com/Backblaze/B2_Command_Line_Tool

import (
	"fmt"
	
	"os/exec"
)

func installB2Tools() {
	// Método 1: Descargar y extraer .tar.gz
	b2tools_tar_url := "https://files.pythonhosted.org/packages/4c/43/4b5704ff889734bca6b8597a6acd324cc0caacc36eb26583f13444c0c61c/b2-4.1.0.tar.gz"
	b2tools_cmd_tar := exec.Command("curl", "-L", b2tools_tar_url, "-o", "package.tar.gz")
	err := b2tools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	b2tools_zip_url := "https://files.pythonhosted.org/packages/4c/43/4b5704ff889734bca6b8597a6acd324cc0caacc36eb26583f13444c0c61c/b2-4.1.0.zip"
	b2tools_cmd_zip := exec.Command("curl", "-L", b2tools_zip_url, "-o", "package.zip")
	err = b2tools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	b2tools_bin_url := "https://files.pythonhosted.org/packages/4c/43/4b5704ff889734bca6b8597a6acd324cc0caacc36eb26583f13444c0c61c/b2-4.1.0.bin"
	b2tools_cmd_bin := exec.Command("curl", "-L", b2tools_bin_url, "-o", "binary.bin")
	err = b2tools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	b2tools_src_url := "https://files.pythonhosted.org/packages/4c/43/4b5704ff889734bca6b8597a6acd324cc0caacc36eb26583f13444c0c61c/b2-4.1.0.src.tar.gz"
	b2tools_cmd_src := exec.Command("curl", "-L", b2tools_src_url, "-o", "source.tar.gz")
	err = b2tools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	b2tools_cmd_direct := exec.Command("./binary")
	err = b2tools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
