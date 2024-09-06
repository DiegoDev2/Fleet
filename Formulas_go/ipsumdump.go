package main

// Ipsumdump - Summarizes TCP/IP dump files into a self-describing ASCII format
// Homepage: https://read.seas.harvard.edu/~kohler/ipsumdump/

import (
	"fmt"
	
	"os/exec"
)

func installIpsumdump() {
	// Método 1: Descargar y extraer .tar.gz
	ipsumdump_tar_url := "https://read.seas.harvard.edu/~kohler/ipsumdump/ipsumdump-1.86.tar.gz"
	ipsumdump_cmd_tar := exec.Command("curl", "-L", ipsumdump_tar_url, "-o", "package.tar.gz")
	err := ipsumdump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipsumdump_zip_url := "https://read.seas.harvard.edu/~kohler/ipsumdump/ipsumdump-1.86.zip"
	ipsumdump_cmd_zip := exec.Command("curl", "-L", ipsumdump_zip_url, "-o", "package.zip")
	err = ipsumdump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipsumdump_bin_url := "https://read.seas.harvard.edu/~kohler/ipsumdump/ipsumdump-1.86.bin"
	ipsumdump_cmd_bin := exec.Command("curl", "-L", ipsumdump_bin_url, "-o", "binary.bin")
	err = ipsumdump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipsumdump_src_url := "https://read.seas.harvard.edu/~kohler/ipsumdump/ipsumdump-1.86.src.tar.gz"
	ipsumdump_cmd_src := exec.Command("curl", "-L", ipsumdump_src_url, "-o", "source.tar.gz")
	err = ipsumdump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipsumdump_cmd_direct := exec.Command("./binary")
	err = ipsumdump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
