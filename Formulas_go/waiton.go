package main

// WaitOn - Provides shell scripts with access to kqueue(3)
// Homepage: https://www.freshports.org/sysutils/wait_on/

import (
	"fmt"
	
	"os/exec"
)

func installWaitOn() {
	// Método 1: Descargar y extraer .tar.gz
	waiton_tar_url := "https://pkg.freebsd.org/ports-distfiles/wait_on-1.1.tar.gz"
	waiton_cmd_tar := exec.Command("curl", "-L", waiton_tar_url, "-o", "package.tar.gz")
	err := waiton_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	waiton_zip_url := "https://pkg.freebsd.org/ports-distfiles/wait_on-1.1.zip"
	waiton_cmd_zip := exec.Command("curl", "-L", waiton_zip_url, "-o", "package.zip")
	err = waiton_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	waiton_bin_url := "https://pkg.freebsd.org/ports-distfiles/wait_on-1.1.bin"
	waiton_cmd_bin := exec.Command("curl", "-L", waiton_bin_url, "-o", "binary.bin")
	err = waiton_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	waiton_src_url := "https://pkg.freebsd.org/ports-distfiles/wait_on-1.1.src.tar.gz"
	waiton_cmd_src := exec.Command("curl", "-L", waiton_src_url, "-o", "source.tar.gz")
	err = waiton_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	waiton_cmd_direct := exec.Command("./binary")
	err = waiton_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bsdmake")
exec.Command("latte", "install", "bsdmake")
}
