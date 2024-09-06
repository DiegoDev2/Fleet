package main

// Ppl - Parma Polyhedra Library: numerical abstractions for analysis, verification
// Homepage: https://www.bugseng.com/ppl

import (
	"fmt"
	
	"os/exec"
)

func installPpl() {
	// Método 1: Descargar y extraer .tar.gz
	ppl_tar_url := "https://www.bugseng.com/products/ppl/download/ftp/releases/1.2/ppl-1.2.tar.xz"
	ppl_cmd_tar := exec.Command("curl", "-L", ppl_tar_url, "-o", "package.tar.gz")
	err := ppl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ppl_zip_url := "https://www.bugseng.com/products/ppl/download/ftp/releases/1.2/ppl-1.2.tar.xz"
	ppl_cmd_zip := exec.Command("curl", "-L", ppl_zip_url, "-o", "package.zip")
	err = ppl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ppl_bin_url := "https://www.bugseng.com/products/ppl/download/ftp/releases/1.2/ppl-1.2.tar.xz"
	ppl_cmd_bin := exec.Command("curl", "-L", ppl_bin_url, "-o", "binary.bin")
	err = ppl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ppl_src_url := "https://www.bugseng.com/products/ppl/download/ftp/releases/1.2/ppl-1.2.tar.xz"
	ppl_cmd_src := exec.Command("curl", "-L", ppl_src_url, "-o", "source.tar.gz")
	err = ppl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ppl_cmd_direct := exec.Command("./binary")
	err = ppl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: m4")
exec.Command("latte", "install", "m4")
}
