package main

// EmacsClangCompleteAsync - Emacs plugin using libclang to complete C/C++ code
// Homepage: https://github.com/Golevka/emacs-clang-complete-async

import (
	"fmt"
	
	"os/exec"
)

func installEmacsClangCompleteAsync() {
	// Método 1: Descargar y extraer .tar.gz
	emacsclangcompleteasync_tar_url := "https://github.com/Golevka/emacs-clang-complete-async/archive/refs/tags/v0.5.tar.gz"
	emacsclangcompleteasync_cmd_tar := exec.Command("curl", "-L", emacsclangcompleteasync_tar_url, "-o", "package.tar.gz")
	err := emacsclangcompleteasync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emacsclangcompleteasync_zip_url := "https://github.com/Golevka/emacs-clang-complete-async/archive/refs/tags/v0.5.zip"
	emacsclangcompleteasync_cmd_zip := exec.Command("curl", "-L", emacsclangcompleteasync_zip_url, "-o", "package.zip")
	err = emacsclangcompleteasync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emacsclangcompleteasync_bin_url := "https://github.com/Golevka/emacs-clang-complete-async/archive/refs/tags/v0.5.bin"
	emacsclangcompleteasync_cmd_bin := exec.Command("curl", "-L", emacsclangcompleteasync_bin_url, "-o", "binary.bin")
	err = emacsclangcompleteasync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emacsclangcompleteasync_src_url := "https://github.com/Golevka/emacs-clang-complete-async/archive/refs/tags/v0.5.src.tar.gz"
	emacsclangcompleteasync_cmd_src := exec.Command("curl", "-L", emacsclangcompleteasync_src_url, "-o", "source.tar.gz")
	err = emacsclangcompleteasync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emacsclangcompleteasync_cmd_direct := exec.Command("./binary")
	err = emacsclangcompleteasync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
