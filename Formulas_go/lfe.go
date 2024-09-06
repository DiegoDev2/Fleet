package main

// Lfe - Concurrent Lisp for the Erlang VM
// Homepage: https://lfe.io/

import (
	"fmt"
	
	"os/exec"
)

func installLfe() {
	// Método 1: Descargar y extraer .tar.gz
	lfe_tar_url := "https://github.com/lfe/lfe/archive/refs/tags/v2.1.4.tar.gz"
	lfe_cmd_tar := exec.Command("curl", "-L", lfe_tar_url, "-o", "package.tar.gz")
	err := lfe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lfe_zip_url := "https://github.com/lfe/lfe/archive/refs/tags/v2.1.4.zip"
	lfe_cmd_zip := exec.Command("curl", "-L", lfe_zip_url, "-o", "package.zip")
	err = lfe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lfe_bin_url := "https://github.com/lfe/lfe/archive/refs/tags/v2.1.4.bin"
	lfe_cmd_bin := exec.Command("curl", "-L", lfe_bin_url, "-o", "binary.bin")
	err = lfe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lfe_src_url := "https://github.com/lfe/lfe/archive/refs/tags/v2.1.4.src.tar.gz"
	lfe_cmd_src := exec.Command("curl", "-L", lfe_src_url, "-o", "source.tar.gz")
	err = lfe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lfe_cmd_direct := exec.Command("./binary")
	err = lfe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: emacs")
exec.Command("latte", "install", "emacs")
	fmt.Println("Instalando dependencia: erlang")
exec.Command("latte", "install", "erlang")
}
