package main

// Cmix - Data compression program with high compression ratio
// Homepage: https://www.byronknoll.com/cmix.html

import (
	"fmt"
	
	"os/exec"
)

func installCmix() {
	// Método 1: Descargar y extraer .tar.gz
	cmix_tar_url := "https://github.com/byronknoll/cmix/archive/refs/tags/v20.tar.gz"
	cmix_cmd_tar := exec.Command("curl", "-L", cmix_tar_url, "-o", "package.tar.gz")
	err := cmix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmix_zip_url := "https://github.com/byronknoll/cmix/archive/refs/tags/v20.zip"
	cmix_cmd_zip := exec.Command("curl", "-L", cmix_zip_url, "-o", "package.zip")
	err = cmix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmix_bin_url := "https://github.com/byronknoll/cmix/archive/refs/tags/v20.bin"
	cmix_cmd_bin := exec.Command("curl", "-L", cmix_bin_url, "-o", "binary.bin")
	err = cmix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmix_src_url := "https://github.com/byronknoll/cmix/archive/refs/tags/v20.src.tar.gz"
	cmix_cmd_src := exec.Command("curl", "-L", cmix_src_url, "-o", "source.tar.gz")
	err = cmix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmix_cmd_direct := exec.Command("./binary")
	err = cmix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
