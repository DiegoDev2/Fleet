package main

// Bbot - OSINT automation tool
// Homepage: https://github.com/blacklanternsecurity/bbot

import (
	"fmt"
	
	"os/exec"
)

func installBbot() {
	// Método 1: Descargar y extraer .tar.gz
	bbot_tar_url := "https://files.pythonhosted.org/packages/ec/c8/1def3db8a9518c7c279f56e20170577cb396c63d2dae8a5f685463ac9708/bbot-1.1.8.tar.gz"
	bbot_cmd_tar := exec.Command("curl", "-L", bbot_tar_url, "-o", "package.tar.gz")
	err := bbot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bbot_zip_url := "https://files.pythonhosted.org/packages/ec/c8/1def3db8a9518c7c279f56e20170577cb396c63d2dae8a5f685463ac9708/bbot-1.1.8.zip"
	bbot_cmd_zip := exec.Command("curl", "-L", bbot_zip_url, "-o", "package.zip")
	err = bbot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bbot_bin_url := "https://files.pythonhosted.org/packages/ec/c8/1def3db8a9518c7c279f56e20170577cb396c63d2dae8a5f685463ac9708/bbot-1.1.8.bin"
	bbot_cmd_bin := exec.Command("curl", "-L", bbot_bin_url, "-o", "binary.bin")
	err = bbot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bbot_src_url := "https://files.pythonhosted.org/packages/ec/c8/1def3db8a9518c7c279f56e20170577cb396c63d2dae8a5f685463ac9708/bbot-1.1.8.src.tar.gz"
	bbot_cmd_src := exec.Command("curl", "-L", bbot_src_url, "-o", "source.tar.gz")
	err = bbot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bbot_cmd_direct := exec.Command("./binary")
	err = bbot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
