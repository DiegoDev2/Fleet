package main

// OcamlAT4 - General purpose programming language in the ML family
// Homepage: https://ocaml.org/

import (
	"fmt"
	
	"os/exec"
)

func installOcamlAT4() {
	// Método 1: Descargar y extraer .tar.gz
	ocamlat4_tar_url := "https://caml.inria.fr/pub/distrib/ocaml-4.14/ocaml-4.14.2.tar.xz"
	ocamlat4_cmd_tar := exec.Command("curl", "-L", ocamlat4_tar_url, "-o", "package.tar.gz")
	err := ocamlat4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocamlat4_zip_url := "https://caml.inria.fr/pub/distrib/ocaml-4.14/ocaml-4.14.2.tar.xz"
	ocamlat4_cmd_zip := exec.Command("curl", "-L", ocamlat4_zip_url, "-o", "package.zip")
	err = ocamlat4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocamlat4_bin_url := "https://caml.inria.fr/pub/distrib/ocaml-4.14/ocaml-4.14.2.tar.xz"
	ocamlat4_cmd_bin := exec.Command("curl", "-L", ocamlat4_bin_url, "-o", "binary.bin")
	err = ocamlat4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocamlat4_src_url := "https://caml.inria.fr/pub/distrib/ocaml-4.14/ocaml-4.14.2.tar.xz"
	ocamlat4_cmd_src := exec.Command("curl", "-L", ocamlat4_src_url, "-o", "source.tar.gz")
	err = ocamlat4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocamlat4_cmd_direct := exec.Command("./binary")
	err = ocamlat4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
