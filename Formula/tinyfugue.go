package main

// TinyFugue - Programmable MUD client
// Homepage: https://tinyfugue.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installTinyFugue() {
	// Método 1: Descargar y extraer .tar.gz
	tinyfugue_tar_url := "https://downloads.sourceforge.net/project/tinyfugue/tinyfugue/5.0%20beta%208/tf-50b8.tar.gz"
	tinyfugue_cmd_tar := exec.Command("curl", "-L", tinyfugue_tar_url, "-o", "package.tar.gz")
	err := tinyfugue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinyfugue_zip_url := "https://downloads.sourceforge.net/project/tinyfugue/tinyfugue/5.0%20beta%208/tf-50b8.zip"
	tinyfugue_cmd_zip := exec.Command("curl", "-L", tinyfugue_zip_url, "-o", "package.zip")
	err = tinyfugue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinyfugue_bin_url := "https://downloads.sourceforge.net/project/tinyfugue/tinyfugue/5.0%20beta%208/tf-50b8.bin"
	tinyfugue_cmd_bin := exec.Command("curl", "-L", tinyfugue_bin_url, "-o", "binary.bin")
	err = tinyfugue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinyfugue_src_url := "https://downloads.sourceforge.net/project/tinyfugue/tinyfugue/5.0%20beta%208/tf-50b8.src.tar.gz"
	tinyfugue_cmd_src := exec.Command("curl", "-L", tinyfugue_src_url, "-o", "source.tar.gz")
	err = tinyfugue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinyfugue_cmd_direct := exec.Command("./binary")
	err = tinyfugue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libnet")
	exec.Command("latte", "install", "libnet").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
