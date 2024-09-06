package main

// Ocaml - General purpose programming language in the ML family
// Homepage: https://ocaml.org/

import (
	"fmt"
	
	"os/exec"
)

func installOcaml() {
	// Método 1: Descargar y extraer .tar.gz
	ocaml_tar_url := "https://caml.inria.fr/pub/distrib/ocaml-5.2/ocaml-5.2.0.tar.xz"
	ocaml_cmd_tar := exec.Command("curl", "-L", ocaml_tar_url, "-o", "package.tar.gz")
	err := ocaml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocaml_zip_url := "https://caml.inria.fr/pub/distrib/ocaml-5.2/ocaml-5.2.0.tar.xz"
	ocaml_cmd_zip := exec.Command("curl", "-L", ocaml_zip_url, "-o", "package.zip")
	err = ocaml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocaml_bin_url := "https://caml.inria.fr/pub/distrib/ocaml-5.2/ocaml-5.2.0.tar.xz"
	ocaml_cmd_bin := exec.Command("curl", "-L", ocaml_bin_url, "-o", "binary.bin")
	err = ocaml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocaml_src_url := "https://caml.inria.fr/pub/distrib/ocaml-5.2/ocaml-5.2.0.tar.xz"
	ocaml_cmd_src := exec.Command("curl", "-L", ocaml_src_url, "-o", "source.tar.gz")
	err = ocaml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocaml_cmd_direct := exec.Command("./binary")
	err = ocaml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
