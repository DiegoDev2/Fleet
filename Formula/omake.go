package main

// Omake - Build system designed for scalability, portability, and concision
// Homepage: http://projects.camlcity.org/projects/omake.html

import (
	"fmt"
	
	"os/exec"
)

func installOmake() {
	// Método 1: Descargar y extraer .tar.gz
	omake_tar_url := "https://github.com/ocaml-omake/omake/archive/refs/tags/omake-0.10.6.tar.gz"
	omake_cmd_tar := exec.Command("curl", "-L", omake_tar_url, "-o", "package.tar.gz")
	err := omake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	omake_zip_url := "https://github.com/ocaml-omake/omake/archive/refs/tags/omake-0.10.6.zip"
	omake_cmd_zip := exec.Command("curl", "-L", omake_zip_url, "-o", "package.zip")
	err = omake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	omake_bin_url := "https://github.com/ocaml-omake/omake/archive/refs/tags/omake-0.10.6.bin"
	omake_cmd_bin := exec.Command("curl", "-L", omake_bin_url, "-o", "binary.bin")
	err = omake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	omake_src_url := "https://github.com/ocaml-omake/omake/archive/refs/tags/omake-0.10.6.src.tar.gz"
	omake_cmd_src := exec.Command("curl", "-L", omake_src_url, "-o", "source.tar.gz")
	err = omake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	omake_cmd_direct := exec.Command("./binary")
	err = omake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
	fmt.Println("Instalando dependencia: ocaml-findlib")
	exec.Command("latte", "install", "ocaml-findlib").Run()
}
