package main

// Pwntools - CTF framework used by Gallopsled in every CTF
// Homepage: https://github.com/Gallopsled/pwntools

import (
	"fmt"
	
	"os/exec"
)

func installPwntools() {
	// Método 1: Descargar y extraer .tar.gz
	pwntools_tar_url := "https://files.pythonhosted.org/packages/5e/c0/7f2a0dd8e1641a25a5066b2235e4556796fbd962ff45bfb9ce9aaec6c74e/pwntools-4.13.0.tar.gz"
	pwntools_cmd_tar := exec.Command("curl", "-L", pwntools_tar_url, "-o", "package.tar.gz")
	err := pwntools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pwntools_zip_url := "https://files.pythonhosted.org/packages/5e/c0/7f2a0dd8e1641a25a5066b2235e4556796fbd962ff45bfb9ce9aaec6c74e/pwntools-4.13.0.zip"
	pwntools_cmd_zip := exec.Command("curl", "-L", pwntools_zip_url, "-o", "package.zip")
	err = pwntools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pwntools_bin_url := "https://files.pythonhosted.org/packages/5e/c0/7f2a0dd8e1641a25a5066b2235e4556796fbd962ff45bfb9ce9aaec6c74e/pwntools-4.13.0.bin"
	pwntools_cmd_bin := exec.Command("curl", "-L", pwntools_bin_url, "-o", "binary.bin")
	err = pwntools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pwntools_src_url := "https://files.pythonhosted.org/packages/5e/c0/7f2a0dd8e1641a25a5066b2235e4556796fbd962ff45bfb9ce9aaec6c74e/pwntools-4.13.0.src.tar.gz"
	pwntools_cmd_src := exec.Command("curl", "-L", pwntools_src_url, "-o", "source.tar.gz")
	err = pwntools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pwntools_cmd_direct := exec.Command("./binary")
	err = pwntools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
