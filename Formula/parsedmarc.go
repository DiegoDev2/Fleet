package main

// Parsedmarc - DMARC report analyzer and visualizer
// Homepage: https://domainaware.github.io/parsedmarc/

import (
	"fmt"
	
	"os/exec"
)

func installParsedmarc() {
	// Método 1: Descargar y extraer .tar.gz
	parsedmarc_tar_url := "https://files.pythonhosted.org/packages/29/95/24f5d01195312127ee1c31799987d041c74fdce49f512122b216d465ed61/parsedmarc-8.15.0.tar.gz"
	parsedmarc_cmd_tar := exec.Command("curl", "-L", parsedmarc_tar_url, "-o", "package.tar.gz")
	err := parsedmarc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parsedmarc_zip_url := "https://files.pythonhosted.org/packages/29/95/24f5d01195312127ee1c31799987d041c74fdce49f512122b216d465ed61/parsedmarc-8.15.0.zip"
	parsedmarc_cmd_zip := exec.Command("curl", "-L", parsedmarc_zip_url, "-o", "package.zip")
	err = parsedmarc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parsedmarc_bin_url := "https://files.pythonhosted.org/packages/29/95/24f5d01195312127ee1c31799987d041c74fdce49f512122b216d465ed61/parsedmarc-8.15.0.bin"
	parsedmarc_cmd_bin := exec.Command("curl", "-L", parsedmarc_bin_url, "-o", "binary.bin")
	err = parsedmarc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parsedmarc_src_url := "https://files.pythonhosted.org/packages/29/95/24f5d01195312127ee1c31799987d041c74fdce49f512122b216d465ed61/parsedmarc-8.15.0.src.tar.gz"
	parsedmarc_cmd_src := exec.Command("curl", "-L", parsedmarc_src_url, "-o", "source.tar.gz")
	err = parsedmarc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parsedmarc_cmd_direct := exec.Command("./binary")
	err = parsedmarc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
