package main

// Hevea - LaTeX-to-HTML translator
// Homepage: https://hevea.inria.fr/

import (
	"fmt"
	
	"os/exec"
)

func installHevea() {
	// Método 1: Descargar y extraer .tar.gz
	hevea_tar_url := "https://hevea.inria.fr/old/hevea-2.36.tar.gz"
	hevea_cmd_tar := exec.Command("curl", "-L", hevea_tar_url, "-o", "package.tar.gz")
	err := hevea_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hevea_zip_url := "https://hevea.inria.fr/old/hevea-2.36.zip"
	hevea_cmd_zip := exec.Command("curl", "-L", hevea_zip_url, "-o", "package.zip")
	err = hevea_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hevea_bin_url := "https://hevea.inria.fr/old/hevea-2.36.bin"
	hevea_cmd_bin := exec.Command("curl", "-L", hevea_bin_url, "-o", "binary.bin")
	err = hevea_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hevea_src_url := "https://hevea.inria.fr/old/hevea-2.36.src.tar.gz"
	hevea_cmd_src := exec.Command("curl", "-L", hevea_src_url, "-o", "source.tar.gz")
	err = hevea_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hevea_cmd_direct := exec.Command("./binary")
	err = hevea_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocamlbuild")
exec.Command("latte", "install", "ocamlbuild")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
}
