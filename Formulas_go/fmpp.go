package main

// Fmpp - Text file preprocessing tool using FreeMarker templates
// Homepage: https://fmpp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installFmpp() {
	// Método 1: Descargar y extraer .tar.gz
	fmpp_tar_url := "https://downloads.sourceforge.net/project/fmpp/fmpp/0.9.16/fmpp_0.9.16.tar.gz"
	fmpp_cmd_tar := exec.Command("curl", "-L", fmpp_tar_url, "-o", "package.tar.gz")
	err := fmpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fmpp_zip_url := "https://downloads.sourceforge.net/project/fmpp/fmpp/0.9.16/fmpp_0.9.16.zip"
	fmpp_cmd_zip := exec.Command("curl", "-L", fmpp_zip_url, "-o", "package.zip")
	err = fmpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fmpp_bin_url := "https://downloads.sourceforge.net/project/fmpp/fmpp/0.9.16/fmpp_0.9.16.bin"
	fmpp_cmd_bin := exec.Command("curl", "-L", fmpp_bin_url, "-o", "binary.bin")
	err = fmpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fmpp_src_url := "https://downloads.sourceforge.net/project/fmpp/fmpp/0.9.16/fmpp_0.9.16.src.tar.gz"
	fmpp_cmd_src := exec.Command("curl", "-L", fmpp_src_url, "-o", "source.tar.gz")
	err = fmpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fmpp_cmd_direct := exec.Command("./binary")
	err = fmpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
