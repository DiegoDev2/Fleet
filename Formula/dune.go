package main

// Dune - Composable build system for OCaml
// Homepage: https://dune.build/

import (
	"fmt"
	
	"os/exec"
)

func installDune() {
	// Método 1: Descargar y extraer .tar.gz
	dune_tar_url := "https://github.com/ocaml/dune/releases/download/3.16.0/dune-3.16.0.tbz"
	dune_cmd_tar := exec.Command("curl", "-L", dune_tar_url, "-o", "package.tar.gz")
	err := dune_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dune_zip_url := "https://github.com/ocaml/dune/releases/download/3.16.0/dune-3.16.0.tbz"
	dune_cmd_zip := exec.Command("curl", "-L", dune_zip_url, "-o", "package.zip")
	err = dune_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dune_bin_url := "https://github.com/ocaml/dune/releases/download/3.16.0/dune-3.16.0.tbz"
	dune_cmd_bin := exec.Command("curl", "-L", dune_bin_url, "-o", "binary.bin")
	err = dune_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dune_src_url := "https://github.com/ocaml/dune/releases/download/3.16.0/dune-3.16.0.tbz"
	dune_cmd_src := exec.Command("curl", "-L", dune_src_url, "-o", "source.tar.gz")
	err = dune_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dune_cmd_direct := exec.Command("./binary")
	err = dune_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
