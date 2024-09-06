package main

// ZshAutocomplete - Real-time type-ahead completion for Zsh
// Homepage: https://github.com/marlonrichert/zsh-autocomplete

import (
	"fmt"
	
	"os/exec"
)

func installZshAutocomplete() {
	// Método 1: Descargar y extraer .tar.gz
	zshautocomplete_tar_url := "https://github.com/marlonrichert/zsh-autocomplete/archive/refs/tags/24.09.04.tar.gz"
	zshautocomplete_cmd_tar := exec.Command("curl", "-L", zshautocomplete_tar_url, "-o", "package.tar.gz")
	err := zshautocomplete_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshautocomplete_zip_url := "https://github.com/marlonrichert/zsh-autocomplete/archive/refs/tags/24.09.04.zip"
	zshautocomplete_cmd_zip := exec.Command("curl", "-L", zshautocomplete_zip_url, "-o", "package.zip")
	err = zshautocomplete_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshautocomplete_bin_url := "https://github.com/marlonrichert/zsh-autocomplete/archive/refs/tags/24.09.04.bin"
	zshautocomplete_cmd_bin := exec.Command("curl", "-L", zshautocomplete_bin_url, "-o", "binary.bin")
	err = zshautocomplete_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshautocomplete_src_url := "https://github.com/marlonrichert/zsh-autocomplete/archive/refs/tags/24.09.04.src.tar.gz"
	zshautocomplete_cmd_src := exec.Command("curl", "-L", zshautocomplete_src_url, "-o", "source.tar.gz")
	err = zshautocomplete_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshautocomplete_cmd_direct := exec.Command("./binary")
	err = zshautocomplete_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: clitest")
	exec.Command("latte", "install", "clitest").Run()
}
