package main

// Cryptography - Cryptographic recipes and primitives for Python
// Homepage: https://cryptography.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installCryptography() {
	// Método 1: Descargar y extraer .tar.gz
	cryptography_tar_url := "https://files.pythonhosted.org/packages/de/ba/0664727028b37e249e73879348cc46d45c5c1a2a2e81e8166462953c5755/cryptography-43.0.1.tar.gz"
	cryptography_cmd_tar := exec.Command("curl", "-L", cryptography_tar_url, "-o", "package.tar.gz")
	err := cryptography_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cryptography_zip_url := "https://files.pythonhosted.org/packages/de/ba/0664727028b37e249e73879348cc46d45c5c1a2a2e81e8166462953c5755/cryptography-43.0.1.zip"
	cryptography_cmd_zip := exec.Command("curl", "-L", cryptography_zip_url, "-o", "package.zip")
	err = cryptography_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cryptography_bin_url := "https://files.pythonhosted.org/packages/de/ba/0664727028b37e249e73879348cc46d45c5c1a2a2e81e8166462953c5755/cryptography-43.0.1.bin"
	cryptography_cmd_bin := exec.Command("curl", "-L", cryptography_bin_url, "-o", "binary.bin")
	err = cryptography_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cryptography_src_url := "https://files.pythonhosted.org/packages/de/ba/0664727028b37e249e73879348cc46d45c5c1a2a2e81e8166462953c5755/cryptography-43.0.1.src.tar.gz"
	cryptography_cmd_src := exec.Command("curl", "-L", cryptography_src_url, "-o", "source.tar.gz")
	err = cryptography_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cryptography_cmd_direct := exec.Command("./binary")
	err = cryptography_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: cffi")
exec.Command("latte", "install", "cffi")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
