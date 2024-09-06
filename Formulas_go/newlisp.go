package main

// Newlisp - Lisp-like, general-purpose scripting language
// Homepage: http://www.newlisp.org/

import (
	"fmt"
	
	"os/exec"
)

func installNewlisp() {
	// Método 1: Descargar y extraer .tar.gz
	newlisp_tar_url := "http://www.newlisp.org/downloads/newlisp-10.7.5.tgz"
	newlisp_cmd_tar := exec.Command("curl", "-L", newlisp_tar_url, "-o", "package.tar.gz")
	err := newlisp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	newlisp_zip_url := "http://www.newlisp.org/downloads/newlisp-10.7.5.tgz"
	newlisp_cmd_zip := exec.Command("curl", "-L", newlisp_zip_url, "-o", "package.zip")
	err = newlisp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	newlisp_bin_url := "http://www.newlisp.org/downloads/newlisp-10.7.5.tgz"
	newlisp_cmd_bin := exec.Command("curl", "-L", newlisp_bin_url, "-o", "binary.bin")
	err = newlisp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	newlisp_src_url := "http://www.newlisp.org/downloads/newlisp-10.7.5.tgz"
	newlisp_cmd_src := exec.Command("curl", "-L", newlisp_src_url, "-o", "source.tar.gz")
	err = newlisp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	newlisp_cmd_direct := exec.Command("./binary")
	err = newlisp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
