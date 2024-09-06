package main

// Apparix - File system navigation via bookmarking directories
// Homepage: https://micans.org/apparix/

import (
	"fmt"
	
	"os/exec"
)

func installApparix() {
	// Método 1: Descargar y extraer .tar.gz
	apparix_tar_url := "https://micans.org/apparix/src/apparix-11-062.tar.gz"
	apparix_cmd_tar := exec.Command("curl", "-L", apparix_tar_url, "-o", "package.tar.gz")
	err := apparix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apparix_zip_url := "https://micans.org/apparix/src/apparix-11-062.zip"
	apparix_cmd_zip := exec.Command("curl", "-L", apparix_zip_url, "-o", "package.zip")
	err = apparix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apparix_bin_url := "https://micans.org/apparix/src/apparix-11-062.bin"
	apparix_cmd_bin := exec.Command("curl", "-L", apparix_bin_url, "-o", "binary.bin")
	err = apparix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apparix_src_url := "https://micans.org/apparix/src/apparix-11-062.src.tar.gz"
	apparix_cmd_src := exec.Command("curl", "-L", apparix_src_url, "-o", "source.tar.gz")
	err = apparix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apparix_cmd_direct := exec.Command("./binary")
	err = apparix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
