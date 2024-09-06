package main

// Log4shib - Forked version of log4cpp for the Shibboleth project
// Homepage: https://wiki.shibboleth.net/confluence/display/OpenSAML/log4shib

import (
	"fmt"
	
	"os/exec"
)

func installLog4shib() {
	// Método 1: Descargar y extraer .tar.gz
	log4shib_tar_url := "https://shibboleth.net/downloads/log4shib/2.0.1/log4shib-2.0.1.tar.gz"
	log4shib_cmd_tar := exec.Command("curl", "-L", log4shib_tar_url, "-o", "package.tar.gz")
	err := log4shib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	log4shib_zip_url := "https://shibboleth.net/downloads/log4shib/2.0.1/log4shib-2.0.1.zip"
	log4shib_cmd_zip := exec.Command("curl", "-L", log4shib_zip_url, "-o", "package.zip")
	err = log4shib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	log4shib_bin_url := "https://shibboleth.net/downloads/log4shib/2.0.1/log4shib-2.0.1.bin"
	log4shib_cmd_bin := exec.Command("curl", "-L", log4shib_bin_url, "-o", "binary.bin")
	err = log4shib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	log4shib_src_url := "https://shibboleth.net/downloads/log4shib/2.0.1/log4shib-2.0.1.src.tar.gz"
	log4shib_cmd_src := exec.Command("curl", "-L", log4shib_src_url, "-o", "source.tar.gz")
	err = log4shib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	log4shib_cmd_direct := exec.Command("./binary")
	err = log4shib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
