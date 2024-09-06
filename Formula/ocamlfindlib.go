package main

// OcamlFindlib - OCaml library manager
// Homepage: http://projects.camlcity.org/projects/findlib.html

import (
	"fmt"
	
	"os/exec"
)

func installOcamlFindlib() {
	// Método 1: Descargar y extraer .tar.gz
	ocamlfindlib_tar_url := "http://download.camlcity.org/download/findlib-1.9.6.tar.gz"
	ocamlfindlib_cmd_tar := exec.Command("curl", "-L", ocamlfindlib_tar_url, "-o", "package.tar.gz")
	err := ocamlfindlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocamlfindlib_zip_url := "http://download.camlcity.org/download/findlib-1.9.6.zip"
	ocamlfindlib_cmd_zip := exec.Command("curl", "-L", ocamlfindlib_zip_url, "-o", "package.zip")
	err = ocamlfindlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocamlfindlib_bin_url := "http://download.camlcity.org/download/findlib-1.9.6.bin"
	ocamlfindlib_cmd_bin := exec.Command("curl", "-L", ocamlfindlib_bin_url, "-o", "binary.bin")
	err = ocamlfindlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocamlfindlib_src_url := "http://download.camlcity.org/download/findlib-1.9.6.src.tar.gz"
	ocamlfindlib_cmd_src := exec.Command("curl", "-L", ocamlfindlib_src_url, "-o", "source.tar.gz")
	err = ocamlfindlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocamlfindlib_cmd_direct := exec.Command("./binary")
	err = ocamlfindlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
