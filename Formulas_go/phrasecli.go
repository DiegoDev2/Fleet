package main

// PhraseCli - Tool to interact with the Phrase API
// Homepage: https://phrase.com/

import (
	"fmt"
	
	"os/exec"
)

func installPhraseCli() {
	// Método 1: Descargar y extraer .tar.gz
	phrasecli_tar_url := "https://github.com/phrase/phrase-cli/archive/refs/tags/2.31.2.tar.gz"
	phrasecli_cmd_tar := exec.Command("curl", "-L", phrasecli_tar_url, "-o", "package.tar.gz")
	err := phrasecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phrasecli_zip_url := "https://github.com/phrase/phrase-cli/archive/refs/tags/2.31.2.zip"
	phrasecli_cmd_zip := exec.Command("curl", "-L", phrasecli_zip_url, "-o", "package.zip")
	err = phrasecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phrasecli_bin_url := "https://github.com/phrase/phrase-cli/archive/refs/tags/2.31.2.bin"
	phrasecli_cmd_bin := exec.Command("curl", "-L", phrasecli_bin_url, "-o", "binary.bin")
	err = phrasecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phrasecli_src_url := "https://github.com/phrase/phrase-cli/archive/refs/tags/2.31.2.src.tar.gz"
	phrasecli_cmd_src := exec.Command("curl", "-L", phrasecli_src_url, "-o", "source.tar.gz")
	err = phrasecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phrasecli_cmd_direct := exec.Command("./binary")
	err = phrasecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
