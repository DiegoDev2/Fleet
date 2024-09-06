package main

// Logcheck - Mail anomalies in the system logfiles to the administrator
// Homepage: https://packages.debian.org/sid/logcheck

import (
	"fmt"
	
	"os/exec"
)

func installLogcheck() {
	// Método 1: Descargar y extraer .tar.gz
	logcheck_tar_url := "https://deb.debian.org/debian/pool/main/l/logcheck/logcheck_1.4.3.tar.xz"
	logcheck_cmd_tar := exec.Command("curl", "-L", logcheck_tar_url, "-o", "package.tar.gz")
	err := logcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	logcheck_zip_url := "https://deb.debian.org/debian/pool/main/l/logcheck/logcheck_1.4.3.tar.xz"
	logcheck_cmd_zip := exec.Command("curl", "-L", logcheck_zip_url, "-o", "package.zip")
	err = logcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	logcheck_bin_url := "https://deb.debian.org/debian/pool/main/l/logcheck/logcheck_1.4.3.tar.xz"
	logcheck_cmd_bin := exec.Command("curl", "-L", logcheck_bin_url, "-o", "binary.bin")
	err = logcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	logcheck_src_url := "https://deb.debian.org/debian/pool/main/l/logcheck/logcheck_1.4.3.tar.xz"
	logcheck_cmd_src := exec.Command("curl", "-L", logcheck_src_url, "-o", "source.tar.gz")
	err = logcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	logcheck_cmd_direct := exec.Command("./binary")
	err = logcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
}
