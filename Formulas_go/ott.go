package main

// Ott - Tool for writing definitions of programming languages and calculi
// Homepage: https://www.cl.cam.ac.uk/~pes20/ott/

import (
	"fmt"
	
	"os/exec"
)

func installOtt() {
	// Método 1: Descargar y extraer .tar.gz
	ott_tar_url := "https://github.com/ott-lang/ott/archive/refs/tags/0.33.tar.gz"
	ott_cmd_tar := exec.Command("curl", "-L", ott_tar_url, "-o", "package.tar.gz")
	err := ott_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ott_zip_url := "https://github.com/ott-lang/ott/archive/refs/tags/0.33.zip"
	ott_cmd_zip := exec.Command("curl", "-L", ott_zip_url, "-o", "package.zip")
	err = ott_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ott_bin_url := "https://github.com/ott-lang/ott/archive/refs/tags/0.33.bin"
	ott_cmd_bin := exec.Command("curl", "-L", ott_bin_url, "-o", "binary.bin")
	err = ott_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ott_src_url := "https://github.com/ott-lang/ott/archive/refs/tags/0.33.src.tar.gz"
	ott_cmd_src := exec.Command("curl", "-L", ott_src_url, "-o", "source.tar.gz")
	err = ott_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ott_cmd_direct := exec.Command("./binary")
	err = ott_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: opam")
exec.Command("latte", "install", "opam")
}
