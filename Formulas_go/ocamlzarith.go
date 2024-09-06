package main

// OcamlZarith - OCaml library for arbitrary-precision arithmetic
// Homepage: https://github.com/ocaml/Zarith

import (
	"fmt"
	
	"os/exec"
)

func installOcamlZarith() {
	// Método 1: Descargar y extraer .tar.gz
	ocamlzarith_tar_url := "https://github.com/ocaml/Zarith/archive/refs/tags/release-1.14.tar.gz"
	ocamlzarith_cmd_tar := exec.Command("curl", "-L", ocamlzarith_tar_url, "-o", "package.tar.gz")
	err := ocamlzarith_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocamlzarith_zip_url := "https://github.com/ocaml/Zarith/archive/refs/tags/release-1.14.zip"
	ocamlzarith_cmd_zip := exec.Command("curl", "-L", ocamlzarith_zip_url, "-o", "package.zip")
	err = ocamlzarith_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocamlzarith_bin_url := "https://github.com/ocaml/Zarith/archive/refs/tags/release-1.14.bin"
	ocamlzarith_cmd_bin := exec.Command("curl", "-L", ocamlzarith_bin_url, "-o", "binary.bin")
	err = ocamlzarith_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocamlzarith_src_url := "https://github.com/ocaml/Zarith/archive/refs/tags/release-1.14.src.tar.gz"
	ocamlzarith_cmd_src := exec.Command("curl", "-L", ocamlzarith_src_url, "-o", "source.tar.gz")
	err = ocamlzarith_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocamlzarith_cmd_direct := exec.Command("./binary")
	err = ocamlzarith_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml-findlib")
exec.Command("latte", "install", "ocaml-findlib")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
}
