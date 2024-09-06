package main

// TrezorAgent - Hardware SSH/GPG agent for Trezor, Keepkey & Ledger
// Homepage: https://github.com/romanz/trezor-agent

import (
	"fmt"
	
	"os/exec"
)

func installTrezorAgent() {
	// Método 1: Descargar y extraer .tar.gz
	trezoragent_tar_url := "https://files.pythonhosted.org/packages/11/bc/aa2bdee9cd81af9ecde0a9e8b5c6c6594a4a0ee7ade950b51a39d54f9e63/trezor_agent-0.12.0.tar.gz"
	trezoragent_cmd_tar := exec.Command("curl", "-L", trezoragent_tar_url, "-o", "package.tar.gz")
	err := trezoragent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trezoragent_zip_url := "https://files.pythonhosted.org/packages/11/bc/aa2bdee9cd81af9ecde0a9e8b5c6c6594a4a0ee7ade950b51a39d54f9e63/trezor_agent-0.12.0.zip"
	trezoragent_cmd_zip := exec.Command("curl", "-L", trezoragent_zip_url, "-o", "package.zip")
	err = trezoragent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trezoragent_bin_url := "https://files.pythonhosted.org/packages/11/bc/aa2bdee9cd81af9ecde0a9e8b5c6c6594a4a0ee7ade950b51a39d54f9e63/trezor_agent-0.12.0.bin"
	trezoragent_cmd_bin := exec.Command("curl", "-L", trezoragent_bin_url, "-o", "binary.bin")
	err = trezoragent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trezoragent_src_url := "https://files.pythonhosted.org/packages/11/bc/aa2bdee9cd81af9ecde0a9e8b5c6c6594a4a0ee7ade950b51a39d54f9e63/trezor_agent-0.12.0.src.tar.gz"
	trezoragent_cmd_src := exec.Command("curl", "-L", trezoragent_src_url, "-o", "source.tar.gz")
	err = trezoragent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trezoragent_cmd_direct := exec.Command("./binary")
	err = trezoragent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
