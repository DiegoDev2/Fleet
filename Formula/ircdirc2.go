package main

// IrcdIrc2 - Original IRC server daemon
// Homepage: http://www.irc.org/

import (
	"fmt"
	
	"os/exec"
)

func installIrcdIrc2() {
	// Método 1: Descargar y extraer .tar.gz
	ircdirc2_tar_url := "http://www.irc.org/ftp/irc/server/irc2.11.2p3.tgz"
	ircdirc2_cmd_tar := exec.Command("curl", "-L", ircdirc2_tar_url, "-o", "package.tar.gz")
	err := ircdirc2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ircdirc2_zip_url := "http://www.irc.org/ftp/irc/server/irc2.11.2p3.tgz"
	ircdirc2_cmd_zip := exec.Command("curl", "-L", ircdirc2_zip_url, "-o", "package.zip")
	err = ircdirc2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ircdirc2_bin_url := "http://www.irc.org/ftp/irc/server/irc2.11.2p3.tgz"
	ircdirc2_cmd_bin := exec.Command("curl", "-L", ircdirc2_bin_url, "-o", "binary.bin")
	err = ircdirc2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ircdirc2_src_url := "http://www.irc.org/ftp/irc/server/irc2.11.2p3.tgz"
	ircdirc2_cmd_src := exec.Command("curl", "-L", ircdirc2_src_url, "-o", "source.tar.gz")
	err = ircdirc2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ircdirc2_cmd_direct := exec.Command("./binary")
	err = ircdirc2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
