package main

// OcamlNum - OCaml legacy Num library for arbitrary-precision arithmetic
// Homepage: https://github.com/ocaml/num

import (
	"fmt"
	
	"os/exec"
)

func installOcamlNum() {
	// Método 1: Descargar y extraer .tar.gz
	ocamlnum_tar_url := "https://github.com/ocaml/num/archive/refs/tags/v1.5.tar.gz"
	ocamlnum_cmd_tar := exec.Command("curl", "-L", ocamlnum_tar_url, "-o", "package.tar.gz")
	err := ocamlnum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocamlnum_zip_url := "https://github.com/ocaml/num/archive/refs/tags/v1.5.zip"
	ocamlnum_cmd_zip := exec.Command("curl", "-L", ocamlnum_zip_url, "-o", "package.zip")
	err = ocamlnum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocamlnum_bin_url := "https://github.com/ocaml/num/archive/refs/tags/v1.5.bin"
	ocamlnum_cmd_bin := exec.Command("curl", "-L", ocamlnum_bin_url, "-o", "binary.bin")
	err = ocamlnum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocamlnum_src_url := "https://github.com/ocaml/num/archive/refs/tags/v1.5.src.tar.gz"
	ocamlnum_cmd_src := exec.Command("curl", "-L", ocamlnum_src_url, "-o", "source.tar.gz")
	err = ocamlnum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocamlnum_cmd_direct := exec.Command("./binary")
	err = ocamlnum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml-findlib")
	exec.Command("latte", "install", "ocaml-findlib").Run()
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
