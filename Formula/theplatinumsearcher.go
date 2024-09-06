package main

// ThePlatinumSearcher - Multi-platform code-search similar to ack and ag
// Homepage: https://github.com/monochromegane/the_platinum_searcher

import (
	"fmt"
	
	"os/exec"
)

func installThePlatinumSearcher() {
	// Método 1: Descargar y extraer .tar.gz
	theplatinumsearcher_tar_url := "https://github.com/monochromegane/the_platinum_searcher/archive/refs/tags/v2.2.0.tar.gz"
	theplatinumsearcher_cmd_tar := exec.Command("curl", "-L", theplatinumsearcher_tar_url, "-o", "package.tar.gz")
	err := theplatinumsearcher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	theplatinumsearcher_zip_url := "https://github.com/monochromegane/the_platinum_searcher/archive/refs/tags/v2.2.0.zip"
	theplatinumsearcher_cmd_zip := exec.Command("curl", "-L", theplatinumsearcher_zip_url, "-o", "package.zip")
	err = theplatinumsearcher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	theplatinumsearcher_bin_url := "https://github.com/monochromegane/the_platinum_searcher/archive/refs/tags/v2.2.0.bin"
	theplatinumsearcher_cmd_bin := exec.Command("curl", "-L", theplatinumsearcher_bin_url, "-o", "binary.bin")
	err = theplatinumsearcher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	theplatinumsearcher_src_url := "https://github.com/monochromegane/the_platinum_searcher/archive/refs/tags/v2.2.0.src.tar.gz"
	theplatinumsearcher_cmd_src := exec.Command("curl", "-L", theplatinumsearcher_src_url, "-o", "source.tar.gz")
	err = theplatinumsearcher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	theplatinumsearcher_cmd_direct := exec.Command("./binary")
	err = theplatinumsearcher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
