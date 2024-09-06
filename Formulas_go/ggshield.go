package main

// Ggshield - Scanner for secrets and sensitive data in code
// Homepage: https://www.gitguardian.com

import (
	"fmt"
	
	"os/exec"
)

func installGgshield() {
	// Método 1: Descargar y extraer .tar.gz
	ggshield_tar_url := "https://files.pythonhosted.org/packages/df/1f/4fef6963ac9846a88a033a633e369a04d06ad3b77bfc7939afade16aadb9/ggshield-1.31.0.tar.gz"
	ggshield_cmd_tar := exec.Command("curl", "-L", ggshield_tar_url, "-o", "package.tar.gz")
	err := ggshield_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ggshield_zip_url := "https://files.pythonhosted.org/packages/df/1f/4fef6963ac9846a88a033a633e369a04d06ad3b77bfc7939afade16aadb9/ggshield-1.31.0.zip"
	ggshield_cmd_zip := exec.Command("curl", "-L", ggshield_zip_url, "-o", "package.zip")
	err = ggshield_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ggshield_bin_url := "https://files.pythonhosted.org/packages/df/1f/4fef6963ac9846a88a033a633e369a04d06ad3b77bfc7939afade16aadb9/ggshield-1.31.0.bin"
	ggshield_cmd_bin := exec.Command("curl", "-L", ggshield_bin_url, "-o", "binary.bin")
	err = ggshield_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ggshield_src_url := "https://files.pythonhosted.org/packages/df/1f/4fef6963ac9846a88a033a633e369a04d06ad3b77bfc7939afade16aadb9/ggshield-1.31.0.src.tar.gz"
	ggshield_cmd_src := exec.Command("curl", "-L", ggshield_src_url, "-o", "source.tar.gz")
	err = ggshield_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ggshield_cmd_direct := exec.Command("./binary")
	err = ggshield_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
