package main

// Defaultbrowser - Command-line tool for getting & setting the default browser
// Homepage: https://github.com/kerma/defaultbrowser

import (
	"fmt"
	
	"os/exec"
)

func installDefaultbrowser() {
	// Método 1: Descargar y extraer .tar.gz
	defaultbrowser_tar_url := "https://github.com/kerma/defaultbrowser/archive/refs/tags/1.1.tar.gz"
	defaultbrowser_cmd_tar := exec.Command("curl", "-L", defaultbrowser_tar_url, "-o", "package.tar.gz")
	err := defaultbrowser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	defaultbrowser_zip_url := "https://github.com/kerma/defaultbrowser/archive/refs/tags/1.1.zip"
	defaultbrowser_cmd_zip := exec.Command("curl", "-L", defaultbrowser_zip_url, "-o", "package.zip")
	err = defaultbrowser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	defaultbrowser_bin_url := "https://github.com/kerma/defaultbrowser/archive/refs/tags/1.1.bin"
	defaultbrowser_cmd_bin := exec.Command("curl", "-L", defaultbrowser_bin_url, "-o", "binary.bin")
	err = defaultbrowser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	defaultbrowser_src_url := "https://github.com/kerma/defaultbrowser/archive/refs/tags/1.1.src.tar.gz"
	defaultbrowser_cmd_src := exec.Command("curl", "-L", defaultbrowser_src_url, "-o", "source.tar.gz")
	err = defaultbrowser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	defaultbrowser_cmd_direct := exec.Command("./binary")
	err = defaultbrowser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
