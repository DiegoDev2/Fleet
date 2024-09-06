package main

// Sqtop - Display information about active connections for a Squid proxy
// Homepage: https://github.com/paleg/sqtop

import (
	"fmt"
	
	"os/exec"
)

func installSqtop() {
	// Método 1: Descargar y extraer .tar.gz
	sqtop_tar_url := "https://github.com/paleg/sqtop/archive/refs/tags/v2015-02-08.tar.gz"
	sqtop_cmd_tar := exec.Command("curl", "-L", sqtop_tar_url, "-o", "package.tar.gz")
	err := sqtop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqtop_zip_url := "https://github.com/paleg/sqtop/archive/refs/tags/v2015-02-08.zip"
	sqtop_cmd_zip := exec.Command("curl", "-L", sqtop_zip_url, "-o", "package.zip")
	err = sqtop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqtop_bin_url := "https://github.com/paleg/sqtop/archive/refs/tags/v2015-02-08.bin"
	sqtop_cmd_bin := exec.Command("curl", "-L", sqtop_bin_url, "-o", "binary.bin")
	err = sqtop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqtop_src_url := "https://github.com/paleg/sqtop/archive/refs/tags/v2015-02-08.src.tar.gz"
	sqtop_cmd_src := exec.Command("curl", "-L", sqtop_src_url, "-o", "source.tar.gz")
	err = sqtop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqtop_cmd_direct := exec.Command("./binary")
	err = sqtop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
