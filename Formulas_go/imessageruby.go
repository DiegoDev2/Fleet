package main

// ImessageRuby - Command-line tool to send iMessage
// Homepage: https://github.com/linjunpop/imessage

import (
	"fmt"
	
	"os/exec"
)

func installImessageRuby() {
	// Método 1: Descargar y extraer .tar.gz
	imessageruby_tar_url := "https://github.com/linjunpop/imessage/archive/refs/tags/v0.4.0.tar.gz"
	imessageruby_cmd_tar := exec.Command("curl", "-L", imessageruby_tar_url, "-o", "package.tar.gz")
	err := imessageruby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imessageruby_zip_url := "https://github.com/linjunpop/imessage/archive/refs/tags/v0.4.0.zip"
	imessageruby_cmd_zip := exec.Command("curl", "-L", imessageruby_zip_url, "-o", "package.zip")
	err = imessageruby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imessageruby_bin_url := "https://github.com/linjunpop/imessage/archive/refs/tags/v0.4.0.bin"
	imessageruby_cmd_bin := exec.Command("curl", "-L", imessageruby_bin_url, "-o", "binary.bin")
	err = imessageruby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imessageruby_src_url := "https://github.com/linjunpop/imessage/archive/refs/tags/v0.4.0.src.tar.gz"
	imessageruby_cmd_src := exec.Command("curl", "-L", imessageruby_src_url, "-o", "source.tar.gz")
	err = imessageruby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imessageruby_cmd_direct := exec.Command("./binary")
	err = imessageruby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
