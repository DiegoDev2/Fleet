package main

// Onionprobe - Test and monitoring tool for Tor Onion Services
// Homepage: https://tpo.pages.torproject.net/onion-services/onionprobe/

import (
	"fmt"
	
	"os/exec"
)

func installOnionprobe() {
	// Método 1: Descargar y extraer .tar.gz
	onionprobe_tar_url := "https://files.pythonhosted.org/packages/aa/a7/881b66594477795314e4a5029f098eb78cf21c843b63bed8d3c7cfcf5fe4/onionprobe-1.2.0.tar.gz"
	onionprobe_cmd_tar := exec.Command("curl", "-L", onionprobe_tar_url, "-o", "package.tar.gz")
	err := onionprobe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onionprobe_zip_url := "https://files.pythonhosted.org/packages/aa/a7/881b66594477795314e4a5029f098eb78cf21c843b63bed8d3c7cfcf5fe4/onionprobe-1.2.0.zip"
	onionprobe_cmd_zip := exec.Command("curl", "-L", onionprobe_zip_url, "-o", "package.zip")
	err = onionprobe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onionprobe_bin_url := "https://files.pythonhosted.org/packages/aa/a7/881b66594477795314e4a5029f098eb78cf21c843b63bed8d3c7cfcf5fe4/onionprobe-1.2.0.bin"
	onionprobe_cmd_bin := exec.Command("curl", "-L", onionprobe_bin_url, "-o", "binary.bin")
	err = onionprobe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onionprobe_src_url := "https://files.pythonhosted.org/packages/aa/a7/881b66594477795314e4a5029f098eb78cf21c843b63bed8d3c7cfcf5fe4/onionprobe-1.2.0.src.tar.gz"
	onionprobe_cmd_src := exec.Command("curl", "-L", onionprobe_src_url, "-o", "source.tar.gz")
	err = onionprobe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onionprobe_cmd_direct := exec.Command("./binary")
	err = onionprobe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: tor")
	exec.Command("latte", "install", "tor").Run()
}
