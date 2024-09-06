package main

// ChromeExport - Convert Chrome's bookmarks and history to HTML bookmarks files
// Homepage: https://github.com/bdesham/chrome-export

import (
	"fmt"
	
	"os/exec"
)

func installChromeExport() {
	// Método 1: Descargar y extraer .tar.gz
	chromeexport_tar_url := "https://github.com/bdesham/chrome-export/archive/refs/tags/v2.0.2.tar.gz"
	chromeexport_cmd_tar := exec.Command("curl", "-L", chromeexport_tar_url, "-o", "package.tar.gz")
	err := chromeexport_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chromeexport_zip_url := "https://github.com/bdesham/chrome-export/archive/refs/tags/v2.0.2.zip"
	chromeexport_cmd_zip := exec.Command("curl", "-L", chromeexport_zip_url, "-o", "package.zip")
	err = chromeexport_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chromeexport_bin_url := "https://github.com/bdesham/chrome-export/archive/refs/tags/v2.0.2.bin"
	chromeexport_cmd_bin := exec.Command("curl", "-L", chromeexport_bin_url, "-o", "binary.bin")
	err = chromeexport_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chromeexport_src_url := "https://github.com/bdesham/chrome-export/archive/refs/tags/v2.0.2.src.tar.gz"
	chromeexport_cmd_src := exec.Command("curl", "-L", chromeexport_src_url, "-o", "source.tar.gz")
	err = chromeexport_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chromeexport_cmd_direct := exec.Command("./binary")
	err = chromeexport_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
