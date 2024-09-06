package main

// Daemonlogger - Network packet logger and soft tap daemon
// Homepage: https://sourceforge.net/projects/daemonlogger/

import (
	"fmt"
	
	"os/exec"
)

func installDaemonlogger() {
	// Método 1: Descargar y extraer .tar.gz
	daemonlogger_tar_url := "https://downloads.sourceforge.net/project/daemonlogger/daemonlogger-1.2.1.tar.gz"
	daemonlogger_cmd_tar := exec.Command("curl", "-L", daemonlogger_tar_url, "-o", "package.tar.gz")
	err := daemonlogger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	daemonlogger_zip_url := "https://downloads.sourceforge.net/project/daemonlogger/daemonlogger-1.2.1.zip"
	daemonlogger_cmd_zip := exec.Command("curl", "-L", daemonlogger_zip_url, "-o", "package.zip")
	err = daemonlogger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	daemonlogger_bin_url := "https://downloads.sourceforge.net/project/daemonlogger/daemonlogger-1.2.1.bin"
	daemonlogger_cmd_bin := exec.Command("curl", "-L", daemonlogger_bin_url, "-o", "binary.bin")
	err = daemonlogger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	daemonlogger_src_url := "https://downloads.sourceforge.net/project/daemonlogger/daemonlogger-1.2.1.src.tar.gz"
	daemonlogger_cmd_src := exec.Command("curl", "-L", daemonlogger_src_url, "-o", "source.tar.gz")
	err = daemonlogger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	daemonlogger_cmd_direct := exec.Command("./binary")
	err = daemonlogger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libdnet")
	exec.Command("latte", "install", "libdnet").Run()
}
