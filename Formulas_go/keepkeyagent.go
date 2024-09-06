package main

// KeepkeyAgent - Keepkey Hardware-based SSH/GPG agent
// Homepage: https://github.com/romanz/trezor-agent

import (
	"fmt"
	
	"os/exec"
)

func installKeepkeyAgent() {
	// Método 1: Descargar y extraer .tar.gz
	keepkeyagent_tar_url := "https://files.pythonhosted.org/packages/65/72/4bf47a7bc8dc93d2ac21672a0db4bc58a78ec5cee3c4bcebd0b4092a9110/keepkey_agent-0.9.0.tar.gz"
	keepkeyagent_cmd_tar := exec.Command("curl", "-L", keepkeyagent_tar_url, "-o", "package.tar.gz")
	err := keepkeyagent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keepkeyagent_zip_url := "https://files.pythonhosted.org/packages/65/72/4bf47a7bc8dc93d2ac21672a0db4bc58a78ec5cee3c4bcebd0b4092a9110/keepkey_agent-0.9.0.zip"
	keepkeyagent_cmd_zip := exec.Command("curl", "-L", keepkeyagent_zip_url, "-o", "package.zip")
	err = keepkeyagent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keepkeyagent_bin_url := "https://files.pythonhosted.org/packages/65/72/4bf47a7bc8dc93d2ac21672a0db4bc58a78ec5cee3c4bcebd0b4092a9110/keepkey_agent-0.9.0.bin"
	keepkeyagent_cmd_bin := exec.Command("curl", "-L", keepkeyagent_bin_url, "-o", "binary.bin")
	err = keepkeyagent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keepkeyagent_src_url := "https://files.pythonhosted.org/packages/65/72/4bf47a7bc8dc93d2ac21672a0db4bc58a78ec5cee3c4bcebd0b4092a9110/keepkey_agent-0.9.0.src.tar.gz"
	keepkeyagent_cmd_src := exec.Command("curl", "-L", keepkeyagent_src_url, "-o", "source.tar.gz")
	err = keepkeyagent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keepkeyagent_cmd_direct := exec.Command("./binary")
	err = keepkeyagent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
