package main

// Sigstore - Codesigning tool for Python packages
// Homepage: https://github.com/sigstore/sigstore-python

import (
	"fmt"
	
	"os/exec"
)

func installSigstore() {
	// Método 1: Descargar y extraer .tar.gz
	sigstore_tar_url := "https://files.pythonhosted.org/packages/f3/63/daee649bc4e97048474fbc184300dfcf511dadfa422875b42226efd3a9fc/sigstore-3.2.0.tar.gz"
	sigstore_cmd_tar := exec.Command("curl", "-L", sigstore_tar_url, "-o", "package.tar.gz")
	err := sigstore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sigstore_zip_url := "https://files.pythonhosted.org/packages/f3/63/daee649bc4e97048474fbc184300dfcf511dadfa422875b42226efd3a9fc/sigstore-3.2.0.zip"
	sigstore_cmd_zip := exec.Command("curl", "-L", sigstore_zip_url, "-o", "package.zip")
	err = sigstore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sigstore_bin_url := "https://files.pythonhosted.org/packages/f3/63/daee649bc4e97048474fbc184300dfcf511dadfa422875b42226efd3a9fc/sigstore-3.2.0.bin"
	sigstore_cmd_bin := exec.Command("curl", "-L", sigstore_bin_url, "-o", "binary.bin")
	err = sigstore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sigstore_src_url := "https://files.pythonhosted.org/packages/f3/63/daee649bc4e97048474fbc184300dfcf511dadfa422875b42226efd3a9fc/sigstore-3.2.0.src.tar.gz"
	sigstore_cmd_src := exec.Command("curl", "-L", sigstore_src_url, "-o", "source.tar.gz")
	err = sigstore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sigstore_cmd_direct := exec.Command("./binary")
	err = sigstore_cmd_direct.Run()
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
