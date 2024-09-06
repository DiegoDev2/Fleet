package main

// Sbcl - Steel Bank Common Lisp system
// Homepage: https://www.sbcl.org/

import (
	"fmt"
	
	"os/exec"
)

func installSbcl() {
	// Método 1: Descargar y extraer .tar.gz
	sbcl_tar_url := "https://downloads.sourceforge.net/project/sbcl/sbcl/2.4.8/sbcl-2.4.8-source.tar.bz2"
	sbcl_cmd_tar := exec.Command("curl", "-L", sbcl_tar_url, "-o", "package.tar.gz")
	err := sbcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sbcl_zip_url := "https://downloads.sourceforge.net/project/sbcl/sbcl/2.4.8/sbcl-2.4.8-source.tar.bz2"
	sbcl_cmd_zip := exec.Command("curl", "-L", sbcl_zip_url, "-o", "package.zip")
	err = sbcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sbcl_bin_url := "https://downloads.sourceforge.net/project/sbcl/sbcl/2.4.8/sbcl-2.4.8-source.tar.bz2"
	sbcl_cmd_bin := exec.Command("curl", "-L", sbcl_bin_url, "-o", "binary.bin")
	err = sbcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sbcl_src_url := "https://downloads.sourceforge.net/project/sbcl/sbcl/2.4.8/sbcl-2.4.8-source.tar.bz2"
	sbcl_cmd_src := exec.Command("curl", "-L", sbcl_src_url, "-o", "source.tar.gz")
	err = sbcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sbcl_cmd_direct := exec.Command("./binary")
	err = sbcl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ecl")
exec.Command("latte", "install", "ecl")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
