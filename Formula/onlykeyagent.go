package main

// OnlykeyAgent - Middleware that lets you use OnlyKey as a hardware SSH/GPG device
// Homepage: https://docs.crp.to/onlykey-agent.html

import (
	"fmt"
	
	"os/exec"
)

func installOnlykeyAgent() {
	// Método 1: Descargar y extraer .tar.gz
	onlykeyagent_tar_url := "https://files.pythonhosted.org/packages/68/80/e89b6c3680bedb1e14e99f0539ac805bddc7d8dd87c58805c64484966b7c/onlykey-agent-1.1.15.tar.gz"
	onlykeyagent_cmd_tar := exec.Command("curl", "-L", onlykeyagent_tar_url, "-o", "package.tar.gz")
	err := onlykeyagent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onlykeyagent_zip_url := "https://files.pythonhosted.org/packages/68/80/e89b6c3680bedb1e14e99f0539ac805bddc7d8dd87c58805c64484966b7c/onlykey-agent-1.1.15.zip"
	onlykeyagent_cmd_zip := exec.Command("curl", "-L", onlykeyagent_zip_url, "-o", "package.zip")
	err = onlykeyagent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onlykeyagent_bin_url := "https://files.pythonhosted.org/packages/68/80/e89b6c3680bedb1e14e99f0539ac805bddc7d8dd87c58805c64484966b7c/onlykey-agent-1.1.15.bin"
	onlykeyagent_cmd_bin := exec.Command("curl", "-L", onlykeyagent_bin_url, "-o", "binary.bin")
	err = onlykeyagent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onlykeyagent_src_url := "https://files.pythonhosted.org/packages/68/80/e89b6c3680bedb1e14e99f0539ac805bddc7d8dd87c58805c64484966b7c/onlykey-agent-1.1.15.src.tar.gz"
	onlykeyagent_cmd_src := exec.Command("curl", "-L", onlykeyagent_src_url, "-o", "source.tar.gz")
	err = onlykeyagent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onlykeyagent_cmd_direct := exec.Command("./binary")
	err = onlykeyagent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: cython")
	exec.Command("latte", "install", "cython").Run()
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: hidapi")
	exec.Command("latte", "install", "hidapi").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
