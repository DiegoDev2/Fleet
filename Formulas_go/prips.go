package main

// Prips - Print the IP addresses in a given range
// Homepage: https://devel.ringlet.net/sysutils/prips/

import (
	"fmt"
	
	"os/exec"
)

func installPrips() {
	// Método 1: Descargar y extraer .tar.gz
	prips_tar_url := "https://devel.ringlet.net/files/sys/prips/prips-1.2.0.tar.xz"
	prips_cmd_tar := exec.Command("curl", "-L", prips_tar_url, "-o", "package.tar.gz")
	err := prips_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prips_zip_url := "https://devel.ringlet.net/files/sys/prips/prips-1.2.0.tar.xz"
	prips_cmd_zip := exec.Command("curl", "-L", prips_zip_url, "-o", "package.zip")
	err = prips_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prips_bin_url := "https://devel.ringlet.net/files/sys/prips/prips-1.2.0.tar.xz"
	prips_cmd_bin := exec.Command("curl", "-L", prips_bin_url, "-o", "binary.bin")
	err = prips_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prips_src_url := "https://devel.ringlet.net/files/sys/prips/prips-1.2.0.tar.xz"
	prips_cmd_src := exec.Command("curl", "-L", prips_src_url, "-o", "source.tar.gz")
	err = prips_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prips_cmd_direct := exec.Command("./binary")
	err = prips_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
