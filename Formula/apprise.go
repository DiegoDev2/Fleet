package main

// Apprise - Send notifications from the command-line to popular notification services
// Homepage: https://pypi.org/project/apprise/

import (
	"fmt"
	
	"os/exec"
)

func installApprise() {
	// Método 1: Descargar y extraer .tar.gz
	apprise_tar_url := "https://files.pythonhosted.org/packages/92/26/19c26dbf32d31129c50a3568022ae1c9d05c4aac056c0661d9bfea0f7810/apprise-1.9.0.tar.gz"
	apprise_cmd_tar := exec.Command("curl", "-L", apprise_tar_url, "-o", "package.tar.gz")
	err := apprise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apprise_zip_url := "https://files.pythonhosted.org/packages/92/26/19c26dbf32d31129c50a3568022ae1c9d05c4aac056c0661d9bfea0f7810/apprise-1.9.0.zip"
	apprise_cmd_zip := exec.Command("curl", "-L", apprise_zip_url, "-o", "package.zip")
	err = apprise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apprise_bin_url := "https://files.pythonhosted.org/packages/92/26/19c26dbf32d31129c50a3568022ae1c9d05c4aac056c0661d9bfea0f7810/apprise-1.9.0.bin"
	apprise_cmd_bin := exec.Command("curl", "-L", apprise_bin_url, "-o", "binary.bin")
	err = apprise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apprise_src_url := "https://files.pythonhosted.org/packages/92/26/19c26dbf32d31129c50a3568022ae1c9d05c4aac056c0661d9bfea0f7810/apprise-1.9.0.src.tar.gz"
	apprise_cmd_src := exec.Command("curl", "-L", apprise_src_url, "-o", "source.tar.gz")
	err = apprise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apprise_cmd_direct := exec.Command("./binary")
	err = apprise_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
