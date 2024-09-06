package main

// Crfsuite - Fast implementation of conditional random fields
// Homepage: https://www.chokkan.org/software/crfsuite/

import (
	"fmt"
	
	"os/exec"
)

func installCrfsuite() {
	// Método 1: Descargar y extraer .tar.gz
	crfsuite_tar_url := "https://github.com/chokkan/crfsuite/archive/refs/tags/0.12.tar.gz"
	crfsuite_cmd_tar := exec.Command("curl", "-L", crfsuite_tar_url, "-o", "package.tar.gz")
	err := crfsuite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crfsuite_zip_url := "https://github.com/chokkan/crfsuite/archive/refs/tags/0.12.zip"
	crfsuite_cmd_zip := exec.Command("curl", "-L", crfsuite_zip_url, "-o", "package.zip")
	err = crfsuite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crfsuite_bin_url := "https://github.com/chokkan/crfsuite/archive/refs/tags/0.12.bin"
	crfsuite_cmd_bin := exec.Command("curl", "-L", crfsuite_bin_url, "-o", "binary.bin")
	err = crfsuite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crfsuite_src_url := "https://github.com/chokkan/crfsuite/archive/refs/tags/0.12.src.tar.gz"
	crfsuite_cmd_src := exec.Command("curl", "-L", crfsuite_src_url, "-o", "source.tar.gz")
	err = crfsuite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crfsuite_cmd_direct := exec.Command("./binary")
	err = crfsuite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: liblbfgs")
exec.Command("latte", "install", "liblbfgs")
}
