package main

// Grepcidr - Filter IP addresses matching IPv4 CIDR/network specification
// Homepage: https://www.pc-tools.net/unix/grepcidr/

import (
	"fmt"
	
	"os/exec"
)

func installGrepcidr() {
	// Método 1: Descargar y extraer .tar.gz
	grepcidr_tar_url := "https://www.pc-tools.net/files/unix/grepcidr-2.0.tar.gz"
	grepcidr_cmd_tar := exec.Command("curl", "-L", grepcidr_tar_url, "-o", "package.tar.gz")
	err := grepcidr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grepcidr_zip_url := "https://www.pc-tools.net/files/unix/grepcidr-2.0.zip"
	grepcidr_cmd_zip := exec.Command("curl", "-L", grepcidr_zip_url, "-o", "package.zip")
	err = grepcidr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grepcidr_bin_url := "https://www.pc-tools.net/files/unix/grepcidr-2.0.bin"
	grepcidr_cmd_bin := exec.Command("curl", "-L", grepcidr_bin_url, "-o", "binary.bin")
	err = grepcidr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grepcidr_src_url := "https://www.pc-tools.net/files/unix/grepcidr-2.0.src.tar.gz"
	grepcidr_cmd_src := exec.Command("curl", "-L", grepcidr_src_url, "-o", "source.tar.gz")
	err = grepcidr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grepcidr_cmd_direct := exec.Command("./binary")
	err = grepcidr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
