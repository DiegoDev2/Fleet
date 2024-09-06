package main

// Camlp5 - Preprocessor and pretty-printer for OCaml
// Homepage: https://camlp5.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installCamlp5() {
	// Método 1: Descargar y extraer .tar.gz
	camlp5_tar_url := "https://github.com/camlp5/camlp5/archive/refs/tags/8.03.00.tar.gz"
	camlp5_cmd_tar := exec.Command("curl", "-L", camlp5_tar_url, "-o", "package.tar.gz")
	err := camlp5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	camlp5_zip_url := "https://github.com/camlp5/camlp5/archive/refs/tags/8.03.00.zip"
	camlp5_cmd_zip := exec.Command("curl", "-L", camlp5_zip_url, "-o", "package.zip")
	err = camlp5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	camlp5_bin_url := "https://github.com/camlp5/camlp5/archive/refs/tags/8.03.00.bin"
	camlp5_cmd_bin := exec.Command("curl", "-L", camlp5_bin_url, "-o", "binary.bin")
	err = camlp5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	camlp5_src_url := "https://github.com/camlp5/camlp5/archive/refs/tags/8.03.00.src.tar.gz"
	camlp5_cmd_src := exec.Command("curl", "-L", camlp5_src_url, "-o", "source.tar.gz")
	err = camlp5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	camlp5_cmd_direct := exec.Command("./binary")
	err = camlp5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml-findlib")
exec.Command("latte", "install", "ocaml-findlib")
	fmt.Println("Instalando dependencia: opam")
exec.Command("latte", "install", "opam")
	fmt.Println("Instalando dependencia: camlp-streams")
exec.Command("latte", "install", "camlp-streams")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
}
