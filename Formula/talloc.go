package main

// Talloc - Hierarchical, reference-counted memory pool with destructors
// Homepage: https://talloc.samba.org/

import (
	"fmt"
	
	"os/exec"
)

func installTalloc() {
	// Método 1: Descargar y extraer .tar.gz
	talloc_tar_url := "https://www.samba.org/ftp/talloc/talloc-2.4.2.tar.gz"
	talloc_cmd_tar := exec.Command("curl", "-L", talloc_tar_url, "-o", "package.tar.gz")
	err := talloc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	talloc_zip_url := "https://www.samba.org/ftp/talloc/talloc-2.4.2.zip"
	talloc_cmd_zip := exec.Command("curl", "-L", talloc_zip_url, "-o", "package.zip")
	err = talloc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	talloc_bin_url := "https://www.samba.org/ftp/talloc/talloc-2.4.2.bin"
	talloc_cmd_bin := exec.Command("curl", "-L", talloc_bin_url, "-o", "binary.bin")
	err = talloc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	talloc_src_url := "https://www.samba.org/ftp/talloc/talloc-2.4.2.src.tar.gz"
	talloc_cmd_src := exec.Command("curl", "-L", talloc_src_url, "-o", "source.tar.gz")
	err = talloc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	talloc_cmd_direct := exec.Command("./binary")
	err = talloc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
