package main

// Coq - Proof assistant for higher-order logic
// Homepage: https://coq.inria.fr/

import (
	"fmt"
	
	"os/exec"
)

func installCoq() {
	// Método 1: Descargar y extraer .tar.gz
	coq_tar_url := "https://github.com/coq/coq/releases/download/V8.20.0/coq-8.20.0.tar.gz"
	coq_cmd_tar := exec.Command("curl", "-L", coq_tar_url, "-o", "package.tar.gz")
	err := coq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coq_zip_url := "https://github.com/coq/coq/releases/download/V8.20.0/coq-8.20.0.zip"
	coq_cmd_zip := exec.Command("curl", "-L", coq_zip_url, "-o", "package.zip")
	err = coq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coq_bin_url := "https://github.com/coq/coq/releases/download/V8.20.0/coq-8.20.0.bin"
	coq_cmd_bin := exec.Command("curl", "-L", coq_bin_url, "-o", "binary.bin")
	err = coq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coq_src_url := "https://github.com/coq/coq/releases/download/V8.20.0/coq-8.20.0.src.tar.gz"
	coq_cmd_src := exec.Command("curl", "-L", coq_src_url, "-o", "source.tar.gz")
	err = coq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coq_cmd_direct := exec.Command("./binary")
	err = coq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dune")
exec.Command("latte", "install", "dune")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: ocaml-findlib")
exec.Command("latte", "install", "ocaml-findlib")
	fmt.Println("Instalando dependencia: ocaml-zarith")
exec.Command("latte", "install", "ocaml-zarith")
}
