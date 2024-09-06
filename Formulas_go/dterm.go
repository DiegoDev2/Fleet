package main

// Dterm - Terminal emulator for use with xterm and friends
// Homepage: https://web.archive.org/web/20230125064301/http://www.knossos.net.nz/resources/free-software/dterm/

import (
	"fmt"
	
	"os/exec"
)

func installDterm() {
	// Método 1: Descargar y extraer .tar.gz
	dterm_tar_url := "https://web.archive.org/web/20230125064301/http://www.knossos.net.nz/downloads/dterm-0.5.tgz"
	dterm_cmd_tar := exec.Command("curl", "-L", dterm_tar_url, "-o", "package.tar.gz")
	err := dterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dterm_zip_url := "https://web.archive.org/web/20230125064301/http://www.knossos.net.nz/downloads/dterm-0.5.tgz"
	dterm_cmd_zip := exec.Command("curl", "-L", dterm_zip_url, "-o", "package.zip")
	err = dterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dterm_bin_url := "https://web.archive.org/web/20230125064301/http://www.knossos.net.nz/downloads/dterm-0.5.tgz"
	dterm_cmd_bin := exec.Command("curl", "-L", dterm_bin_url, "-o", "binary.bin")
	err = dterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dterm_src_url := "https://web.archive.org/web/20230125064301/http://www.knossos.net.nz/downloads/dterm-0.5.tgz"
	dterm_cmd_src := exec.Command("curl", "-L", dterm_src_url, "-o", "source.tar.gz")
	err = dterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dterm_cmd_direct := exec.Command("./binary")
	err = dterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
