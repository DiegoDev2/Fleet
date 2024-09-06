package main

// Joker - Small Clojure interpreter, linter and formatter
// Homepage: https://joker-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installJoker() {
	// Método 1: Descargar y extraer .tar.gz
	joker_tar_url := "https://github.com/candid82/joker/archive/refs/tags/v1.4.0.tar.gz"
	joker_cmd_tar := exec.Command("curl", "-L", joker_tar_url, "-o", "package.tar.gz")
	err := joker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	joker_zip_url := "https://github.com/candid82/joker/archive/refs/tags/v1.4.0.zip"
	joker_cmd_zip := exec.Command("curl", "-L", joker_zip_url, "-o", "package.zip")
	err = joker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	joker_bin_url := "https://github.com/candid82/joker/archive/refs/tags/v1.4.0.bin"
	joker_cmd_bin := exec.Command("curl", "-L", joker_bin_url, "-o", "binary.bin")
	err = joker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	joker_src_url := "https://github.com/candid82/joker/archive/refs/tags/v1.4.0.src.tar.gz"
	joker_cmd_src := exec.Command("curl", "-L", joker_src_url, "-o", "source.tar.gz")
	err = joker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	joker_cmd_direct := exec.Command("./binary")
	err = joker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
