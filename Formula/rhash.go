package main

// Rhash - Utility for computing and verifying hash sums of files
// Homepage: https://sourceforge.net/projects/rhash/

import (
	"fmt"
	
	"os/exec"
)

func installRhash() {
	// Método 1: Descargar y extraer .tar.gz
	rhash_tar_url := "https://downloads.sourceforge.net/project/rhash/rhash/1.4.4/rhash-1.4.4-src.tar.gz"
	rhash_cmd_tar := exec.Command("curl", "-L", rhash_tar_url, "-o", "package.tar.gz")
	err := rhash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rhash_zip_url := "https://downloads.sourceforge.net/project/rhash/rhash/1.4.4/rhash-1.4.4-src.zip"
	rhash_cmd_zip := exec.Command("curl", "-L", rhash_zip_url, "-o", "package.zip")
	err = rhash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rhash_bin_url := "https://downloads.sourceforge.net/project/rhash/rhash/1.4.4/rhash-1.4.4-src.bin"
	rhash_cmd_bin := exec.Command("curl", "-L", rhash_bin_url, "-o", "binary.bin")
	err = rhash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rhash_src_url := "https://downloads.sourceforge.net/project/rhash/rhash/1.4.4/rhash-1.4.4-src.src.tar.gz"
	rhash_cmd_src := exec.Command("curl", "-L", rhash_src_url, "-o", "source.tar.gz")
	err = rhash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rhash_cmd_direct := exec.Command("./binary")
	err = rhash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
