package main

// Hmmer - Build profile HMMs and scan against sequence databases
// Homepage: http://hmmer.org/

import (
	"fmt"
	
	"os/exec"
)

func installHmmer() {
	// Método 1: Descargar y extraer .tar.gz
	hmmer_tar_url := "http://eddylab.org/software/hmmer/hmmer-3.4.tar.gz"
	hmmer_cmd_tar := exec.Command("curl", "-L", hmmer_tar_url, "-o", "package.tar.gz")
	err := hmmer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hmmer_zip_url := "http://eddylab.org/software/hmmer/hmmer-3.4.zip"
	hmmer_cmd_zip := exec.Command("curl", "-L", hmmer_zip_url, "-o", "package.zip")
	err = hmmer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hmmer_bin_url := "http://eddylab.org/software/hmmer/hmmer-3.4.bin"
	hmmer_cmd_bin := exec.Command("curl", "-L", hmmer_bin_url, "-o", "binary.bin")
	err = hmmer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hmmer_src_url := "http://eddylab.org/software/hmmer/hmmer-3.4.src.tar.gz"
	hmmer_cmd_src := exec.Command("curl", "-L", hmmer_src_url, "-o", "source.tar.gz")
	err = hmmer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hmmer_cmd_direct := exec.Command("./binary")
	err = hmmer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
