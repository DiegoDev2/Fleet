package main

// Liblzf - Very small, very fast data compression library
// Homepage: http://oldhome.schmorp.de/marc/liblzf.html

import (
	"fmt"
	
	"os/exec"
)

func installLiblzf() {
	// Método 1: Descargar y extraer .tar.gz
	liblzf_tar_url := "http://dist.schmorp.de/liblzf/liblzf-3.6.tar.gz"
	liblzf_cmd_tar := exec.Command("curl", "-L", liblzf_tar_url, "-o", "package.tar.gz")
	err := liblzf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblzf_zip_url := "http://dist.schmorp.de/liblzf/liblzf-3.6.zip"
	liblzf_cmd_zip := exec.Command("curl", "-L", liblzf_zip_url, "-o", "package.zip")
	err = liblzf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblzf_bin_url := "http://dist.schmorp.de/liblzf/liblzf-3.6.bin"
	liblzf_cmd_bin := exec.Command("curl", "-L", liblzf_bin_url, "-o", "binary.bin")
	err = liblzf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblzf_src_url := "http://dist.schmorp.de/liblzf/liblzf-3.6.src.tar.gz"
	liblzf_cmd_src := exec.Command("curl", "-L", liblzf_src_url, "-o", "source.tar.gz")
	err = liblzf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblzf_cmd_direct := exec.Command("./binary")
	err = liblzf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
