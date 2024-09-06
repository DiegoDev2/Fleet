package main

// Proftpd - Highly configurable GPL-licensed FTP server software
// Homepage: http://www.proftpd.org/

import (
	"fmt"
	
	"os/exec"
)

func installProftpd() {
	// Método 1: Descargar y extraer .tar.gz
	proftpd_tar_url := "https://github.com/proftpd/proftpd/archive/refs/tags/v1.3.8b.tar.gz"
	proftpd_cmd_tar := exec.Command("curl", "-L", proftpd_tar_url, "-o", "package.tar.gz")
	err := proftpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proftpd_zip_url := "https://github.com/proftpd/proftpd/archive/refs/tags/v1.3.8b.zip"
	proftpd_cmd_zip := exec.Command("curl", "-L", proftpd_zip_url, "-o", "package.zip")
	err = proftpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proftpd_bin_url := "https://github.com/proftpd/proftpd/archive/refs/tags/v1.3.8b.bin"
	proftpd_cmd_bin := exec.Command("curl", "-L", proftpd_bin_url, "-o", "binary.bin")
	err = proftpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proftpd_src_url := "https://github.com/proftpd/proftpd/archive/refs/tags/v1.3.8b.src.tar.gz"
	proftpd_cmd_src := exec.Command("curl", "-L", proftpd_src_url, "-o", "source.tar.gz")
	err = proftpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proftpd_cmd_direct := exec.Command("./binary")
	err = proftpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
