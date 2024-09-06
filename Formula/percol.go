package main

// Percol - Interactive grep tool
// Homepage: https://github.com/mooz/percol

import (
	"fmt"
	
	"os/exec"
)

func installPercol() {
	// Método 1: Descargar y extraer .tar.gz
	percol_tar_url := "https://files.pythonhosted.org/packages/50/ea/282b2df42d6be8d4292206ea9169742951c39374af43ae0d6f9fff0af599/percol-0.2.1.tar.gz"
	percol_cmd_tar := exec.Command("curl", "-L", percol_tar_url, "-o", "package.tar.gz")
	err := percol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	percol_zip_url := "https://files.pythonhosted.org/packages/50/ea/282b2df42d6be8d4292206ea9169742951c39374af43ae0d6f9fff0af599/percol-0.2.1.zip"
	percol_cmd_zip := exec.Command("curl", "-L", percol_zip_url, "-o", "package.zip")
	err = percol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	percol_bin_url := "https://files.pythonhosted.org/packages/50/ea/282b2df42d6be8d4292206ea9169742951c39374af43ae0d6f9fff0af599/percol-0.2.1.bin"
	percol_cmd_bin := exec.Command("curl", "-L", percol_bin_url, "-o", "binary.bin")
	err = percol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	percol_src_url := "https://files.pythonhosted.org/packages/50/ea/282b2df42d6be8d4292206ea9169742951c39374af43ae0d6f9fff0af599/percol-0.2.1.src.tar.gz"
	percol_cmd_src := exec.Command("curl", "-L", percol_src_url, "-o", "source.tar.gz")
	err = percol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	percol_cmd_direct := exec.Command("./binary")
	err = percol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
