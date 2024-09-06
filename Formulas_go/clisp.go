package main

// Clisp - GNU CLISP, a Common Lisp implementation
// Homepage: https://clisp.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installClisp() {
	// Método 1: Descargar y extraer .tar.gz
	clisp_tar_url := "https://alpha.gnu.org/gnu/clisp/clisp-2.49.92.tar.bz2"
	clisp_cmd_tar := exec.Command("curl", "-L", clisp_tar_url, "-o", "package.tar.gz")
	err := clisp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clisp_zip_url := "https://alpha.gnu.org/gnu/clisp/clisp-2.49.92.tar.bz2"
	clisp_cmd_zip := exec.Command("curl", "-L", clisp_zip_url, "-o", "package.zip")
	err = clisp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clisp_bin_url := "https://alpha.gnu.org/gnu/clisp/clisp-2.49.92.tar.bz2"
	clisp_cmd_bin := exec.Command("curl", "-L", clisp_bin_url, "-o", "binary.bin")
	err = clisp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clisp_src_url := "https://alpha.gnu.org/gnu/clisp/clisp-2.49.92.tar.bz2"
	clisp_cmd_src := exec.Command("curl", "-L", clisp_src_url, "-o", "source.tar.gz")
	err = clisp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clisp_cmd_direct := exec.Command("./binary")
	err = clisp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libsigsegv")
exec.Command("latte", "install", "libsigsegv")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
