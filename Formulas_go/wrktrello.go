package main

// WrkTrello - Command-line interface to Trello
// Homepage: https://github.com/blangel/wrk

import (
	"fmt"
	
	"os/exec"
)

func installWrkTrello() {
	// Método 1: Descargar y extraer .tar.gz
	wrktrello_tar_url := "https://github.s3.amazonaws.com/downloads/blangel/wrk/wrk-1.0.1.tar.gz"
	wrktrello_cmd_tar := exec.Command("curl", "-L", wrktrello_tar_url, "-o", "package.tar.gz")
	err := wrktrello_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wrktrello_zip_url := "https://github.s3.amazonaws.com/downloads/blangel/wrk/wrk-1.0.1.zip"
	wrktrello_cmd_zip := exec.Command("curl", "-L", wrktrello_zip_url, "-o", "package.zip")
	err = wrktrello_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wrktrello_bin_url := "https://github.s3.amazonaws.com/downloads/blangel/wrk/wrk-1.0.1.bin"
	wrktrello_cmd_bin := exec.Command("curl", "-L", wrktrello_bin_url, "-o", "binary.bin")
	err = wrktrello_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wrktrello_src_url := "https://github.s3.amazonaws.com/downloads/blangel/wrk/wrk-1.0.1.src.tar.gz"
	wrktrello_cmd_src := exec.Command("curl", "-L", wrktrello_src_url, "-o", "source.tar.gz")
	err = wrktrello_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wrktrello_cmd_direct := exec.Command("./binary")
	err = wrktrello_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
