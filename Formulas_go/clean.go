package main

// Clean - Search for files matching a regex and delete them
// Homepage: https://clean.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installClean() {
	// Método 1: Descargar y extraer .tar.gz
	clean_tar_url := "https://downloads.sourceforge.net/project/clean/clean/3.4/clean-3.4.tar.bz2"
	clean_cmd_tar := exec.Command("curl", "-L", clean_tar_url, "-o", "package.tar.gz")
	err := clean_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clean_zip_url := "https://downloads.sourceforge.net/project/clean/clean/3.4/clean-3.4.tar.bz2"
	clean_cmd_zip := exec.Command("curl", "-L", clean_zip_url, "-o", "package.zip")
	err = clean_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clean_bin_url := "https://downloads.sourceforge.net/project/clean/clean/3.4/clean-3.4.tar.bz2"
	clean_cmd_bin := exec.Command("curl", "-L", clean_bin_url, "-o", "binary.bin")
	err = clean_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clean_src_url := "https://downloads.sourceforge.net/project/clean/clean/3.4/clean-3.4.tar.bz2"
	clean_cmd_src := exec.Command("curl", "-L", clean_src_url, "-o", "source.tar.gz")
	err = clean_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clean_cmd_direct := exec.Command("./binary")
	err = clean_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
