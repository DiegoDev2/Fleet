package main

// Wrangler - Refactoring tool for Erlang with emacs and Eclipse integration
// Homepage: https://refactoringtools.github.io/docs/wrangler/

import (
	"fmt"
	
	"os/exec"
)

func installWrangler() {
	// Método 1: Descargar y extraer .tar.gz
	wrangler_tar_url := "https://github.com/RefactoringTools/wrangler/archive/refs/tags/wrangler1.2.tar.gz"
	wrangler_cmd_tar := exec.Command("curl", "-L", wrangler_tar_url, "-o", "package.tar.gz")
	err := wrangler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wrangler_zip_url := "https://github.com/RefactoringTools/wrangler/archive/refs/tags/wrangler1.2.zip"
	wrangler_cmd_zip := exec.Command("curl", "-L", wrangler_zip_url, "-o", "package.zip")
	err = wrangler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wrangler_bin_url := "https://github.com/RefactoringTools/wrangler/archive/refs/tags/wrangler1.2.bin"
	wrangler_cmd_bin := exec.Command("curl", "-L", wrangler_bin_url, "-o", "binary.bin")
	err = wrangler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wrangler_src_url := "https://github.com/RefactoringTools/wrangler/archive/refs/tags/wrangler1.2.src.tar.gz"
	wrangler_cmd_src := exec.Command("curl", "-L", wrangler_src_url, "-o", "source.tar.gz")
	err = wrangler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wrangler_cmd_direct := exec.Command("./binary")
	err = wrangler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: erlang")
exec.Command("latte", "install", "erlang")
}
