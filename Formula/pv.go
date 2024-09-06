package main

// Pv - Monitor data's progress through a pipe
// Homepage: https://www.ivarch.com/programs/pv.shtml

import (
	"fmt"
	
	"os/exec"
)

func installPv() {
	// Método 1: Descargar y extraer .tar.gz
	pv_tar_url := "https://www.ivarch.com/programs/sources/pv-1.8.13.tar.gz"
	pv_cmd_tar := exec.Command("curl", "-L", pv_tar_url, "-o", "package.tar.gz")
	err := pv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pv_zip_url := "https://www.ivarch.com/programs/sources/pv-1.8.13.zip"
	pv_cmd_zip := exec.Command("curl", "-L", pv_zip_url, "-o", "package.zip")
	err = pv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pv_bin_url := "https://www.ivarch.com/programs/sources/pv-1.8.13.bin"
	pv_cmd_bin := exec.Command("curl", "-L", pv_bin_url, "-o", "binary.bin")
	err = pv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pv_src_url := "https://www.ivarch.com/programs/sources/pv-1.8.13.src.tar.gz"
	pv_cmd_src := exec.Command("curl", "-L", pv_src_url, "-o", "source.tar.gz")
	err = pv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pv_cmd_direct := exec.Command("./binary")
	err = pv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
