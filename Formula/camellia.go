package main

// Camellia - Image Processing & Computer Vision library written in C
// Homepage: https://camellia.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCamellia() {
	// Método 1: Descargar y extraer .tar.gz
	camellia_tar_url := "https://downloads.sourceforge.net/project/camellia/Unix_Linux%20Distribution/v2.7.0/CamelliaLib-2.7.0.tar.gz"
	camellia_cmd_tar := exec.Command("curl", "-L", camellia_tar_url, "-o", "package.tar.gz")
	err := camellia_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	camellia_zip_url := "https://downloads.sourceforge.net/project/camellia/Unix_Linux%20Distribution/v2.7.0/CamelliaLib-2.7.0.zip"
	camellia_cmd_zip := exec.Command("curl", "-L", camellia_zip_url, "-o", "package.zip")
	err = camellia_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	camellia_bin_url := "https://downloads.sourceforge.net/project/camellia/Unix_Linux%20Distribution/v2.7.0/CamelliaLib-2.7.0.bin"
	camellia_cmd_bin := exec.Command("curl", "-L", camellia_bin_url, "-o", "binary.bin")
	err = camellia_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	camellia_src_url := "https://downloads.sourceforge.net/project/camellia/Unix_Linux%20Distribution/v2.7.0/CamelliaLib-2.7.0.src.tar.gz"
	camellia_cmd_src := exec.Command("curl", "-L", camellia_src_url, "-o", "source.tar.gz")
	err = camellia_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	camellia_cmd_direct := exec.Command("./binary")
	err = camellia_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
