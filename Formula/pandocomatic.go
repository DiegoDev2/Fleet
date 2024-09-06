package main

// Pandocomatic - Automate the use of pandoc
// Homepage: https://heerdebeer.org/Software/markdown/pandocomatic/

import (
	"fmt"
	
	"os/exec"
)

func installPandocomatic() {
	// Método 1: Descargar y extraer .tar.gz
	pandocomatic_tar_url := "https://github.com/htdebeer/pandocomatic/archive/refs/tags/1.1.3.tar.gz"
	pandocomatic_cmd_tar := exec.Command("curl", "-L", pandocomatic_tar_url, "-o", "package.tar.gz")
	err := pandocomatic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pandocomatic_zip_url := "https://github.com/htdebeer/pandocomatic/archive/refs/tags/1.1.3.zip"
	pandocomatic_cmd_zip := exec.Command("curl", "-L", pandocomatic_zip_url, "-o", "package.zip")
	err = pandocomatic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pandocomatic_bin_url := "https://github.com/htdebeer/pandocomatic/archive/refs/tags/1.1.3.bin"
	pandocomatic_cmd_bin := exec.Command("curl", "-L", pandocomatic_bin_url, "-o", "binary.bin")
	err = pandocomatic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pandocomatic_src_url := "https://github.com/htdebeer/pandocomatic/archive/refs/tags/1.1.3.src.tar.gz"
	pandocomatic_cmd_src := exec.Command("curl", "-L", pandocomatic_src_url, "-o", "source.tar.gz")
	err = pandocomatic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pandocomatic_cmd_direct := exec.Command("./binary")
	err = pandocomatic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
