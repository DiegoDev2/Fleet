package main

// Scdoc - Small man page generator
// Homepage: https://sr.ht/~sircmpwn/scdoc/

import (
	"fmt"
	
	"os/exec"
)

func installScdoc() {
	// Método 1: Descargar y extraer .tar.gz
	scdoc_tar_url := "https://git.sr.ht/~sircmpwn/scdoc/archive/1.11.3.tar.gz"
	scdoc_cmd_tar := exec.Command("curl", "-L", scdoc_tar_url, "-o", "package.tar.gz")
	err := scdoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scdoc_zip_url := "https://git.sr.ht/~sircmpwn/scdoc/archive/1.11.3.zip"
	scdoc_cmd_zip := exec.Command("curl", "-L", scdoc_zip_url, "-o", "package.zip")
	err = scdoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scdoc_bin_url := "https://git.sr.ht/~sircmpwn/scdoc/archive/1.11.3.bin"
	scdoc_cmd_bin := exec.Command("curl", "-L", scdoc_bin_url, "-o", "binary.bin")
	err = scdoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scdoc_src_url := "https://git.sr.ht/~sircmpwn/scdoc/archive/1.11.3.src.tar.gz"
	scdoc_cmd_src := exec.Command("curl", "-L", scdoc_src_url, "-o", "source.tar.gz")
	err = scdoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scdoc_cmd_direct := exec.Command("./binary")
	err = scdoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
