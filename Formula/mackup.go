package main

// Mackup - Keep your Mac's application settings in sync
// Homepage: https://github.com/lra/mackup

import (
	"fmt"
	
	"os/exec"
)

func installMackup() {
	// Método 1: Descargar y extraer .tar.gz
	mackup_tar_url := "https://files.pythonhosted.org/packages/63/37/8f5ee72905948757f284e7a4fea1cd8b7203f13e57d2cf4917f2f1afa7a8/mackup-0.8.41.tar.gz"
	mackup_cmd_tar := exec.Command("curl", "-L", mackup_tar_url, "-o", "package.tar.gz")
	err := mackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mackup_zip_url := "https://files.pythonhosted.org/packages/63/37/8f5ee72905948757f284e7a4fea1cd8b7203f13e57d2cf4917f2f1afa7a8/mackup-0.8.41.zip"
	mackup_cmd_zip := exec.Command("curl", "-L", mackup_zip_url, "-o", "package.zip")
	err = mackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mackup_bin_url := "https://files.pythonhosted.org/packages/63/37/8f5ee72905948757f284e7a4fea1cd8b7203f13e57d2cf4917f2f1afa7a8/mackup-0.8.41.bin"
	mackup_cmd_bin := exec.Command("curl", "-L", mackup_bin_url, "-o", "binary.bin")
	err = mackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mackup_src_url := "https://files.pythonhosted.org/packages/63/37/8f5ee72905948757f284e7a4fea1cd8b7203f13e57d2cf4917f2f1afa7a8/mackup-0.8.41.src.tar.gz"
	mackup_cmd_src := exec.Command("curl", "-L", mackup_src_url, "-o", "source.tar.gz")
	err = mackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mackup_cmd_direct := exec.Command("./binary")
	err = mackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
