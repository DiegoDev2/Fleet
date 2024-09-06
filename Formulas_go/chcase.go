package main

// Chcase - Perl file-renaming script
// Homepage: https://web.archive.org/web/20210514155949/http://primaledge.ca/chcase.html

import (
	"fmt"
	
	"os/exec"
)

func installChcase() {
	// Método 1: Descargar y extraer .tar.gz
	chcase_tar_url := "https://web.archive.org/web/20210731164711/http://primaledge.ca/chcase"
	chcase_cmd_tar := exec.Command("curl", "-L", chcase_tar_url, "-o", "package.tar.gz")
	err := chcase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chcase_zip_url := "https://web.archive.org/web/20210731164711/http://primaledge.ca/chcase"
	chcase_cmd_zip := exec.Command("curl", "-L", chcase_zip_url, "-o", "package.zip")
	err = chcase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chcase_bin_url := "https://web.archive.org/web/20210731164711/http://primaledge.ca/chcase"
	chcase_cmd_bin := exec.Command("curl", "-L", chcase_bin_url, "-o", "binary.bin")
	err = chcase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chcase_src_url := "https://web.archive.org/web/20210731164711/http://primaledge.ca/chcase"
	chcase_cmd_src := exec.Command("curl", "-L", chcase_src_url, "-o", "source.tar.gz")
	err = chcase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chcase_cmd_direct := exec.Command("./binary")
	err = chcase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
