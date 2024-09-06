package main

// Inform6 - Design system for interactive fiction
// Homepage: https://inform-fiction.org/inform6.html

import (
	"fmt"
	
	"os/exec"
)

func installInform6() {
	// Método 1: Descargar y extraer .tar.gz
	inform6_tar_url := "https://ifarchive.org/if-archive/infocom/compilers/inform6/source/inform-6.42-r4.tar.gz"
	inform6_cmd_tar := exec.Command("curl", "-L", inform6_tar_url, "-o", "package.tar.gz")
	err := inform6_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inform6_zip_url := "https://ifarchive.org/if-archive/infocom/compilers/inform6/source/inform-6.42-r4.zip"
	inform6_cmd_zip := exec.Command("curl", "-L", inform6_zip_url, "-o", "package.zip")
	err = inform6_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inform6_bin_url := "https://ifarchive.org/if-archive/infocom/compilers/inform6/source/inform-6.42-r4.bin"
	inform6_cmd_bin := exec.Command("curl", "-L", inform6_bin_url, "-o", "binary.bin")
	err = inform6_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inform6_src_url := "https://ifarchive.org/if-archive/infocom/compilers/inform6/source/inform-6.42-r4.src.tar.gz"
	inform6_cmd_src := exec.Command("curl", "-L", inform6_src_url, "-o", "source.tar.gz")
	err = inform6_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inform6_cmd_direct := exec.Command("./binary")
	err = inform6_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
