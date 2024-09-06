package main

// Ol - Purely functional dialect of Lisp
// Homepage: https://yuriy-chumak.github.io/ol/

import (
	"fmt"
	
	"os/exec"
)

func installOl() {
	// Método 1: Descargar y extraer .tar.gz
	ol_tar_url := "https://github.com/yuriy-chumak/ol/archive/refs/tags/2.5.1.tar.gz"
	ol_cmd_tar := exec.Command("curl", "-L", ol_tar_url, "-o", "package.tar.gz")
	err := ol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ol_zip_url := "https://github.com/yuriy-chumak/ol/archive/refs/tags/2.5.1.zip"
	ol_cmd_zip := exec.Command("curl", "-L", ol_zip_url, "-o", "package.zip")
	err = ol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ol_bin_url := "https://github.com/yuriy-chumak/ol/archive/refs/tags/2.5.1.bin"
	ol_cmd_bin := exec.Command("curl", "-L", ol_bin_url, "-o", "binary.bin")
	err = ol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ol_src_url := "https://github.com/yuriy-chumak/ol/archive/refs/tags/2.5.1.src.tar.gz"
	ol_cmd_src := exec.Command("curl", "-L", ol_src_url, "-o", "source.tar.gz")
	err = ol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ol_cmd_direct := exec.Command("./binary")
	err = ol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
