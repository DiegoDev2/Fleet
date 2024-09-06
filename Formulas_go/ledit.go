package main

// Ledit - Line editor for interactive commands
// Homepage: https://pauillac.inria.fr/~ddr/ledit/

import (
	"fmt"
	
	"os/exec"
)

func installLedit() {
	// Método 1: Descargar y extraer .tar.gz
	ledit_tar_url := "https://github.com/chetmurthy/ledit/archive/refs/tags/ledit-2-06.tar.gz"
	ledit_cmd_tar := exec.Command("curl", "-L", ledit_tar_url, "-o", "package.tar.gz")
	err := ledit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ledit_zip_url := "https://github.com/chetmurthy/ledit/archive/refs/tags/ledit-2-06.zip"
	ledit_cmd_zip := exec.Command("curl", "-L", ledit_zip_url, "-o", "package.zip")
	err = ledit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ledit_bin_url := "https://github.com/chetmurthy/ledit/archive/refs/tags/ledit-2-06.bin"
	ledit_cmd_bin := exec.Command("curl", "-L", ledit_bin_url, "-o", "binary.bin")
	err = ledit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ledit_src_url := "https://github.com/chetmurthy/ledit/archive/refs/tags/ledit-2-06.src.tar.gz"
	ledit_cmd_src := exec.Command("curl", "-L", ledit_src_url, "-o", "source.tar.gz")
	err = ledit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ledit_cmd_direct := exec.Command("./binary")
	err = ledit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml-findlib")
exec.Command("latte", "install", "ocaml-findlib")
	fmt.Println("Instalando dependencia: camlp-streams")
exec.Command("latte", "install", "camlp-streams")
	fmt.Println("Instalando dependencia: camlp5")
exec.Command("latte", "install", "camlp5")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
}
