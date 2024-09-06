package main

// ZshHistorySubstringSearch - Zsh port of Fish shell's history search
// Homepage: https://github.com/zsh-users/zsh-history-substring-search

import (
	"fmt"
	
	"os/exec"
)

func installZshHistorySubstringSearch() {
	// Método 1: Descargar y extraer .tar.gz
	zshhistorysubstringsearch_tar_url := "https://github.com/zsh-users/zsh-history-substring-search/archive/refs/tags/v1.1.0.tar.gz"
	zshhistorysubstringsearch_cmd_tar := exec.Command("curl", "-L", zshhistorysubstringsearch_tar_url, "-o", "package.tar.gz")
	err := zshhistorysubstringsearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshhistorysubstringsearch_zip_url := "https://github.com/zsh-users/zsh-history-substring-search/archive/refs/tags/v1.1.0.zip"
	zshhistorysubstringsearch_cmd_zip := exec.Command("curl", "-L", zshhistorysubstringsearch_zip_url, "-o", "package.zip")
	err = zshhistorysubstringsearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshhistorysubstringsearch_bin_url := "https://github.com/zsh-users/zsh-history-substring-search/archive/refs/tags/v1.1.0.bin"
	zshhistorysubstringsearch_cmd_bin := exec.Command("curl", "-L", zshhistorysubstringsearch_bin_url, "-o", "binary.bin")
	err = zshhistorysubstringsearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshhistorysubstringsearch_src_url := "https://github.com/zsh-users/zsh-history-substring-search/archive/refs/tags/v1.1.0.src.tar.gz"
	zshhistorysubstringsearch_cmd_src := exec.Command("curl", "-L", zshhistorysubstringsearch_src_url, "-o", "source.tar.gz")
	err = zshhistorysubstringsearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshhistorysubstringsearch_cmd_direct := exec.Command("./binary")
	err = zshhistorysubstringsearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
