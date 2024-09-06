package main

// Unison - File synchronization tool
// Homepage: https://www.cis.upenn.edu/~bcpierce/unison/

import (
	"fmt"
	
	"os/exec"
)

func installUnison() {
	// Método 1: Descargar y extraer .tar.gz
	unison_tar_url := "https://github.com/bcpierce00/unison/archive/refs/tags/v2.53.5.tar.gz"
	unison_cmd_tar := exec.Command("curl", "-L", unison_tar_url, "-o", "package.tar.gz")
	err := unison_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unison_zip_url := "https://github.com/bcpierce00/unison/archive/refs/tags/v2.53.5.zip"
	unison_cmd_zip := exec.Command("curl", "-L", unison_zip_url, "-o", "package.zip")
	err = unison_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unison_bin_url := "https://github.com/bcpierce00/unison/archive/refs/tags/v2.53.5.bin"
	unison_cmd_bin := exec.Command("curl", "-L", unison_bin_url, "-o", "binary.bin")
	err = unison_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unison_src_url := "https://github.com/bcpierce00/unison/archive/refs/tags/v2.53.5.src.tar.gz"
	unison_cmd_src := exec.Command("curl", "-L", unison_src_url, "-o", "source.tar.gz")
	err = unison_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unison_cmd_direct := exec.Command("./binary")
	err = unison_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
}
