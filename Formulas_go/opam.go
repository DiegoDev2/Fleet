package main

// Opam - OCaml package manager
// Homepage: https://opam.ocaml.org

import (
	"fmt"
	
	"os/exec"
)

func installOpam() {
	// Método 1: Descargar y extraer .tar.gz
	opam_tar_url := "https://github.com/ocaml/opam/releases/download/2.2.1/opam-full-2.2.1.tar.gz"
	opam_cmd_tar := exec.Command("curl", "-L", opam_tar_url, "-o", "package.tar.gz")
	err := opam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opam_zip_url := "https://github.com/ocaml/opam/releases/download/2.2.1/opam-full-2.2.1.zip"
	opam_cmd_zip := exec.Command("curl", "-L", opam_zip_url, "-o", "package.zip")
	err = opam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opam_bin_url := "https://github.com/ocaml/opam/releases/download/2.2.1/opam-full-2.2.1.bin"
	opam_cmd_bin := exec.Command("curl", "-L", opam_bin_url, "-o", "binary.bin")
	err = opam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opam_src_url := "https://github.com/ocaml/opam/releases/download/2.2.1/opam-full-2.2.1.src.tar.gz"
	opam_cmd_src := exec.Command("curl", "-L", opam_src_url, "-o", "source.tar.gz")
	err = opam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opam_cmd_direct := exec.Command("./binary")
	err = opam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: gpatch")
exec.Command("latte", "install", "gpatch")
}
