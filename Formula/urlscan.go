package main

// Urlscan - View/select the URLs in an email message or file
// Homepage: https://github.com/firecat53/urlscan

import (
	"fmt"
	
	"os/exec"
)

func installUrlscan() {
	// Método 1: Descargar y extraer .tar.gz
	urlscan_tar_url := "https://files.pythonhosted.org/packages/d2/1b/83a6cfd26a4037d7271713f8aa51750fdfc5c850c5ebc93161073fd03b6c/urlscan-1.0.3.tar.gz"
	urlscan_cmd_tar := exec.Command("curl", "-L", urlscan_tar_url, "-o", "package.tar.gz")
	err := urlscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	urlscan_zip_url := "https://files.pythonhosted.org/packages/d2/1b/83a6cfd26a4037d7271713f8aa51750fdfc5c850c5ebc93161073fd03b6c/urlscan-1.0.3.zip"
	urlscan_cmd_zip := exec.Command("curl", "-L", urlscan_zip_url, "-o", "package.zip")
	err = urlscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	urlscan_bin_url := "https://files.pythonhosted.org/packages/d2/1b/83a6cfd26a4037d7271713f8aa51750fdfc5c850c5ebc93161073fd03b6c/urlscan-1.0.3.bin"
	urlscan_cmd_bin := exec.Command("curl", "-L", urlscan_bin_url, "-o", "binary.bin")
	err = urlscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	urlscan_src_url := "https://files.pythonhosted.org/packages/d2/1b/83a6cfd26a4037d7271713f8aa51750fdfc5c850c5ebc93161073fd03b6c/urlscan-1.0.3.src.tar.gz"
	urlscan_cmd_src := exec.Command("curl", "-L", urlscan_src_url, "-o", "source.tar.gz")
	err = urlscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	urlscan_cmd_direct := exec.Command("./binary")
	err = urlscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
