package main

// Wiki - Fetch summaries from MediaWiki wikis, like Wikipedia
// Homepage: https://github.com/walle/wiki

import (
	"fmt"
	
	"os/exec"
)

func installWiki() {
	// Método 1: Descargar y extraer .tar.gz
	wiki_tar_url := "https://github.com/walle/wiki/archive/refs/tags/v1.4.1.tar.gz"
	wiki_cmd_tar := exec.Command("curl", "-L", wiki_tar_url, "-o", "package.tar.gz")
	err := wiki_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wiki_zip_url := "https://github.com/walle/wiki/archive/refs/tags/v1.4.1.zip"
	wiki_cmd_zip := exec.Command("curl", "-L", wiki_zip_url, "-o", "package.zip")
	err = wiki_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wiki_bin_url := "https://github.com/walle/wiki/archive/refs/tags/v1.4.1.bin"
	wiki_cmd_bin := exec.Command("curl", "-L", wiki_bin_url, "-o", "binary.bin")
	err = wiki_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wiki_src_url := "https://github.com/walle/wiki/archive/refs/tags/v1.4.1.src.tar.gz"
	wiki_cmd_src := exec.Command("curl", "-L", wiki_src_url, "-o", "source.tar.gz")
	err = wiki_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wiki_cmd_direct := exec.Command("./binary")
	err = wiki_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
