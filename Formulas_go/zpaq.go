package main

// Zpaq - Incremental, journaling command-line archiver
// Homepage: https://mattmahoney.net/dc/zpaq.html

import (
	"fmt"
	
	"os/exec"
)

func installZpaq() {
	// Método 1: Descargar y extraer .tar.gz
	zpaq_tar_url := "https://mattmahoney.net/dc/zpaq715.zip"
	zpaq_cmd_tar := exec.Command("curl", "-L", zpaq_tar_url, "-o", "package.tar.gz")
	err := zpaq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zpaq_zip_url := "https://mattmahoney.net/dc/zpaq715.zip"
	zpaq_cmd_zip := exec.Command("curl", "-L", zpaq_zip_url, "-o", "package.zip")
	err = zpaq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zpaq_bin_url := "https://mattmahoney.net/dc/zpaq715.zip"
	zpaq_cmd_bin := exec.Command("curl", "-L", zpaq_bin_url, "-o", "binary.bin")
	err = zpaq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zpaq_src_url := "https://mattmahoney.net/dc/zpaq715.zip"
	zpaq_cmd_src := exec.Command("curl", "-L", zpaq_src_url, "-o", "source.tar.gz")
	err = zpaq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zpaq_cmd_direct := exec.Command("./binary")
	err = zpaq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
