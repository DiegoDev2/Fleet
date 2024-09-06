package main

// Ocamlbuild - Generic build tool for OCaml
// Homepage: https://github.com/ocaml/ocamlbuild

import (
	"fmt"
	
	"os/exec"
)

func installOcamlbuild() {
	// Método 1: Descargar y extraer .tar.gz
	ocamlbuild_tar_url := "https://github.com/ocaml/ocamlbuild/archive/refs/tags/0.15.0.tar.gz"
	ocamlbuild_cmd_tar := exec.Command("curl", "-L", ocamlbuild_tar_url, "-o", "package.tar.gz")
	err := ocamlbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocamlbuild_zip_url := "https://github.com/ocaml/ocamlbuild/archive/refs/tags/0.15.0.zip"
	ocamlbuild_cmd_zip := exec.Command("curl", "-L", ocamlbuild_zip_url, "-o", "package.zip")
	err = ocamlbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocamlbuild_bin_url := "https://github.com/ocaml/ocamlbuild/archive/refs/tags/0.15.0.bin"
	ocamlbuild_cmd_bin := exec.Command("curl", "-L", ocamlbuild_bin_url, "-o", "binary.bin")
	err = ocamlbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocamlbuild_src_url := "https://github.com/ocaml/ocamlbuild/archive/refs/tags/0.15.0.src.tar.gz"
	ocamlbuild_cmd_src := exec.Command("curl", "-L", ocamlbuild_src_url, "-o", "source.tar.gz")
	err = ocamlbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocamlbuild_cmd_direct := exec.Command("./binary")
	err = ocamlbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
