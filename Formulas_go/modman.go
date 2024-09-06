package main

// Modman - Module deployment script geared towards Magento development
// Homepage: https://github.com/colinmollenhour/modman

import (
	"fmt"
	
	"os/exec"
)

func installModman() {
	// Método 1: Descargar y extraer .tar.gz
	modman_tar_url := "https://github.com/colinmollenhour/modman/archive/refs/tags/1.14.tar.gz"
	modman_cmd_tar := exec.Command("curl", "-L", modman_tar_url, "-o", "package.tar.gz")
	err := modman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	modman_zip_url := "https://github.com/colinmollenhour/modman/archive/refs/tags/1.14.zip"
	modman_cmd_zip := exec.Command("curl", "-L", modman_zip_url, "-o", "package.zip")
	err = modman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	modman_bin_url := "https://github.com/colinmollenhour/modman/archive/refs/tags/1.14.bin"
	modman_cmd_bin := exec.Command("curl", "-L", modman_bin_url, "-o", "binary.bin")
	err = modman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	modman_src_url := "https://github.com/colinmollenhour/modman/archive/refs/tags/1.14.src.tar.gz"
	modman_cmd_src := exec.Command("curl", "-L", modman_src_url, "-o", "source.tar.gz")
	err = modman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	modman_cmd_direct := exec.Command("./binary")
	err = modman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
