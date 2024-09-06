package main

// Scws - Simple Chinese Word Segmentation
// Homepage: https://github.com/hightman/scws

import (
	"fmt"
	
	"os/exec"
)

func installScws() {
	// Método 1: Descargar y extraer .tar.gz
	scws_tar_url := "http://www.xunsearch.com/scws/down/scws-1.2.3.tar.bz2"
	scws_cmd_tar := exec.Command("curl", "-L", scws_tar_url, "-o", "package.tar.gz")
	err := scws_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scws_zip_url := "http://www.xunsearch.com/scws/down/scws-1.2.3.tar.bz2"
	scws_cmd_zip := exec.Command("curl", "-L", scws_zip_url, "-o", "package.zip")
	err = scws_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scws_bin_url := "http://www.xunsearch.com/scws/down/scws-1.2.3.tar.bz2"
	scws_cmd_bin := exec.Command("curl", "-L", scws_bin_url, "-o", "binary.bin")
	err = scws_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scws_src_url := "http://www.xunsearch.com/scws/down/scws-1.2.3.tar.bz2"
	scws_cmd_src := exec.Command("curl", "-L", scws_src_url, "-o", "source.tar.gz")
	err = scws_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scws_cmd_direct := exec.Command("./binary")
	err = scws_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
