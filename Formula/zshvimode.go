package main

// ZshViMode - Better and friendly vi(vim) mode plugin for ZSH
// Homepage: https://github.com/jeffreytse/zsh-vi-mode

import (
	"fmt"
	
	"os/exec"
)

func installZshViMode() {
	// Método 1: Descargar y extraer .tar.gz
	zshvimode_tar_url := "https://github.com/jeffreytse/zsh-vi-mode/archive/refs/tags/v0.11.0.tar.gz"
	zshvimode_cmd_tar := exec.Command("curl", "-L", zshvimode_tar_url, "-o", "package.tar.gz")
	err := zshvimode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshvimode_zip_url := "https://github.com/jeffreytse/zsh-vi-mode/archive/refs/tags/v0.11.0.zip"
	zshvimode_cmd_zip := exec.Command("curl", "-L", zshvimode_zip_url, "-o", "package.zip")
	err = zshvimode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshvimode_bin_url := "https://github.com/jeffreytse/zsh-vi-mode/archive/refs/tags/v0.11.0.bin"
	zshvimode_cmd_bin := exec.Command("curl", "-L", zshvimode_bin_url, "-o", "binary.bin")
	err = zshvimode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshvimode_src_url := "https://github.com/jeffreytse/zsh-vi-mode/archive/refs/tags/v0.11.0.src.tar.gz"
	zshvimode_cmd_src := exec.Command("curl", "-L", zshvimode_src_url, "-o", "source.tar.gz")
	err = zshvimode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshvimode_cmd_direct := exec.Command("./binary")
	err = zshvimode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
