package main

// Lft - Layer Four Traceroute (LFT), an advanced traceroute tool
// Homepage: https://pwhois.org/lft/

import (
	"fmt"
	
	"os/exec"
)

func installLft() {
	// Método 1: Descargar y extraer .tar.gz
	lft_tar_url := "https://pwhois.org/dl/index.who?file=lft-3.91.tar.gz"
	lft_cmd_tar := exec.Command("curl", "-L", lft_tar_url, "-o", "package.tar.gz")
	err := lft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lft_zip_url := "https://pwhois.org/dl/index.who?file=lft-3.91.zip"
	lft_cmd_zip := exec.Command("curl", "-L", lft_zip_url, "-o", "package.zip")
	err = lft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lft_bin_url := "https://pwhois.org/dl/index.who?file=lft-3.91.bin"
	lft_cmd_bin := exec.Command("curl", "-L", lft_bin_url, "-o", "binary.bin")
	err = lft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lft_src_url := "https://pwhois.org/dl/index.who?file=lft-3.91.src.tar.gz"
	lft_cmd_src := exec.Command("curl", "-L", lft_src_url, "-o", "source.tar.gz")
	err = lft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lft_cmd_direct := exec.Command("./binary")
	err = lft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
