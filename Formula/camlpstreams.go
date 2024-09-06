package main

// CamlpStreams - Stream and Genlex libraries for use with Camlp4 and Camlp5
// Homepage: https://github.com/ocaml/camlp-streams

import (
	"fmt"
	
	"os/exec"
)

func installCamlpStreams() {
	// Método 1: Descargar y extraer .tar.gz
	camlpstreams_tar_url := "https://github.com/ocaml/camlp-streams/archive/refs/tags/v5.0.1.tar.gz"
	camlpstreams_cmd_tar := exec.Command("curl", "-L", camlpstreams_tar_url, "-o", "package.tar.gz")
	err := camlpstreams_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	camlpstreams_zip_url := "https://github.com/ocaml/camlp-streams/archive/refs/tags/v5.0.1.zip"
	camlpstreams_cmd_zip := exec.Command("curl", "-L", camlpstreams_zip_url, "-o", "package.zip")
	err = camlpstreams_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	camlpstreams_bin_url := "https://github.com/ocaml/camlp-streams/archive/refs/tags/v5.0.1.bin"
	camlpstreams_cmd_bin := exec.Command("curl", "-L", camlpstreams_bin_url, "-o", "binary.bin")
	err = camlpstreams_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	camlpstreams_src_url := "https://github.com/ocaml/camlp-streams/archive/refs/tags/v5.0.1.src.tar.gz"
	camlpstreams_cmd_src := exec.Command("curl", "-L", camlpstreams_src_url, "-o", "source.tar.gz")
	err = camlpstreams_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	camlpstreams_cmd_direct := exec.Command("./binary")
	err = camlpstreams_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dune")
	exec.Command("latte", "install", "dune").Run()
	fmt.Println("Instalando dependencia: ocaml-findlib")
	exec.Command("latte", "install", "ocaml-findlib").Run()
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
