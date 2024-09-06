package main

// Aider - AI pair programming in your terminal
// Homepage: https://aider.chat/

import (
	"fmt"
	
	"os/exec"
)

func installAider() {
	// Método 1: Descargar y extraer .tar.gz
	aider_tar_url := "https://files.pythonhosted.org/packages/7f/91/2ff612a0b6a4b51b70eb43a1cb77d54988627ed0c36d19037ebb6ed5c5ec/aider_chat-0.55.0.tar.gz"
	aider_cmd_tar := exec.Command("curl", "-L", aider_tar_url, "-o", "package.tar.gz")
	err := aider_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aider_zip_url := "https://files.pythonhosted.org/packages/7f/91/2ff612a0b6a4b51b70eb43a1cb77d54988627ed0c36d19037ebb6ed5c5ec/aider_chat-0.55.0.zip"
	aider_cmd_zip := exec.Command("curl", "-L", aider_zip_url, "-o", "package.zip")
	err = aider_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aider_bin_url := "https://files.pythonhosted.org/packages/7f/91/2ff612a0b6a4b51b70eb43a1cb77d54988627ed0c36d19037ebb6ed5c5ec/aider_chat-0.55.0.bin"
	aider_cmd_bin := exec.Command("curl", "-L", aider_bin_url, "-o", "binary.bin")
	err = aider_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aider_src_url := "https://files.pythonhosted.org/packages/7f/91/2ff612a0b6a4b51b70eb43a1cb77d54988627ed0c36d19037ebb6ed5c5ec/aider_chat-0.55.0.src.tar.gz"
	aider_cmd_src := exec.Command("curl", "-L", aider_src_url, "-o", "source.tar.gz")
	err = aider_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aider_cmd_direct := exec.Command("./binary")
	err = aider_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cython")
exec.Command("latte", "install", "cython")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cffi")
exec.Command("latte", "install", "cffi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: scipy")
exec.Command("latte", "install", "scipy")
}
