package main

// Jerm - Communication terminal through serial and TCP/IP interfaces
// Homepage: https://web.archive.org/web/20160719014241/bsddiary.net/jerm/

import (
	"fmt"
	
	"os/exec"
)

func installJerm() {
	// Método 1: Descargar y extraer .tar.gz
	jerm_tar_url := "https://web.archive.org/web/20160719014241/bsddiary.net/jerm/jerm-8096.tar.gz"
	jerm_cmd_tar := exec.Command("curl", "-L", jerm_tar_url, "-o", "package.tar.gz")
	err := jerm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jerm_zip_url := "https://web.archive.org/web/20160719014241/bsddiary.net/jerm/jerm-8096.zip"
	jerm_cmd_zip := exec.Command("curl", "-L", jerm_zip_url, "-o", "package.zip")
	err = jerm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jerm_bin_url := "https://web.archive.org/web/20160719014241/bsddiary.net/jerm/jerm-8096.bin"
	jerm_cmd_bin := exec.Command("curl", "-L", jerm_bin_url, "-o", "binary.bin")
	err = jerm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jerm_src_url := "https://web.archive.org/web/20160719014241/bsddiary.net/jerm/jerm-8096.src.tar.gz"
	jerm_cmd_src := exec.Command("curl", "-L", jerm_src_url, "-o", "source.tar.gz")
	err = jerm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jerm_cmd_direct := exec.Command("./binary")
	err = jerm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
