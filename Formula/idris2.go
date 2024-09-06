package main

// Idris2 - Pure functional programming language with dependent types
// Homepage: https://www.idris-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installIdris2() {
	// Método 1: Descargar y extraer .tar.gz
	idris2_tar_url := "https://github.com/idris-lang/Idris2/archive/refs/tags/v0.7.0.tar.gz"
	idris2_cmd_tar := exec.Command("curl", "-L", idris2_tar_url, "-o", "package.tar.gz")
	err := idris2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	idris2_zip_url := "https://github.com/idris-lang/Idris2/archive/refs/tags/v0.7.0.zip"
	idris2_cmd_zip := exec.Command("curl", "-L", idris2_zip_url, "-o", "package.zip")
	err = idris2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	idris2_bin_url := "https://github.com/idris-lang/Idris2/archive/refs/tags/v0.7.0.bin"
	idris2_cmd_bin := exec.Command("curl", "-L", idris2_bin_url, "-o", "binary.bin")
	err = idris2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	idris2_src_url := "https://github.com/idris-lang/Idris2/archive/refs/tags/v0.7.0.src.tar.gz"
	idris2_cmd_src := exec.Command("curl", "-L", idris2_src_url, "-o", "source.tar.gz")
	err = idris2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	idris2_cmd_direct := exec.Command("./binary")
	err = idris2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: chezscheme")
	exec.Command("latte", "install", "chezscheme").Run()
	fmt.Println("Instalando dependencia: zsh")
	exec.Command("latte", "install", "zsh").Run()
}
