package main

// Dolphie - Feature-rich top tool for monitoring MySQL
// Homepage: https://github.com/charles-001/dolphie

import (
	"fmt"
	
	"os/exec"
)

func installDolphie() {
	// Método 1: Descargar y extraer .tar.gz
	dolphie_tar_url := "https://files.pythonhosted.org/packages/7c/ae/724873b2d46539f1ad320510c4ee66e2be5dcad4a5cc4ff3470e23d69f5d/dolphie-6.0.4.tar.gz"
	dolphie_cmd_tar := exec.Command("curl", "-L", dolphie_tar_url, "-o", "package.tar.gz")
	err := dolphie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dolphie_zip_url := "https://files.pythonhosted.org/packages/7c/ae/724873b2d46539f1ad320510c4ee66e2be5dcad4a5cc4ff3470e23d69f5d/dolphie-6.0.4.zip"
	dolphie_cmd_zip := exec.Command("curl", "-L", dolphie_zip_url, "-o", "package.zip")
	err = dolphie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dolphie_bin_url := "https://files.pythonhosted.org/packages/7c/ae/724873b2d46539f1ad320510c4ee66e2be5dcad4a5cc4ff3470e23d69f5d/dolphie-6.0.4.bin"
	dolphie_cmd_bin := exec.Command("curl", "-L", dolphie_bin_url, "-o", "binary.bin")
	err = dolphie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dolphie_src_url := "https://files.pythonhosted.org/packages/7c/ae/724873b2d46539f1ad320510c4ee66e2be5dcad4a5cc4ff3470e23d69f5d/dolphie-6.0.4.src.tar.gz"
	dolphie_cmd_src := exec.Command("curl", "-L", dolphie_src_url, "-o", "source.tar.gz")
	err = dolphie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dolphie_cmd_direct := exec.Command("./binary")
	err = dolphie_cmd_direct.Run()
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
