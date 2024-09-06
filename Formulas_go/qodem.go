package main

// Qodem - Terminal emulator and BBS client
// Homepage: https://qodem.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installQodem() {
	// Método 1: Descargar y extraer .tar.gz
	qodem_tar_url := "https://downloads.sourceforge.net/project/qodem/qodem/1.0.1/qodem-1.0.1.tar.gz"
	qodem_cmd_tar := exec.Command("curl", "-L", qodem_tar_url, "-o", "package.tar.gz")
	err := qodem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qodem_zip_url := "https://downloads.sourceforge.net/project/qodem/qodem/1.0.1/qodem-1.0.1.zip"
	qodem_cmd_zip := exec.Command("curl", "-L", qodem_zip_url, "-o", "package.zip")
	err = qodem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qodem_bin_url := "https://downloads.sourceforge.net/project/qodem/qodem/1.0.1/qodem-1.0.1.bin"
	qodem_cmd_bin := exec.Command("curl", "-L", qodem_bin_url, "-o", "binary.bin")
	err = qodem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qodem_src_url := "https://downloads.sourceforge.net/project/qodem/qodem/1.0.1/qodem-1.0.1.src.tar.gz"
	qodem_cmd_src := exec.Command("curl", "-L", qodem_src_url, "-o", "source.tar.gz")
	err = qodem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qodem_cmd_direct := exec.Command("./binary")
	err = qodem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
