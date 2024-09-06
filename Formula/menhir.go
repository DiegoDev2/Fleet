package main

// Menhir - LR(1) parser generator for the OCaml programming language
// Homepage: http://cristal.inria.fr/~fpottier/menhir

import (
	"fmt"
	
	"os/exec"
)

func installMenhir() {
	// Método 1: Descargar y extraer .tar.gz
	menhir_tar_url := "https://gitlab.inria.fr/fpottier/menhir/-/archive/20240715/menhir-20240715.tar.bz2"
	menhir_cmd_tar := exec.Command("curl", "-L", menhir_tar_url, "-o", "package.tar.gz")
	err := menhir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	menhir_zip_url := "https://gitlab.inria.fr/fpottier/menhir/-/archive/20240715/menhir-20240715.tar.bz2"
	menhir_cmd_zip := exec.Command("curl", "-L", menhir_zip_url, "-o", "package.zip")
	err = menhir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	menhir_bin_url := "https://gitlab.inria.fr/fpottier/menhir/-/archive/20240715/menhir-20240715.tar.bz2"
	menhir_cmd_bin := exec.Command("curl", "-L", menhir_bin_url, "-o", "binary.bin")
	err = menhir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	menhir_src_url := "https://gitlab.inria.fr/fpottier/menhir/-/archive/20240715/menhir-20240715.tar.bz2"
	menhir_cmd_src := exec.Command("curl", "-L", menhir_src_url, "-o", "source.tar.gz")
	err = menhir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	menhir_cmd_direct := exec.Command("./binary")
	err = menhir_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dune")
	exec.Command("latte", "install", "dune").Run()
	fmt.Println("Instalando dependencia: ocamlbuild")
	exec.Command("latte", "install", "ocamlbuild").Run()
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
