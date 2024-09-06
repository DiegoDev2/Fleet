package main

// Abcl - Armed Bear Common Lisp: a full implementation of Common Lisp
// Homepage: https://abcl.org/

import (
	"fmt"
	
	"os/exec"
)

func installAbcl() {
	// Método 1: Descargar y extraer .tar.gz
	abcl_tar_url := "https://abcl.org/releases/1.9.2/abcl-src-1.9.2.tar.gz"
	abcl_cmd_tar := exec.Command("curl", "-L", abcl_tar_url, "-o", "package.tar.gz")
	err := abcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abcl_zip_url := "https://abcl.org/releases/1.9.2/abcl-src-1.9.2.zip"
	abcl_cmd_zip := exec.Command("curl", "-L", abcl_zip_url, "-o", "package.zip")
	err = abcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abcl_bin_url := "https://abcl.org/releases/1.9.2/abcl-src-1.9.2.bin"
	abcl_cmd_bin := exec.Command("curl", "-L", abcl_bin_url, "-o", "binary.bin")
	err = abcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abcl_src_url := "https://abcl.org/releases/1.9.2/abcl-src-1.9.2.src.tar.gz"
	abcl_cmd_src := exec.Command("curl", "-L", abcl_src_url, "-o", "source.tar.gz")
	err = abcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abcl_cmd_direct := exec.Command("./binary")
	err = abcl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
	exec.Command("latte", "install", "ant").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: rlwrap")
	exec.Command("latte", "install", "rlwrap").Run()
}
