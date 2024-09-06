package main

// Khal - CLI calendar application
// Homepage: https://lostpackets.de/khal/

import (
	"fmt"
	
	"os/exec"
)

func installKhal() {
	// Método 1: Descargar y extraer .tar.gz
	khal_tar_url := "https://files.pythonhosted.org/packages/d3/58/665551b1fea58a70d0f70fb539d2cd6be9ec106f36023d62c3ec5c7b2de1/khal-0.11.3.tar.gz"
	khal_cmd_tar := exec.Command("curl", "-L", khal_tar_url, "-o", "package.tar.gz")
	err := khal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	khal_zip_url := "https://files.pythonhosted.org/packages/d3/58/665551b1fea58a70d0f70fb539d2cd6be9ec106f36023d62c3ec5c7b2de1/khal-0.11.3.zip"
	khal_cmd_zip := exec.Command("curl", "-L", khal_zip_url, "-o", "package.zip")
	err = khal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	khal_bin_url := "https://files.pythonhosted.org/packages/d3/58/665551b1fea58a70d0f70fb539d2cd6be9ec106f36023d62c3ec5c7b2de1/khal-0.11.3.bin"
	khal_cmd_bin := exec.Command("curl", "-L", khal_bin_url, "-o", "binary.bin")
	err = khal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	khal_src_url := "https://files.pythonhosted.org/packages/d3/58/665551b1fea58a70d0f70fb539d2cd6be9ec106f36023d62c3ec5c7b2de1/khal-0.11.3.src.tar.gz"
	khal_cmd_src := exec.Command("curl", "-L", khal_src_url, "-o", "source.tar.gz")
	err = khal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	khal_cmd_direct := exec.Command("./binary")
	err = khal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
