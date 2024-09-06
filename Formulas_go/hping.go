package main

// Hping - Command-line oriented TCP/IP packet assembler/analyzer
// Homepage: http://www.hping.org/

import (
	"fmt"
	
	"os/exec"
)

func installHping() {
	// Método 1: Descargar y extraer .tar.gz
	hping_tar_url := "http://www.hping.org/hping3-20051105.tar.gz"
	hping_cmd_tar := exec.Command("curl", "-L", hping_tar_url, "-o", "package.tar.gz")
	err := hping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hping_zip_url := "http://www.hping.org/hping3-20051105.zip"
	hping_cmd_zip := exec.Command("curl", "-L", hping_zip_url, "-o", "package.zip")
	err = hping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hping_bin_url := "http://www.hping.org/hping3-20051105.bin"
	hping_cmd_bin := exec.Command("curl", "-L", hping_bin_url, "-o", "binary.bin")
	err = hping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hping_src_url := "http://www.hping.org/hping3-20051105.src.tar.gz"
	hping_cmd_src := exec.Command("curl", "-L", hping_src_url, "-o", "source.tar.gz")
	err = hping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hping_cmd_direct := exec.Command("./binary")
	err = hping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
