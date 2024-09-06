package main

// Jid - Json incremental digger
// Homepage: https://github.com/simeji/jid

import (
	"fmt"
	
	"os/exec"
)

func installJid() {
	// Método 1: Descargar y extraer .tar.gz
	jid_tar_url := "https://github.com/simeji/jid/archive/refs/tags/v0.7.6.tar.gz"
	jid_cmd_tar := exec.Command("curl", "-L", jid_tar_url, "-o", "package.tar.gz")
	err := jid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jid_zip_url := "https://github.com/simeji/jid/archive/refs/tags/v0.7.6.zip"
	jid_cmd_zip := exec.Command("curl", "-L", jid_zip_url, "-o", "package.zip")
	err = jid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jid_bin_url := "https://github.com/simeji/jid/archive/refs/tags/v0.7.6.bin"
	jid_cmd_bin := exec.Command("curl", "-L", jid_bin_url, "-o", "binary.bin")
	err = jid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jid_src_url := "https://github.com/simeji/jid/archive/refs/tags/v0.7.6.src.tar.gz"
	jid_cmd_src := exec.Command("curl", "-L", jid_src_url, "-o", "source.tar.gz")
	err = jid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jid_cmd_direct := exec.Command("./binary")
	err = jid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
