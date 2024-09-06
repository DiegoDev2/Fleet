package main

// Keyring - Easy way to access the system keyring service from python
// Homepage: https://github.com/jaraco/keyring

import (
	"fmt"
	
	"os/exec"
)

func installKeyring() {
	// Método 1: Descargar y extraer .tar.gz
	keyring_tar_url := "https://files.pythonhosted.org/packages/32/30/bfdde7294ba6bb2f519950687471dc6a0996d4f77ab30d75c841fa4994ed/keyring-25.3.0.tar.gz"
	keyring_cmd_tar := exec.Command("curl", "-L", keyring_tar_url, "-o", "package.tar.gz")
	err := keyring_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keyring_zip_url := "https://files.pythonhosted.org/packages/32/30/bfdde7294ba6bb2f519950687471dc6a0996d4f77ab30d75c841fa4994ed/keyring-25.3.0.zip"
	keyring_cmd_zip := exec.Command("curl", "-L", keyring_zip_url, "-o", "package.zip")
	err = keyring_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keyring_bin_url := "https://files.pythonhosted.org/packages/32/30/bfdde7294ba6bb2f519950687471dc6a0996d4f77ab30d75c841fa4994ed/keyring-25.3.0.bin"
	keyring_cmd_bin := exec.Command("curl", "-L", keyring_bin_url, "-o", "binary.bin")
	err = keyring_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keyring_src_url := "https://files.pythonhosted.org/packages/32/30/bfdde7294ba6bb2f519950687471dc6a0996d4f77ab30d75c841fa4994ed/keyring-25.3.0.src.tar.gz"
	keyring_cmd_src := exec.Command("curl", "-L", keyring_src_url, "-o", "source.tar.gz")
	err = keyring_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keyring_cmd_direct := exec.Command("./binary")
	err = keyring_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
}
